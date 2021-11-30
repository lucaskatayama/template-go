package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	v1 "github.com/lucaskatayama/hexgo/internal/ui/rest/handlers/v1"
	"github.com/lucaskatayama/hexgo/pkg/log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	app := fiber.New(fiber.Config{
		Prefork:   false,
		Immutable: true,
		AppName:   "hexgo",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(requestid.New())

	app.Mount("/v1", v1.Router())

	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		if err := app.Listen(net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT"))); err != nil {
			log.Panicf("starting http server: %+v\n", err)
		}
	}()

	<-done

	if err := app.Shutdown(); err != nil {
		log.Errorf("shutting down server: %+v", err)
	}

	log.Infof("server stopped")
}
