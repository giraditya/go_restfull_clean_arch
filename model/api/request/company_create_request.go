package request

type CompanyCreateRequest struct {
	Name string `validate:"required,max=45,min=1" json:"name"`
	Code string `validate:"required,max=5,min=1" json:"code"`
}
