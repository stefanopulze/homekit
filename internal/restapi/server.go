package restapi

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var server *fiber.App

func Start(opts *ServerOpts) {
	server = fiber.New()
	server.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("HomeKit Server")
	})

	server.Listen(fmt.Sprintf(":%d", opts.Port))
}

func Shutdown() {
	logrus.Debug("RestApi server shutting down")
	server.Shutdown()
	logrus.Info("RestApi server shutdown")
}
