package presenter

import "github.com/gofiber/fiber/v2"

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func SuccessResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func CreatedResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "created",
		"error":  nil,
	}
}
