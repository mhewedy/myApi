package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myApi/cmd/api/routes"
	"myApi/pkg/commons/dbutil"
	"myApi/pkg/profile"
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

	db := dbutil.New()
	dbutil.Migrate(db)
	injectDB(db)

	router := gin.Default()
	routes.Configure(router)

	go func() {
		time.Sleep(50 * time.Millisecond)
		println("\nmyApi started at: http://localhost:8080/swagger/index.html\n")
	}()

	_ = router.Run()
}

func injectDB(db *gorm.DB) {
	profile.DB = db
}
