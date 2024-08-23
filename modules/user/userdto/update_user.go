package userdto

type UpdateUserDTO struct {
	Name     *string `json:"name"`
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password" validate:"omitempty,min=4"`
}
