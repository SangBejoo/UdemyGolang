package handler

import (
	"strconv"

	"belajar-http/usecase"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	usecase usecase.TodoUsecase
}

func NewTodoHandler(usecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: usecase}
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	todo, err := h.usecase.CreateTodo(req.Title, req.Description)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(todo)
}

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	todos, err := h.usecase.GetAllTodos()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(todos)
}

func (h *TodoHandler) GetTodoByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	todo, err := h.usecase.GetTodoByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}
	return c.JSON(todo)
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	todo, err := h.usecase.UpdateTodo(uint(id), req.Title, req.Description, req.Completed)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(todo)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	err = h.usecase.DeleteTodo(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
