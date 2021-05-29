package user

import (
	"github.com/labstack/echo/v4"
	"myApi/pkg/user/jwt"
)

//
// @Summary Get user profile
// @Description
// @Produce json
// @Success 200 {object} user.profileDTO "User Profile"
// @Router /users/profile [get]
// @Security ApiKeyAuth
func Profile(c echo.Context) error {

	id, _ := jwt.GetUser(c)

	user, err := findUserById(id)

	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
