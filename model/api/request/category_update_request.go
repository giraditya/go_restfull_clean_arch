package request

type CategoryUpdateRequest struct {
	ID   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1"`
}
