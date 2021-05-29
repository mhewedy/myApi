package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"myApi/cmd/api/routes"
	"myApi/pkg/commons/dbutil"
	"myApi/pkg/health"
	"myApi/pkg/user"
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
func main() {
	address := ":8080"

	db := dbutil.New()
	dbutil.Migrate(db)
	injectDBRefs(db)

	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("⇨ myApi started at: http://localhost%s/swagger/index.html\n", address)
	}()

	// setting up echo
	logFormat := `time=${time_rfc3339_nano} level=${level} file=${short_file} line=${line}`
	log.SetHeader(logFormat)
	e := echo.New()
	e.Logger.SetHeader(logFormat)
	routes.Configure(e)
	e.Logger.Fatal(e.Start(address))
}

func injectDBRefs(db *gorm.DB) {
	health.DB = db
	user.DB = db
}
