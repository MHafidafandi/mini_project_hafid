package request

type UserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password" validate:"required"`
}
