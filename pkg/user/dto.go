package user

import (
	"myApi/pkg/commons/errors"
)

type createUserDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (dto createUserDTO) validate() error {
	// TODO
	return nil
}

func (dto createUserDTO) alreadyExists() (bool, error) {
	// TODO
	return false, nil
}

// -----------------------------

type loginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (l loginDTO) validate() error {
	if l.Username == "" || l.Password == "" {
		return errors.New("username_or_password_validation_failed")
	}
	if len(l.Password) < 6 {
		return errors.New("password_validation_failed")
	}
	return nil
}
