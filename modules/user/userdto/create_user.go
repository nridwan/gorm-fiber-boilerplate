package userdto

import "gofiber-boilerplate/modules/user/usermodel"

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

func (dto *CreateUserDTO) ToModel() *usermodel.UserModel {
	return &usermodel.UserModel{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: &dto.Password,
	}
}
