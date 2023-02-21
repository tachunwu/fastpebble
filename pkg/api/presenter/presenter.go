package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tachunwu/fastpebble/pkg/entity"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value []byte `json:"value,omitempty"`
}

func SuccessResponse(data *[]entity.KeyValue) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
