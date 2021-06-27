package routes

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "myApi/docs"
	"myApi/pkg/health"
	"myApi/pkg/user"
	"myApi/pkg/user/jwt"
)

func Configure(e *echo.Echo) {

	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	prometheus.NewPrometheus("myApi", nil).Use(e)

	addPublicRoutes(e)
	addSecuredRoutes(e.Group("", middleware.JWT([]byte(jwt.Secret))))

	e.HTTPErrorHandler = buildHTTPErrorHandlerFunc(e)
}

func addPublicRoutes(e *echo.Echo) {

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/api/v1/health", health.Check)

	e.POST("/api/v1/login", user.Login)
	e.POST("/api/v1/register", user.RegisterUser)
}

func addSecuredRoutes(g *echo.Group) {

	g.GET("/api/v1/users/profile", user.Profile)

}
