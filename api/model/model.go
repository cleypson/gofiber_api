package model

import (
	"gofiber_api/api/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Author struct
type Author struct {
	gorm.Model
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	email     string `json: "email"`
	Password  string `json: "password"`
}

func GetAuthors(c *fiber.Ctx) {
	db := database.DBConn
	var authors []Author
	db.Find(&authors)
	c.JSON(authors)
}

func GetAuthor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var author Author
	db.Find(&author, id)
	c.JSON(author)
}

func NewAuthor(c *fiber.Ctx) {
	db := database.DBConn
	author := new(Author)
	if err := c.QueryParser(author); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&author)
	c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var author Author
	db.First(&author, id)
	db.Delete(&author)
	c.Send("Author Successfully deleted")
}

// Topic struct
type Topic struct {
	gorm.Model
	Title  string `json: "title"`
	Author string `json: "author"`
}

func GetTopics(c *fiber.Ctx) {
	db := database.DBConn
	var topics []Topic
	db.Find(&topics)
	c.JSON(topics)
}

func GetTopic(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var topic Topic
	db.Find(&topic, id)
	c.JSON(topic)
}

func NewTopic(c *fiber.Ctx) {
	db := database.DBConn
	topic := new(Topic)
	if err := c.QueryParser(topic); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&topic)
	c.JSON(topic)
}

func DeleteTopic(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var topic Topic
	db.First(&topic, id)
	db.Delete(&topic)
	c.Send("Topic Successfully deleted")
}

// Commentary struct
type Commentary struct {
	gorm.Model
	Text       string      `json: "Text"`
	Author     string      `json: "author"`
	Topic      *Topic      `json: "topic_id"`
	Commentary *Commentary `json: "commentary_id"`
}

func GetCommentaries(c *fiber.Ctx) {
	db := database.DBConn
	var commentaries []Commentary
	db.Find(&commentaries)
	c.JSON(commentaries)
}

func GetCommentary(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var commentary Commentary
	db.Find(&commentary, id)
	c.JSON(commentary)
}

func NewCommentary(c *fiber.Ctx) {
	db := database.DBConn
	commentary := new(Commentary)
	if err := c.QueryParser(commentary); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&commentary)
	c.JSON(commentary)
}

func DeleteCommentary(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var commentary Commentary
	db.First(&commentary, id)
	db.Delete(&commentary)
	c.Send("Commentary Successfully deleted")
}

// Topic struct
type Rating struct {
	gorm.Model
	Topic      *Topic      `json: "topic_id"`
	Commentary *Commentary `json: "commentary_id"`
	Author     string      `json: "author_id"`
	Rating     int         `json:"rating"`
}

func GetRatings(c *fiber.Ctx) {
	db := database.DBConn
	var ratings []Rating
	db.Find(&ratings)
	c.JSON(ratings)
}

func GetRating(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var rating Rating
	db.Find(&rating, id)
	c.JSON(rating)
}

func NewRating(c *fiber.Ctx) {
	db := database.DBConn
	var rating Rating
	if err := c.QueryParser(rating); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&rating)
	c.JSON(rating)
}

func DeleteRating(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var rating Rating
	db.First(&rating, id)

	db.Delete(&rating)
	c.Send("Rating Successfully deleted")
}
