package domain

type Order struct {
	ID       string
	Date     string
	Customer Customer
	BranchID Branch
}
