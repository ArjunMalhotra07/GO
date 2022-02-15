package main

import "fmt"

type Supplier interface {
	PlaceOrder()
}

type SupplierOne struct{}

func (s *SupplierOne) PlaceOrder() {
	fmt.Println("Sent email to place Order")
}

type SupplierTwo struct{}

func (s *SupplierTwo) PlaceOrder() {
	fmt.Println("Sent api request to place Order")
}
func main() {
	supplierOne := &SupplierOne{}
	supplierTwo := &SupplierTwo{}

	//supplierOne.PlaceOrder()
	//supplierTwo.PlaceOrder()
	OrderPlacer(supplierOne, supplierTwo)

}

func OrderPlacer(suppliers ...Supplier) {
	for _, supplier := range suppliers {
		supplier.PlaceOrder()
	}
}
