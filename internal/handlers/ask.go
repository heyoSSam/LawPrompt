package handlers

import (
	"LawPrompt/internal/model"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type RequestBody struct {
	Question string `json:"question"`
}

func PostAskModel(c echo.Context) error {
	var req RequestBody
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	answer, err := model.AskModel(req.Question)
	if err != nil {
		log.Println("xxx")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"answer": answer,
	})
}
