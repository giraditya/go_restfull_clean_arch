package test

import (
	"fmt"
	"testing"
)

type PurchaseService interface {
	CreatePurchase() string
	RemovePurchase() string
}

type SalesService interface {
	CreateSales() string
	RemoveSales() string
}

type Item struct {
	Name string
	Qty  int
}

type Buyer struct {
	Name        string
	PhoneNumber int
}

type Seller struct {
	Name        string
	PhoneNumber int
}

type Purchase struct {
	ID     int
	Items  []Item
	Buyer  Buyer
	Seller Seller
}

type Sales struct {
	ID     int
	Items  []Item
	Buyer  Buyer
	Seller Seller
}

func (p Purchase) CreatePurchase() string {
	return "Successfull created order from " + p.Buyer.Name
}

func (p Purchase) RemovePurchase() string {
	return "Successfull remove order from " + p.Buyer.Name
}

func TestCreatePurchase(t *testing.T) {
	var purchaseService PurchaseService = Purchase{
		ID: 1,
		Items: []Item{
			{
				Name: "Cloth",
				Qty:  2,
			},
			{
				Name: "Jeans",
				Qty:  3,
			},
		},
		Buyer: Buyer{
			Name:        "Giri",
			PhoneNumber: 82117918363,
		},
		Seller: Seller{
			Name:        "Imox",
			PhoneNumber: 82393939992,
		},
	}

	fmt.Println(purchaseService.CreatePurchase())
}

func TestRemovePurchase(t *testing.T) {
	var purchaseService PurchaseService = Purchase{
		ID: 1,
		Items: []Item{
			{
				Name: "Cloth",
				Qty:  2,
			},
			{
				Name: "Jeans",
				Qty:  3,
			},
		},
		Buyer: Buyer{
			Name:        "Giri",
			PhoneNumber: 82117918363,
		},
		Seller: Seller{
			Name:        "Imox",
			PhoneNumber: 82393939992,
		},
	}

	fmt.Println(purchaseService.RemovePurchase())
}

func (s Sales) CreateSales() string {
	return "Succes create sales to " + s.Buyer.Name
}

func (s Sales) RemoveSales() string {
	return "Success remove sales from " + s.Seller.Name
}

func TestCreateSales(t *testing.T) {
	var salesService SalesService = Sales{
		ID: 1,
		Items: []Item{
			{
				Name: "Cloth",
				Qty:  2,
			},
			{
				Name: "Jeans",
				Qty:  3,
			},
		},
		Buyer: Buyer{
			Name:        "Giri",
			PhoneNumber: 82117918363,
		},
		Seller: Seller{
			Name:        "Imox",
			PhoneNumber: 82393939992,
		},
	}

	fmt.Println(salesService.CreateSales())
}

func TestRemoveSales(t *testing.T) {
	var salesService SalesService = Sales{
		ID: 1,
		Items: []Item{
			{
				Name: "Cloth",
				Qty:  2,
			},
			{
				Name: "Jeans",
				Qty:  3,
			},
		},
		Buyer: Buyer{
			Name:        "Giri",
			PhoneNumber: 82117918363,
		},
		Seller: Seller{
			Name:        "Imox",
			PhoneNumber: 82393939992,
		},
	}

	fmt.Println(salesService.RemoveSales())
}
