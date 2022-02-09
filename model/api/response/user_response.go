package response

type UserResponse struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Company  string `json:"company"`
}
