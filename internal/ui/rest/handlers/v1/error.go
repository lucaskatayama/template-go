package v1

import "github.com/gofiber/fiber/v2"

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "unknown error"

	switch t := err.(type) {
	case *fiber.Error:
		code = t.Code
		msg = t.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"msg >>>>>>>>": msg,
	})
}
