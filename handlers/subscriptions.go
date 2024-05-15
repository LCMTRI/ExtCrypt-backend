package handlers

import (
	"go-workspace/ext-crypt-backend/models"
	"go-workspace/ext-crypt-backend/repository"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllSubscriptions(c echo.Context) error {
	subscriptions, err := repository.GetAllSubscriptions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Message{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, subscriptions)
}
