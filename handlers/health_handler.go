package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheckHandler(c echo.Context) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return c.String(http.StatusInternalServerError, "error get jwt cookie")
	}
	fmt.Println("cookie: ", cookie.Value)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err)
	// }
	return c.String(http.StatusOK, "oke")
}
