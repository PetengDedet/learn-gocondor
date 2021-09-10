package input

type UserCreateRequestValidator struct {
	Username string `form:"username" json:"username" binding:"required,alphanum,min=3,max=255"`
	Name     string `form:"name" json:"name" binding:"required,alphanum,min=3,max=255"`
	Email    string `form:"email" json:"email" binding:"required,email,max=255"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=255"`
}
