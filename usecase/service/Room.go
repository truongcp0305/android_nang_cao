package service

import (
	"log"

	"github.com/gorilla/websocket"
)

type Room struct {
	clients map[*websocket.Conn]bool
}

func NewRoom() *Room {
	return &Room{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (r *Room) Join(client *websocket.Conn) {
	if len(r.clients) == 2 {
		return
	}
	r.clients[client] = true
	ok := "connect sucsess"
	err := client.WriteMessage(websocket.TextMessage, []byte(ok))
	if err != nil {
		log.Println("Lỗi khi gửi thông điệp:", err)
	}
	if len(r.clients) > 1 {
		for client := range r.clients {
			message := "Connect other sucsess"
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Lỗi khi gửi thông điệp:", err)
			}
		}
	}
}

func (r *Room) Leave(client *websocket.Conn) {
	delete(r.clients, client)
}

func (r *Room) Broadcast(message []byte, sender *websocket.Conn) {
	for client := range r.clients {
		if client != sender {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Lỗi khi gửi thông điệp:", err)
			}
		}
	}
}

// func NewRoom() *Room {
// 	return &Room{
// 		clients: make(map[*websocket.Conn]bool),
// 	}
// }

// func (r *Room) Join(client *websocket.Conn) {
// 	r.clients[client] = true
// 	ok := "connect sucsess"
// 	err := client.WriteMessage(websocket.TextMessage, []byte(ok))
// 	if err != nil {
// 		log.Println("Lỗi khi gửi thông điệp:", err)
// 	}
// 	if len(r.clients) > 1 {
// 		for client := range r.clients {
// 			message := "Connect other sucsess"
// 			err := client.WriteMessage(websocket.TextMessage, []byte(message))
// 			if err != nil {
// 				log.Println("Lỗi khi gửi thông điệp:", err)
// 			}
// 		}
// 	}
// }

// func (r *Room) Leave(client *websocket.Conn) {
// 	delete(r.clients, client)
// }

// func (r *Room) Broadcast(message []byte, sender *websocket.Conn) {
// 	for client := range r.clients {
// 		if client != sender {
// 			err := client.WriteMessage(websocket.TextMessage, message)
// 			if err != nil {
// 				log.Println("Lỗi khi gửi thông điệp:", err)
// 			}
// 		}
// 	}
// }
