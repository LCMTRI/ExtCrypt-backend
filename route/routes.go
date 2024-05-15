package route

import (
	"go-workspace/ext-crypt-backend/handlers"

	"github.com/labstack/echo"
)

// func handleFunc(c echo.Context) error {
// 	ln, _ := net.Listen("tcp", ":1234") // listen on all interfaces
// 	conn, _ := ln.Accept()
// 	conn.SetDeadline(time.Now().Add(3 * time.Second))

// 	// time.AfterFunc(5*time.Second, func() {
// 	// 	fmt.Println("Shutting down server...")
// 	// 	conn.Close()
// 	// })
// 	message, _ := bufio.NewReader(conn).ReadString('\n') // will listen for message to process ending in newline (\n)
// 	fmt.Print("Message Received:", string(message))
// 	if message != "" {
// 		return c.String(http.StatusOK, "checked")
// 	}
// 	return c.String(http.StatusOK, "not receive message")
// }

func Private(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheckHandler)
	// e.GET("/check", handleFunc)
	userGr := e.Group("/users")
	userGr.POST("/signin", handlers.SignIn)
	userGr.POST("/signout", handlers.SignIn)

	e.POST("/form", handlers.OptionFormHandler2)
	// e.POST("ticket", )

	e.GET("/subscriptions", handlers.GetAllSubscriptions)
	// userGr.GET("/set-cookie", func(c echo.Context) error {
	// 	// Set a cookie
	// 	cookie := &http.Cookie{
	// 		Name:    "my_cookie",
	// 		Value:   "cookie_value",
	// 		Expires: time.Now().Add(10 * time.Second),
	// 		Path:    "/",
	// 	}
	// 	c.SetCookie(cookie)

	// 	return c.String(http.StatusOK, "Cookie set")
	// })

	// e.GET("/get-cookie", func(c echo.Context) error {
	// 	// Retrieve cookie from the request
	// 	cookie, err := c.Cookie("my_cookie")
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Get the value of the cookie
	// 	cookieValue := cookie.Value

	// 	return c.String(http.StatusOK, "Cookie value: "+cookieValue)
	// })
}
