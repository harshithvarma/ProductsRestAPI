package data

import (
	"encoding/json"
	"io"
)

//ToJson serializes the given interface into a string based JSON format
func (p *ProductsList) ToJson(w io.Writer) error{
	e:=json.NewEncoder(w)
	return e.Encode(p)
}

//FromJson decodes the body in request to a product struct
func (p *Product) FromJson(r io.Reader) error{
	e:=json.NewDecoder(r)
	return e.Decode(p)
}