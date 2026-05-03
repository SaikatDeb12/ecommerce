package models

type SignUpModel struct {
	Name     string `json:"name" validate:"required,min=3,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=15"`
	Address  string `json:"address" validate:"required,max=20"`
	Role     string `json:"role" validate:"required,oneof=admin customer delivery_agent"`
}

type SignInModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=15"`
}

type ErrorModel struct {
	Error      string
	StatusCode int
	Message    string
}
