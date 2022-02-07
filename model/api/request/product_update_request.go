package request

type ProductUpdateRequest struct {
	ID         int    `validate:"required"`
	Name       string `validate:"required,max=200,min=1"`
	Size       string `validate:"required,max=5,min=1"`
	CategoryID int    `validate:"required"`
}
