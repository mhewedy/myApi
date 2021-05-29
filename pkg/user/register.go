package user

import (
	"github.com/labstack/echo/v4"
	"myApi/pkg/commons/errors"
	"net/http"
)

//
// @Summary Register new user
// @Description
// @Accept json
// @Param  profile body user.createUserDTO true "Register user"
// @Success 200
// @Router /api/register [post]
func RegisterUser(c echo.Context) error {

	var dto createUserDTO
	if err := c.Bind(&dto); err != nil {
		return err
	}
	if err := dto.validate(); err != nil {
		return err
	}

	ok, err := dto.alreadyExists()
	if err != nil {
		return err
	}
	if ok {
		return errors.Client("username_already_exists", dto.Username)
	}

	user := &User{
		Username: dto.Username,
		Password: hash(dto.Password),
		Active:   true,
	}

	if err := user.register(); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
