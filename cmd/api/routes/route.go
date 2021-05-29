package routes

import (
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

	addPublicRoutes(e)
	addSecuredRoutes(e.Group("", middleware.JWT([]byte(jwt.Secret))))

	e.HTTPErrorHandler = buildHTTPErrorHandlerFunc(e)
}

func addPublicRoutes(e *echo.Echo) {

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/api/health", health.Check)

	e.POST("/api/login", user.Login)
	e.POST("/api/register", user.RegisterUser)
}

func addSecuredRoutes(g *echo.Group) {

	g.GET("/api/users/profile", user.Profile)

}
