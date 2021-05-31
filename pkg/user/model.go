package user

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"myApi/pkg/commons"
	"myApi/pkg/commons/errors"
	"strings"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Active   bool
}

func (u *User) register() error {
	log.Infof("create user: %s", u.Username)
	return commons.DB.Create(u).Error
}

func (u *User) login() error {

	var dbu User
	err := commons.DB.
		Select("id", "password").
		Where("lower(username) = ? and active = true", strings.ToLower(u.Username)).
		First(&dbu).Error

	if err != nil {
		verifyFakeHash()
		return errors.Client("invalid_username_or_password", fmt.Sprintf("user not found: %s", u.Username))
	}

	if !verifyHash(dbu.Password, u.Password) {
		return errors.Client("invalid_username_or_password", u.Username)
	}

	u.ID = dbu.ID
	return nil
}

func findUserById(id uint) (*profileDTO, error) {

	var p profileDTO

	if err := commons.DB.Model(&User{}).Where("id = ? ", id).First(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}
