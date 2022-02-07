package response

type ProductResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Size     string `json:"size"`
	Category string `json:"category"`
}
