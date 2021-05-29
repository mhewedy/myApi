package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mhewedy/go-conf"
	"time"
)

var Secret = conf.Get("jwt.secret", "v3ryS3cureP@ss")

func CreateToken(id uint, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = username
	claims["exp"] = time.Now().AddDate(10, 0, 0).Unix() // expires in 10 years
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(Secret))
}

func GetUser(c echo.Context) (id uint, username string) {
	user := c.Get("user").(*jwt.Token)
	log.Infof("access token: %s", user.Raw)
	claims := user.Claims.(jwt.MapClaims)

	id = uint(claims["id"].(float64))
	username = claims["name"].(string)
	return
}
