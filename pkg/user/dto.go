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

func (dto createUserDTO) checkAlreadyExists() error {

	var count int64
	err := DB.
		Model(&User{}).
		Where("username = ?", dto.Username).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.Client("username_already_exists", dto.Username)
	}

	return nil
}

// --------------------------------------------------------

type loginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
