package transactions

import (
	"minpro_arya/features/transactions"
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	ID               int `gorm:"primary_key"`
	CustomerID       int
	ProductID        int
	Transaction_Code string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func toDomain(tr Transactions) transactions.Domain {
	return transactions.Domain{
		ID:               tr.ID,
		CustomerID:       tr.CustomerID,
		ProductID:        tr.ProductID,
		Transaction_code: tr.Transaction_Code,
		CreatedAt:        tr.CreatedAt,
		UpdatedAt:        tr.UpdatedAt,
	}
}

func fromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		ID:               domain.ID,
		CustomerID:       domain.CustomerID,
		ProductID:        domain.ProductID,
		Transaction_Code: domain.Transaction_code,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func toDomainList(data []Transactions) []transactions.Domain {
	result := []transactions.Domain{}

	for _, trans := range data {
		result = append(result, toDomain(trans))
	}
	return result
}
