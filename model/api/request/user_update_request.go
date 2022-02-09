package request

type UserUpdateRequest struct {
	ID       int    `validate:"required" json:"id"`
	Name     string `validate:"required,max=45,min=1" json:"name"`
	Address  string `validate:"required" json:"address"`
	Username string `validate:"required,max=45,min=1" json:"username"`
	Password string `validate:"required,max=45,min=1" json:"password"`
}
