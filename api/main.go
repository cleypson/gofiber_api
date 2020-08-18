package main

import (
	"fmt"
	"gofiber_api/api/database"
	"gofiber_api/api/model"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func index(c *fiber.Ctx) {
	c.SendFile("dist/index.html")
}

func setupRoutes(app *fiber.App) {
	// index
	app.Get("/", index)
	// authors
	app.Get("/api/v1/author", model.GetAuthors)
	app.Get("/api/v1/author/:id", model.GetAuthor)
	app.Post("/api/v1/author", model.NewAuthor)
	app.Delete("/api/v1/author/:id", model.DeleteAuthor)
	// topics
	app.Get("/api/v1/topic", model.GetTopics)
	app.Get("/api/v1/topic/:id", model.GetTopic)
	app.Post("/api/v1/topic", model.NewTopic)
	app.Delete("/api/v1/topic/:id", model.DeleteTopic)
	// commentaries
	app.Get("/api/v1/commentary", model.GetCommentaries)
	app.Get("/api/v1/commentary/:id", model.GetCommentary)
	app.Post("/api/v1/commentary", model.NewCommentary)
	app.Delete("/api/v1/commentary/:id", model.DeleteCommentary)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "sqlite3.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&model.Author{})
	database.DBConn.AutoMigrate(&model.Topic{})
	database.DBConn.AutoMigrate(&model.Commentary{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Static("/", "dist")

	setupRoutes(app)

	app.Listen(8000)
	defer database.DBConn.Close()
}
