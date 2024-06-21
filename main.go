package main

import (
	"bytes"
	"firstApi/models"
	"firstApi/routes"
	"firstApi/util"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)  

 
func main() {
	_ = godotenv.Load(".env")

	e := echo.New()

	e.Validator = &util.Validator{Instance: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:`${time_rfc3339} | ${method} ${uri} | ${status} | ${custom} | ${remote_ip} | ${user_agent}` + "\n", 
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
	
	log.SetLevel(log.DebugLevel)
	 
	db:= models.ConnectDatabase("goDB")
	models.Migrate(db)
	
	e.Use(middleware.Recover()) 
	
	api := e.Group("/api")
	
	routes.SetupRoute(api)
	  
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
} 
 