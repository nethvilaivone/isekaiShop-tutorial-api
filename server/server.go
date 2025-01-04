package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/neth/isekai-shop/config"
	"gorm.io/gorm"
)

type echoServer struct {
	app  *echo.Echo
	conf *config.Config
	db   *gorm.DB
}

var (
	once   sync.Once
	server *echoServer
)

func NewechoServer(conf *config.Config, db *gorm.DB) *echoServer {

	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{

			conf: conf,
			db:   db,
			app:  echoApp,
		}
	})

	return server
}

func (s *echoServer) Start() {

    getCoresmidleare := getCoresmidleare(s.conf.Server.AllowOrigins)
	getTimeOutmidleare := getTimeOutmidleare(s.conf.Server.Timeout)
	getBodyLimit := getBodyLimit(s.conf.Server.BodyLimit)
	
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(getCoresmidleare)
	s.app.Use(getTimeOutmidleare)
	s.app.Use(getBodyLimit)


	s.app.GET("/v1/health", s.healthcheck)

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go s.graceFullyShuttdown(ch)
	s.httpListening()
}

func (s *echoServer) healthcheck(c echo.Context) error {
	return c.String(http.StatusAccepted, "ok")
}

func (s *echoServer) httpListening() {
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)

	if err := s.app.Start(serverUrl); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatal("shuting down the server", err)
	}

}

func (s *echoServer) graceFullyShuttdown(quitCh chan os.Signal) {

	ctx := context.Background()
	<-quitCh
	s.app.Logger.Info("Received shutdown signal, waiting for 5 seconds...")
	time.Sleep(5 * time.Second) // Wait for 5 seconds
	s.app.Logger.Fatal("Shutting down...")
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatal("ERROR", err.Error())
	}

}


func getTimeOutmidleare(timeout time.Duration) echo.MiddlewareFunc {

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		Timeout:      timeout * time.Second,
		ErrorMessage: "Request timeout",
	})

}


func getCoresmidleare(allowOregin []string) echo.MiddlewareFunc{
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: allowOregin,
		AllowMethods: []string{echo.GET , echo.POST , echo.PUT , echo.DELETE , echo.PATCH},
		AllowHeaders: []string{echo.HeaderOrigin , echo.HeaderContentType , echo.HeaderAccept},
	})
	
}


func getBodyLimit(bodylomit string ) echo.MiddlewareFunc {

return middleware.BodyLimit(bodylomit)
}

//  enwbr stream Pluser ,
// io writer
// Redis sorted set ;
// go fan out is the coppy and send
// not use ctx in the function
// abstract or interface can not use pointer
// go context is use how can we cancel the go routine
// why it happend when you need to understand something , what it really is  ?
// return is mean to go of the function
// port value  = 6ຫມືນ
