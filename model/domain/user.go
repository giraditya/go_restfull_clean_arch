package domain

type User struct {
	ID       int
	Name     string
	Address  string
	Username string
	Password string
	AuthKey  string
	Company  Company
}
