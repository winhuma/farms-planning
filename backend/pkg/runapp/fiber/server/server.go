package server

import (
	"farms-planning/pkg/runapp/fiber/middleware"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
)

func New() *fiber.App {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(middleware.FiberLog())
	return app
}

func Run(app *fiber.App, port string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	serverShutdown := make(chan struct{})

	go func() {
		<-c
		log.Info().Msg("...Gracefully shutting down...")
		app.Shutdown()
		serverShutdown <- struct{}{}
	}()

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal().Msgf("%s", err)
	}
	<-serverShutdown
	log.Info().Msg("...Running cleanup tasks...")
}
