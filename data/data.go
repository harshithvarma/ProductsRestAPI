package data

import (
	"encoding/json"
	"fmt"
	"io"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int		`json:"id"`
	Name        string	`json:"name" validate:"required"`
	Description string	`json:"description"`
	Price       int		`json:"price" validate:"required,gt=0"`
}

type ProductsList []*Product

func (p *ProductsList) ToJson(w io.Writer) error{
	e:=json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error{
	e:=json.NewDecoder(r)
	return e.Decode(p)
}

func(p *Product) Validate() error{
	validate:=validator.New()
	return validate.Struct(p)
}

func GetProductsList() ProductsList {
	return productsList
}
func AddProduct(p *Product) {
	id := GetId()
	p.ID = id
	productsList = append(productsList, p)
}
func UpdateProduct(id int, p *Product) error {
	_,pos,err:=findProduct(id)
	if err!=nil{
		return err
	}

	productsList[pos]=p
	return nil
}

func GetId() int {
	n := productsList[len(productsList)-1]
	return n.ID + 1
}

var ErrProductNotFound = fmt.Errorf("product not found")
func findProduct(id int) (*Product, int, error) {
	for i, p := range productsList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil,-1,ErrProductNotFound
}

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
