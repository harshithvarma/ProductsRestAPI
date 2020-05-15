package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

//ErrProductNotFound is an error raised when a product cannot be found in the database
var ErrProductNotFound = fmt.Errorf("product not found")

//Product defines the structure for an API
type Product struct {
	ID          int		`json:"id"`
	Name        string	`json:"name" validate:"required"`
	Description string	`json:"description"`
	Price       int		`json:"price" validate:"required,gt=0"`
}

//ProductsList is a slice of product struct
type ProductsList []*Product

func(p *Product) Validate() error{
	validate:=validator.New()
	return validate.Struct(p)
}

//GetProductsList returns a slice of all the products in ProductsList
func GetProductsList() ProductsList {
	return productsList
}

//AddProduct appends the product from request into our ProductsList
func AddProduct(p *Product) {
	id := GetId()
	p.ID = id
	productsList = append(productsList, p)
}

//UpdateProduct updates the productList with given product from the request body
func UpdateProduct(id int, p *Product) error {
	_,pos,err:=findProduct(id)
	if err!=nil{
		return err
	}

	productsList[pos]=p
	return nil
}

//getId is used to retrieve the latest ID entry in the ProductsList
func GetId() int {
	n := productsList[len(productsList)-1]
	return n.ID + 1
}

//findProduct is used to find the product with ID as search parameter
func findProduct(id int) (*Product, int, error) {
	for i, p := range productsList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil,-1,ErrProductNotFound
}

//productsList is false database with manual entries into the slice of product
var productsList = []*Product{
	&Product{
		ID:          1,
		Name:        "tea",
		Description: "tea powder from imported from china",
		Price:       100,
	},
	&Product{
		ID:          2,
		Name:        "coffee",
		Description: "Made in ooty",
		Price:       200,
	},
	&Product{
		ID:          3,
		Name:        "osmania biscuits",
		Description: "Made in Hyderabad",
		Price:       300,
	},
}
