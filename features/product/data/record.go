package data

import (
	"minpro_arya/features/product"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	CompanyID   int
	Name        string
	Description string
	NetWeight   int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(pr Product) product.Domain {
	return product.Domain{
		ID:          pr.ID,
		CompanyID:   pr.CompanyID,
		Name:        pr.Name,
		Description: pr.Description,
		NetWeight:   pr.NetWeight,
		Price:       pr.Price,
		CreatedAt:   pr.CreatedAt,
		UpdatedAt:   pr.UpdatedAt,
	}
}

func fromDomain(domain product.Domain) Product {
	return Product{
		ID:          domain.ID,
		CompanyID:   domain.CompanyID,
		Name:        domain.Name,
		Description: domain.Description,
		NetWeight:   domain.NetWeight,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func toDomainUpdate(pr Product) product.Domain {
	return product.Domain{
		ID:          pr.ID,
		CompanyID:   pr.CompanyID,
		Name:        pr.Name,
		Description: pr.Description,
		NetWeight:   pr.NetWeight,
		Price:       pr.Price,
		CreatedAt:   pr.CreatedAt,
		UpdatedAt:   pr.UpdatedAt,
	}
}
func toDomainList(data []Product) []product.Domain {
	result := []product.Domain{}

	for _, pr := range data {
		result = append(result, toDomain(pr))
	}
	return result
}
