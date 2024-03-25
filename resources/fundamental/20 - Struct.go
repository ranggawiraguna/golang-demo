package resources

import (
	"fmt"
	"golang-demo/utils"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	id   string
	name string
	age  int
}

type Product struct {
	id    string
	name  string
	stock int
}

type Order struct {
	id       string
	product  Product
	customer Customer
	quantity int
}

func printOrderDetail(order Order) {
	fmt.Printf("Customer Name\t: %s\n", order.customer.name)
	fmt.Printf("Customer Age\t: %d\n", order.customer.age)
	fmt.Printf("Product Name\t: %s\n", order.product.name)
	fmt.Printf("Order Quantity\t: %d\n", order.quantity)
}

func Struct() {
	var product Product
	var customer Customer
	var order Order

	utils.LineBreak()
	fmt.Println("Basic Struct")
	fmt.Println("")
	customer.id = uuid.NewString()
	customer.name = "Rangga Wiraguna"
	customer.age = time.Now().Year() - 2000
	product.id = uuid.NewString()
	product.name = "MacBook Pro"
	product.stock = 200
	order.id = uuid.NewString()
	order.product = product
	order.customer = customer
	fmt.Println(customer)
	fmt.Println(product)
	fmt.Println(order)
	utils.LineBreak()

	utils.LineBreak()
	fmt.Println("Pass Struct as Function Arguments")
	fmt.Println("")
	printOrderDetail(order)
	utils.LineBreak()
}
