package domain

type Product struct {
	ID       int
	Name     string
	Size     string
	Price    int
	Category Category
}
