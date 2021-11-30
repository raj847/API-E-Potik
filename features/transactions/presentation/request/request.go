package request

import (
	"minpro_arya/features/transactions"
)

type Transactions struct {
	ProductID int `json:"product_id"`
}

func (req *Transactions) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		ProductID: req.ProductID,
	}
}
