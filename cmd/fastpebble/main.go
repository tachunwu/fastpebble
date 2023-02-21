package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tachunwu/fastpebble/pkg/api"
	"github.com/tachunwu/fastpebble/pkg/fastpebble"
	"github.com/tachunwu/fastpebble/pkg/transport"
)

func main() {
	// Make sure we clean up before exiting.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
	NewApp()

	for {
		runtime.Gosched()
	}

}

type App struct {
	repository fastpebble.Repository
	transport  transport.Transport
	service    fastpebble.Service
}

func NewApp() *App {

	t := transport.NewHTTPServer()
	repo := fastpebble.NewRepository()
	svc := fastpebble.NewService(repo)

	server := fiber.New()
	server.Use(cors.New())
	r := server.Group("/v1")
	api.NewRouter(r, svc)

	t.Init(":34567", server)
	t.Serve()

	app := &App{
		transport:  t,
		repository: repo,
		service:    svc,
	}

	return app
}
