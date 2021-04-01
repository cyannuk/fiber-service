package api

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"fiber-service/config"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog/log"
	"gopkg.in/reform.v1"

	"fiber-service/api/errors"
	"fiber-service/interface/service"
)

type application struct {
	*fiber.App
	userService  service.UserService
	orderService service.OrderService
	sessionStore *session.Store
	config       config.ServerConfig
}

type Application interface {
	Start() error
}

func (application *application) shutdown() {
	err := application.Shutdown()
	if err != nil {
		log.Error().Err(err).Msg("Application")
	}
	err = application.sessionStore.Reset()
	if err != nil {
		log.Error().Err(err).Msg("Application")
	}
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	switch err {
	case errors.ErrNoParameter:
	case errors.ErrInvalidParameter:
		err = fiber.ErrBadRequest
	case reform.ErrNoRows:
		err = fiber.ErrNotFound
	}
	return fiber.DefaultErrorHandler(ctx, err)
}

func (application *application) Start() error {
	certFile, keyFile := application.config.Certificate()
	if len(certFile) > 0 {
		return application.ListenTLS(application.config.BindAddress(), certFile, keyFile)
	} else {
		return application.Listen(application.config.BindAddress())
	}
}

func NewApplication(userService service.UserService, orderService service.OrderService, config config.ServerConfig) Application {
	application := application{fiber.New(fiber.Config{ErrorHandler: errorHandler, JSONEncoder: json.Marshal}), userService, orderService,
		session.New(session.Config{KeyLookup: "cookie:X-SESSION-ID"}), config}

	application.Post("/login", application.Login)
	application.Get("/users", application.GetUsers)
	application.Post("/users", application.CreateUser)
	application.Get("/users/me", application.GetUser)
	application.Get("/users/:user_id", application.GetUserById)
	application.Delete("/users/:user_id", application.DeleteUserById)
	application.Get("/users/:user_id/orders", application.GetOrders)
	application.Get("/users/:user_id/orders/:order_id", application.GetOrder)
	application.Get("/user_orders", application.GetUserOrders)

	once := sync.Once{}
	channel := make(chan os.Signal, 3)
	signal.Notify(channel, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		<-channel
		once.Do(application.shutdown)
	}()

	return &application
}
