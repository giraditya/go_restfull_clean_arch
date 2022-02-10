package domain

import "time"

type Order struct {
	ID          string
	CreatedDate time.Time
	UpdatedDate time.Time
	Customer    Customer
	Branch      Branch
}
