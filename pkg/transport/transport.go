package transport

import "github.com/gofiber/fiber/v2"

type Transport interface {
	Init(addr string, server *fiber.App)
	Serve()
	Close()
}
