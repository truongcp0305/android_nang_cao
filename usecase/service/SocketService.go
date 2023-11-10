package service

import (
	"android-service/model"
	"android-service/usecase/repository"
	"fmt"
	"time"
)

type SocketService struct {
	db        repository.Database
	clients   map[string]*model.MatchStatus
	questions []model.Word
	conns     map[string]*Conn
}

type Conn struct {
	id   string
	mess chan bool
	res  chan string
}

func NewSocketService(db repository.Database) *SocketService {
	return &SocketService{
		clients:   make(map[string]*model.MatchStatus),
		db:        db,
		questions: []model.Word{},
		conns:     make(map[string]*Conn),
	}
}

func (s *SocketService) GetStatus(stt model.MatchStatus) model.MatchStatus {
	if _, ok := s.clients[stt.Id]; !ok {
		return model.MatchStatus{
			Message: "not found client",
		}
	}
	s.clients[stt.Id].Status = stt.Status
	s.clients[stt.Id].Point = stt.Point
	for _, cl := range s.clients {
		if cl.Id != stt.Id {
			s.conns[stt.Id].mess <- true
			return *cl
		}
	}
	return model.MatchStatus{
		Message: "not found op",
	}
}

func (s *SocketService) Join(id string, level string) model.MatchStatus {
	if len(s.questions) == 0 {
		s.questions, _ = s.db.GetWordsForQuestion(level)
	}
	if _, ok := s.clients[id]; ok {
		return *s.clients[id]
	}
	if len(s.clients) == 2 {
		//s.clients = make(map[string]*model.MatchStatus)
		return model.MatchStatus{
			Message: "sever quá tải",
		}
	}
	s.clients[id] = &model.MatchStatus{
		Id:     id,
		Status: "finding",
		Point:  "0",
	}
	if len(s.clients) == 2 {
		for _, c := range s.clients {
			s.clients[c.Id].Status = "matching"
			s.clients[c.Id].Questions = s.questions
		}
		s.HandleConnection()
	}
	return *s.clients[id]
}

func (s *SocketService) HandleConnection() {
	for id := range s.clients {
		s.conns[id] = &Conn{
			id:   id,
			mess: make(chan bool, 1),
			res:  make(chan string, 1),
		}
		go s.conns[id].KeepConection(s.conns[id].mess, s.conns[id].res)
		go s.HandleDisconnect(id)
	}
}

func (s *SocketService) HandleDisconnect(id string) {
	for {
		select {
		case mes, _ := <-s.conns[id].res:
			if mes == id {
				s.Leave(id)
				delete(s.conns, id)
				fmt.Println("disconnect " + id)
				return
			}
		}
	}
}

func (c *Conn) KeepConection(conn <-chan bool, disConn chan<- string) {
	fmt.Println("kepp " + c.id)
	timer := time.NewTimer(15 * time.Second)
	for {
		select {
		case mes, ok := <-conn:
			if !ok {
				disConn <- c.id
				return
			}
			if mes {
				fmt.Println("message from " + c.id)
				timer.Reset(15 * time.Second)
			}
		case <-timer.C:
			disConn <- c.id
			return
		}
	}
}

func (s *SocketService) Leave(id string) {
	if _, ok := s.clients[id]; !ok {
		return
	}
	s.clients[id].Status = "done"
	for _, cl := range s.clients {
		if cl.Status != "done" {
			return
		}
	}
	s.questions = []model.Word{}
	for k := range s.clients {
		delete(s.clients, k)
	}
}
