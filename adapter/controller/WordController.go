package controller

import (
	"android-service/adapter/incoming"
	"android-service/usecase/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WordController struct {
	wordService service.WordService
}

func NewWordController(w service.WordService) WordController {
	return WordController{
		wordService: w,
	}
}

func (wC *WordController) Insert(c echo.Context) error {
	var params incoming.InsertWord
	err := c.Bind(&params)
	if err != nil || params.Texts == "" || params.Level == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err = wC.wordService.InsertWord(params.Texts, params.Level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Insert sucsess")
}

func (wC *WordController) GetQuestions(c echo.Context) error {
	var param incoming.GetQuestion
	err := c.Bind(&param)
	if param.Level == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	result, err := wC.wordService.GetQuestions(param.Level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
