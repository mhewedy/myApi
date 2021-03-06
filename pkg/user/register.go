package user

import (
	"github.com/labstack/echo/v4"
	"myApi/pkg/commons"
	"net/http"
)

//
// @Summary Register new user
// @Description
// @Accept json
// @Param  profile body user.createUserDTO true "Register user"
// @Success 200
// @Router /register [post]
// @Tags user
func RegisterUser(c echo.Context) error {

	var dto createUserDTO
	if err := c.Bind(&dto); err != nil {
		return err
	}

	if err := commons.Validate.Struct(dto); err != nil {
		return err
	}

	if err := dto.checkAlreadyExists(); err != nil {
		return err
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
