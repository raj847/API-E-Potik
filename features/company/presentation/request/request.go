package request

import (
	"minpro_arya/features/company"
)

type Company struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone_Number string `json:"phone_number"`
	Photo        string `json:"photo"`
}

func (req *Company) ToDomain() *company.Domain {
	return &company.Domain{
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		Name:         req.Name,
		Phone_Number: req.Phone_Number,
		Photo:        req.Photo,
	}
}

type CompanyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
