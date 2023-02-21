package transport

import (
	"github.com/gofiber/fiber/v2"
)

func NewHTTPServer() Transport {
	return new(HTTPServer)
}

type HTTPServer struct {
	server *fiber.App
	addr   string
}

func (s *HTTPServer) Init(addr string, server *fiber.App) {
	s.addr = addr
	s.server = server
}

func (s *HTTPServer) Serve() {

	s.server.Listen(s.addr)

}

func (s *HTTPServer) Close() {
	s.server.Shutdown()
}
