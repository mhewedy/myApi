package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CreateUserDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email"`
}

//
// @Summary Add a new user
// @Description get string by ID
// @Accept json
// @Param  profile body profile.CreateUserDTO true "Create profile"
// @Success 200
// @Router /profiles [post]
func Create(c *gin.Context) (interface{}, error) {

	var userDTO CreateUserDTO

	err := c.BindJSON(&userDTO)
	if err != nil {
		return nil, err
	}

	if err := create(&userDTO); err != nil {
		return nil, err
	}

	return nil, nil
}

func create(userDTO *CreateUserDTO) error {

	logrus.Infof("creating user: %+v", userDTO)

	user := User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}

	return DB.Create(&user).Error
}
