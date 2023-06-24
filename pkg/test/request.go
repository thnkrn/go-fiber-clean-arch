package test

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	middleware "github.com/thnkrn/go-fiber-clean-arch/pkg/api/middleware"
)

type HTTPRequest struct {
	Method      string
	Path        string
	Body        interface{}
	Description string
}

type errortDependencies struct {
}

func createErrorHandler() (*middleware.ErrorHandler, *errortDependencies) {
	errorHandler := middleware.NewErrorHandler()

	return errorHandler, &errortDependencies{}
}

func RequestHandler(urlPattern string, request *http.Request, handler fiber.Handler) *http.Response {
	errorHandler, _ := createErrorHandler()

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler.FiberErrorHandler(),
	})

	switch request.Method {
	case http.MethodPost:
		app.Post(urlPattern, handler)
	case http.MethodPut:
		app.Put(urlPattern, handler)
	case http.MethodDelete:
		app.Delete(urlPattern, handler)
	default:
		app.Get(urlPattern, handler)
	}

	response, _ := app.Test(request, -1)

	return response
}
