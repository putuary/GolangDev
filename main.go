package main

import (
	"chat-messaging/configs"
	"chat-messaging/handlers"
	"chat-messaging/migrations"

	"github.com/gin-gonic/gin"

	_ "net/http/pprof"
)

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Migrate(db)
	
	r := gin.Default()
	r.GET("/users", handlers.GetUsers(db))
	r.POST("/users", handlers.CreateUser(db))

	r.GET("/messages", handlers.GetMessages(db))
	r.POST("/messages", handlers.CreateMessage(db))
	r.PUT("/messages/:id", handlers.UpdateMessage(db))
	r.DELETE("/messages/:id", handlers.DeleteMessage(db))

	r.GET("/groups", handlers.GetGroups(db))
	r.POST("/groups", handlers.CreateGroup(db))
	r.PUT("/groups/:id", handlers.UpdateGroup(db))
	r.DELETE("/groups/:id", handlers.DeleteGroup(db))
	
	r.Run(":8081")
}