package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"myApi/cmd/api/routes"
	"myApi/pkg/commons"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	address := ":8080"

	db := commons.InitDB()
	commons.InitValidate()

	commons.Migrate(db)

	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("⇨ myApi started at: http://localhost%s/swagger/index.html\n", address)
	}()

	// setting up echo logger
	logFormat := `time=${time_rfc3339_nano} level=${level} file=${short_file} line=${line}`
	log.SetHeader(logFormat)

	// setting up echo
	e := echo.New()
	e.Logger.SetHeader(logFormat)
	routes.Configure(e)
	e.Logger.Fatal(e.Start(address))
}
