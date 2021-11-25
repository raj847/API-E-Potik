package request

import "minpro_arya/features/customer"

type Customer struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Dob          string `json:"dob"`
	Phone_Number string `json:"phone_number"`
	Photo        string `json:"photo"`
}

func (req *Customer) ToDomain() *customer.Domain {
	return &customer.Domain{
		Username:     req.Username,
		Password:     req.Password,
		Email:        req.Email,
		Name:         req.Name,
		Dob:          req.Dob,
		Phone_Number: req.Phone_Number,
		Photo:        req.Photo,
	}
}

type CustomerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
