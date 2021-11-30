package v1

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler fiber.Handler
}

var Routes = []Route{
	{
		Path:    "/todo",
		Method:  http.MethodGet,
		Handler: ListTodos,
	},
}

func Router() *fiber.App {
	api := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})

	for _, route := range Routes {
		api.Add(route.Method, route.Path, route.Handler)
	}
	return api
}
