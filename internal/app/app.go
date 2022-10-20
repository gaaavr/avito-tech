package app

import (
	"avito/internal/parser"
	"avito/internal/service"
	"avito/pkg/logger"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

// App -  object responsible for working with the request and extracting data from it
type App struct {
	defaultServer *fasthttp.Server
	services      *service.Service
	parser        *parser.Parser
	logger        logger.Logger
}

// New - constructor function for App
func NewApp(services *service.Service) *App {
	return &App{
		defaultServer: &fasthttp.Server{},
		services:      services,
		parser:        parser.NewParser(),
		logger:        logger.New(),
	}
}

// Routing - method for registering all handlers in the router
func (a *App) Routing() *router.Router {
	router := router.New()
	router.POST("/accrual_funds", a.LogRequests(a.accrualFunds))
	router.POST("/get_balance", a.LogRequests(a.getBalance))
	router.POST("/create_order", a.LogRequests(a.blockFunds))
	router.POST("/charge_funds", a.LogRequests(a.chargeFunds))
	return router
}

func (a *App) Run() {
	router := a.Routing()

	a.defaultServer.Handler = router.Handler
	a.logger.Info("start service")
	osSignals := make(chan os.Signal, 1)
	go func() {
		err := a.defaultServer.ListenAndServe(fmt.Sprintf("%s:%d", "localhost", 8080))
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-osSignals
	if err := a.defaultServer.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
