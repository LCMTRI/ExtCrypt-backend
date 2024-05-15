package main

import (
	"go-workspace/ext-crypt-backend/database"
	"go-workspace/ext-crypt-backend/repository"
	"go-workspace/ext-crypt-backend/route"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database.ConnectDb()
	repository.GetSubscriptionsColl()
	repository.GetUsersColl()
	e := echo.New()
	e.Use(middleware.CORS())
	route.Private(e)

	e.Logger.Fatal(e.Start(":8000"))
}

// func main() {
// 	ln, _ := net.Listen("tcp", ":1234") // listen on all interfaces
// 	conn, _ := ln.Accept()
// 	// conn.SetDeadline(time.Now().Add(3 * time.Second))

// 	time.AfterFunc(5*time.Second, func() {
// 		fmt.Println("Shutting down server...")
// 		conn.Close()
// 	})
// 	message, _ := bufio.NewReader(conn).ReadString('\n') // will listen for message to process ending in newline (\n)
// 	fmt.Print("Message Received:", string(message))
// }
