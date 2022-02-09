package request

type UserCreateRequest struct {
	Name     string `validate:"required,max=45,min=1" json:"name"`
	Address  string `validate:"required" json:"address"`
	Username string `validate:"required,max=45,min=1" json:"username"`
	Password string `validate:"required,max=45,min=1" json:"password"`
}
