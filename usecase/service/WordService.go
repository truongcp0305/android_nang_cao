package service

import (
	"android-service/model"
	"android-service/usecase/repository"
	"errors"
	"strconv"
	"strings"
)

type WordService struct {
	db repository.Database
}

func NewWordService(db repository.Database) *WordService {
	return &WordService{
		db: db,
	}
}

func (s *WordService) InsertWord(texts string, level string) error {
	lev, err := strconv.ParseInt(level, 10, 64)
	if err != nil {
		return errors.New("Invalid query params")
	}
	text := strings.Split(texts, ",")
	words := []model.Word{}
	for _, t := range text {
		word := model.Word{
			Text:  t,
			Level: lev,
		}
		words = append(words, word)
	}
	err = s.db.InsertWord(words)
	if err != nil {
		return err
	}
	return nil
}
