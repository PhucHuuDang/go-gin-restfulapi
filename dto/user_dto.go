package dto

type CreateUser struct {
	Name      string  `json:"name" example:"John"`
	Email     string  `json:"email" example:"john@gmail.com"`
	Age       int     `json:"age" example:"18"`
	AvatarURL *string `json:"avatar_url" example:"https://avatar.com/a.png"`
}
type UpdateUserDTO struct {
	Name      string  `json:"name" example:"John"`
	Email     string  `json:"email" example:"john@gmail.com"`
	Age       int     `json:"age" example:"18"`
	AvatarURL *string `json:"avatar_url"`
}
type DeleteUserDTO struct {
	ID uint `json:"id" binding:"required"`
}
