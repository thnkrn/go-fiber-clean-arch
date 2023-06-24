package api

import (
	"fmt"
	"log"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	handler "github.com/thnkrn/go-fiber-clean-arch/pkg/api/handler"
	middleware "github.com/thnkrn/go-fiber-clean-arch/pkg/api/middleware"
	config "github.com/thnkrn/go-fiber-clean-arch/pkg/config"
)

type Middlewares struct {
	ErrorHandler   *middleware.ErrorHandler
	Authentication *middleware.Authentication
}

type Handlers struct {
	UserHandler *handler.UserHandler
}

type ServerHTTP struct {
	app *fiber.App
}

func NewServerHTTP(middlewares *Middlewares, handlers Handlers, cfg config.Config) *ServerHTTP {
	app := fiber.New(
		fiber.Config{
			// NOTE: enable SO_REUSEPORT,
			// https://pkg.go.dev/github.com/valyala/fasthttp/reuseport, https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/, https://github.com/gofiber/fiber/issues/180
			Prefork: cfg.Prefork,
			// NOTE: Override default JSON encoding, ref: https://docs.gofiber.io/guide/faster-fiber#custom-json-encoder-decoder
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
			// NOTE: Override default error handler
			ErrorHandler: middlewares.ErrorHandler.FiberErrorHandler(),
		})

	log.Printf((fmt.Sprintf("Server is started with PID: %v and PPID: %v", os.Getpid(), os.Getppid())))

	// NOTE: Enable log tracing from Fiber, https://docs.gofiber.io/api/middleware/logger
	if cfg.Tracing {
		app.Use(logger.New())
	}

	if cfg.Recover {
		app.Use(recover.New())
	}

	// NOTE: Healthcheck
	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/login", middleware.LoginHandler)

	userAPI := app.Group("/api", middlewares.Authentication.Authentication())
	userAPI.Get("users", handlers.UserHandler.FindAll)
	userAPI.Get("users/:id<minLen(1)>", handlers.UserHandler.FindByID)
	userAPI.Post("users", handlers.UserHandler.Create)
	userAPI.Delete("users/:id<minLen(1)>", handlers.UserHandler.Delete)
	userAPI.Put("users/:id<minLen(1)>", handlers.UserHandler.Update)
	userAPI.Get("users/name/:text<minLen(1)>", handlers.UserHandler.FindByMatchName)

	return &ServerHTTP{app}
}

func (sh *ServerHTTP) Start() {
	sh.app.Listen(":8080")
}
