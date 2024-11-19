package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var quizzCollection *mongo.Collection

func main() {
	app := fiber.New()

	setupDb()

	app.Get("/", index)
	app.Get("/api/quizzes", getQuizzes)

	log.Fatal(app.Listen(":3000"))
}

func setupDb() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}

	quizzCollection = client.Database("quiz").Collection("quizzes") 
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func getQuizzes(c *fiber.Ctx) error {
	cursor, err := quizzCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	quizzes := []map[string]any{}
	err = cursor.All(context.Background(), &quizzes)
	if err != nil {
		panic(err)
	}
	return c.JSON(quizzes)
}
