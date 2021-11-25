package data

import (
	"minpro_arya/features/customer"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID           int       `json:"id" form:"id" gorm:"primary_key"`
	Username     string    `json:"username" form:"username" gorm:"unique"`
	Email        string    `json:"email" form:"email" gorm:"unique"`
	Password     string    `json:"password" form:"password"`
	Name         string    `json:"name" form:"name"`
	Dob          string    `json:"dob" form:"dob"`
	Phone_Number string    `json:"phone_number" form:"phone_number"`
	Photo        string    `json:"photo" form:"photo"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}

func toDomain(cus Customer) customer.Domain {
	return customer.Domain{
		ID:           cus.ID,
		Username:     cus.Username,
		Email:        cus.Email,
		Password:     cus.Password,
		Name:         cus.Name,
		Dob:          cus.Dob,
		Phone_Number: cus.Phone_Number,
		Photo:        cus.Photo,
		CreatedAt:    cus.CreatedAt,
		UpdatedAt:    cus.UpdatedAt,
	}
}

func fromDomain(domain customer.Domain) Customer {
	return Customer{
		ID:           domain.ID,
		Username:     domain.Username,
		Email:        domain.Email,
		Password:     domain.Password,
		Name:         domain.Name,
		Dob:          domain.Dob,
		Phone_Number: domain.Phone_Number,
		Photo:        domain.Photo,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
