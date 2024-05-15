package handlers

import (
	"fmt"
	"go-workspace/ext-crypt-backend/models"
	"go-workspace/ext-crypt-backend/repository"
	"net/http"

	"github.com/labstack/echo"
)

func SignIn(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	fmt.Println("user: ", user)

	email, err := repository.SignIn(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Message{
			Msg: "Sign in with Google failed: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &models.Message{
		Msg: email,
	})
}
