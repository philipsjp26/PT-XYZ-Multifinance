package http

import (
	"go_playground/internal/core/port/service"

	"github.com/gofiber/fiber/v2"
)

type baseController struct {
	app *fiber.App
	srv service.LivenessService
}

func NewBaseController(app *fiber.App, service service.LivenessService) *baseController {
	return &baseController{
		app: app,
		srv: service,
	}
}

func (l *baseController) InitRouter() {

	api := l.app.Group("/api/")
	api.Get("liveness", func(c *fiber.Ctx) error {
		livenessResponse := l.srv.GetLiveness()
		return c.JSON(livenessResponse)
	})
}
