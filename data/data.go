package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Product struct {
	ID          int
	Name        string
	Description string
	price       int
}

type ProductsList []*Product

func (p *Product) ToJson(w io.Writer) error{
	e:=json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error{
	e:=json.NewDecoder(r)
	return e.Decode(p)
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
	p.ID=id
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
		price:       100,
	},
	&Product{
		ID:          2,
		Name:        "coffee",
		Description: "Made in ooty",
		price:       200,
	},
	&Product{
		ID:          3,
		Name:        "osmania biscuits",
		Description: "Made in Hyderabad",
		price:       300,
	},
}
