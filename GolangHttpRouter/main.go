package main

import (
	"belajar-http/entity"
	"belajar-http/handler"
	"belajar-http/repository"
	"belajar-http/usecase"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres" // changed from sqlite
	"gorm.io/gorm"
)

func main() {
	// initialize database
	dsn := "host=localhost user=postgre password=postgre dbname=todos port=5432 sslmode=disable" // adjust as needed
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})                                     // changed from sqlite
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to connect database: %v", err))
	}

	// migrate the schema
	db.AutoMigrate(&entity.Todo{})

	// wire up layers
	todoRepo := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	// todo routes
	app.Post("/todos", todoHandler.CreateTodo)
	app.Get("/todos", todoHandler.GetAllTodos)
	app.Get("/todos/:id", todoHandler.GetTodoByID)
	app.Put("/todos/:id", todoHandler.UpdateTodo)
	app.Delete("/todos/:id", todoHandler.DeleteTodo)

	// start server

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to start server: %v", err))
	}
}
