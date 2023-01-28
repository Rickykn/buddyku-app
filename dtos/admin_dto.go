package dtos

type AdminRegisterDTO struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AdminRegisterResponse struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
