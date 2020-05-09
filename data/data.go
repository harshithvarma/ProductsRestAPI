package data

type Product struct {
	ID          int
	Name        string
	Description string
	price       int
}

var productsList=[]*Product{
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