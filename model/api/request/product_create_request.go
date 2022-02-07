package request

type ProductCreateRequest struct {
	Name       string `validate:"required,max=200,min=1" json:"name"`
	Size       string `validate:"required,max=5,min=1" json:"size"`
	CategoryID int    `validate:"required" json:"categoryID"`
}
