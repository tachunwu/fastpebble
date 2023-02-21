package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tachunwu/fastpebble/pkg/fastpebble"
)

func NewRouter(r fiber.Router, service fastpebble.Service) {
	r.Get("/:key", Get(service))
	r.Post("/:key", Set(service))
	r.Delete("/:key", Delete(service))
	r.Post("/batchs", Batch(service))
}
