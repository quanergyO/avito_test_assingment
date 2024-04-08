package types

type Role int

const (
	User Role = iota + 1
	Admin
)

type UserType struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     Role   `json:"role,omitempty"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
