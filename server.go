package main

import (
	"bytes"
	"firstApi/repository"
	"firstApi/routes"
	"firstApi/util"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct{}

// NewServer creates a new instance of Server.
func NewServer() *Server {
	return &Server{}
}

// ListenAndServe represents the main entry point of the program.
//
// Sets up DB connection logic, registers routes and listens for connections.
func (s *Server) ListenAndServe() error {
	config := util.NewConfig()
	e := echo.New()

	e.Validator = &util.Validator{Instance: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	setupMiddlewares(e)

	db, err := repository.ConnectDatabase("goDB")
	if err != nil {
		return err
	}
	repository.Migrate(db)

	api := e.Group("/api")

	routes.SetupRoute(api)

	return e.Start(":" + config.Port)
}

// setupMiddlewares sets up middlewares for logging, timming and recovering.
func setupMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `${time_rfc3339} | ${method} ${uri} | ${status} | ${custom} | ${remote_ip} | ${user_agent}` + "\n",
		CustomTimeFormat: time.RFC3339,
		CustomTagFunc: func(c echo.Context, buf *bytes.Buffer) (int, error) {
			start := c.Get("start_time").(time.Time)
			end := time.Now()
			latencyStr := util.CustomLatency(start, end)
			return buf.WriteString(latencyStr)
		},
		Output: nil,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("start_time", time.Now())
			return next(c)
		}
	})

	e.Use(middleware.Recover())
}
