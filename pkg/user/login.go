package user

import (
	"github.com/labstack/echo/v4"
	"myApi/pkg/user/jwt"
	"net/http"
)

//
// @Summary Login
// @Description
// @Accept json
// @Param  profile body user.loginDTO true "Login"
// @Success 200
// @Router /api/login [post]
func Login(c echo.Context) error {

	var l loginDTO
	if err := c.Bind(&l); err != nil {
		return err
	}
	if err := l.validate(); err != nil {
		return err
	}

	u := &User{
		Username: l.Username,
		Password: l.Password,
	}

	if err := u.login(); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	token, err := jwt.CreateToken(u.ID, u.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
