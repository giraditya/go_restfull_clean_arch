package request

type UserGenerateAuthKeyRequest struct {
	Username string `validate:"required,max=45,min=1" json:"username"`
	Password string `validate:"required,max=45,min=1" json:"password"`
}
