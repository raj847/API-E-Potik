package company

import "time"

type Domain struct {
	ID           int
	Username     string
	Email        string
	Password     string
	Name         string
	Phone_Number string
	Photo        string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(email, password string) (Domain, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(email, password string) (Domain, error)
}
