package request

import "minpro_arya/features/product"

type Product struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	NetWeight   int    `json:"net-weight"`
	Price       int    `json:"price"`
}

func (req *Product) ToDomain() *product.Domain {
	return &product.Domain{
		Name:        req.Name,
		Description: req.Description,
		NetWeight:   req.NetWeight,
		Price:       req.Price,
	}
}
