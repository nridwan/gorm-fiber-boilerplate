package userdto

import (
	"gofiber-boilerplate/modules/user/usermodel"
)

type UserDTO = usermodel.ReadonlyUserModel

func MapUserModelToDTO(model *usermodel.UserModel) *UserDTO {
	return &UserDTO{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}
