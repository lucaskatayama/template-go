package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaskatayama/hexgo/internal/app"
	"net/http"
)

func ListTodos(c *fiber.Ctx) error {
	todos, err := app.Todo.List(c.UserContext())
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "v1 error")
	}
	return c.JSON(todos)
}
