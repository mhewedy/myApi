package user

import (
	"github.com/go-playground/validator"
	"myApi/pkg/commons/errors"
)

var (
	Validate *validator.Validate
)

// --------------------------------------------------------

type createUserDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

func (dto createUserDTO) alreadyExists() error {
	// TODO
	//if ok {
	//	return errors.Client("username_already_exists", dto.Username)
	//}
	return nil
}

// --------------------------------------------------------

type loginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (l loginDTO) validate() error {
	if l.Username == "" || l.Password == "" {
		return errors.Client("username_or_password_validation_failed")
	}
	if len(l.Password) < 6 {
		return errors.Client("password_validation_failed")
	}
	return nil
}
