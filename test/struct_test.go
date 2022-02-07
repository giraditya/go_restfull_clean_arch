package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Items struct {
	Name    string
	Size    string
	Price   int
	Qty     int
	Details ItemDetails
}

type ItemDetails struct {
	Brand    string
	Supplier Supplier
}

type Supplier struct {
	Name        string
	PhoneNumber int
}

type Order struct {
	ID          int
	OrderNumber string
	Items       []Items
}

func (o *Order) CreateOrder() {
	result, _ := json.Marshal(o)
	fmt.Println(string(result))
}

func TestCreateOrder(t *testing.T) {
	order := Order{
		ID:          1,
		OrderNumber: "OR001",
		Items: []Items{
			{
				Name:  "Baju",
				Size:  "M",
				Price: 20000,
				Qty:   2,
				Details: ItemDetails{
					Brand: "Erigo",
					Supplier: Supplier{
						Name:        "Giri",
						PhoneNumber: 82117918363,
					},
				},
			},
			{
				Name:  "Celana",
				Size:  "M",
				Price: 30000,
				Qty:   3,
				Details: ItemDetails{
					Brand: "Proshop",
					Supplier: Supplier{
						Name:        "Rafi",
						PhoneNumber: 82383992233,
					},
				},
			},
		},
	}
	order.CreateOrder()
}
