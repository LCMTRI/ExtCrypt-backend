package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheckHandler(c echo.Context) error {

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err)
	// }
	return c.String(http.StatusOK, "Health check ok")
}
