// Package app contains the entire structure of the application, it describes the functions of
// initializing the application components and its launch. And also this package contains all the handlers
// that are responsible for working with the request and getting data from it
package app

import (
	"avito/configs"
	"avito/internal/parser"
	"avito/internal/repository"
	"avito/internal/service"
	"avito/pkg/logger"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"syscall"
)

// App -  object responsible for working with the request and extracting data from it
type App struct {
	config        *configs.Common
	defaultServer *fasthttp.Server
	services      *service.Service
	parser        *parser.Parser
	logger        logger.Logger
	store         *reportStore
}

// NewApp - constructor function for App
func NewApp() *App {
	return &App{
		defaultServer: &fasthttp.Server{},
		parser:        parser.NewParser(),
		logger:        logger.New(),
	}
}

// ParseConfig - function for parsing config from env
func (a *App) ParseConfig() {
	if err := godotenv.Load(".env"); err != nil {
		a.logger.Fatalf("cannot load env: %s", err.Error())
	}
	var cfg configs.Common
	if err := env.Parse(&cfg); err != nil {
		a.logger.Fatalf("cannot parse config: %s", err.Error())
	}
	a.config = &cfg
}

// InitApi -  initializes the application
func InitApi() {
	a := NewApp()
	a.ParseConfig()
	repo, err := repository.NewPostgresDB(a.config.ConfigDB)
	if err != nil {
		a.logger.Fatalf("init db error: %s", err.Error())
	}
	r := repository.NewRepository(repo)
	s := service.NewService(r)
	a.services = s
	storage := reportStore{
		reports: make(map[string][]byte),
	}
	a.store = &storage
	router := a.Routing()
	a.defaultServer.Handler = router.Handler
	a.Run()
	repo.Close()
}

func (a *App) Run() {
	a.logger.Info("start service")
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		err := a.defaultServer.ListenAndServe(fmt.Sprintf("%s:%d", a.config.ServiceHost, a.config.ServicePort))
		if err != nil {
			a.logger.Fatalf("problem with run app: %s", err.Error())
		}
	}()
	s := <-signalChannel
	a.logger.Infof("Got signal: %s. Initiate gracefully stop.\n", s.String())
	if err := a.defaultServer.Shutdown(); err != nil {
		a.logger.Fatal(err)
	}
}
