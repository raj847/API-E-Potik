package transactions

import "time"

type Domain struct {
	ID               int
	CustomerID       int
	ProductID        int
	Transaction_code string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	Trans(customerID int, domain *Domain) (Domain, error)
	GetTransByID(id int) (Domain, error)
	GetAllTrans() ([]Domain, error)
	GetAllCustomerTrans(customerID int) ([]Domain, error)
}

type Repository interface {
	Trans(customerID int, domain *Domain) (Domain, error)
	GetTransByID(id int) (Domain, error)
	GetAllTrans() ([]Domain, error)
	GetAllCustomerTrans(customerID int) ([]Domain, error)
}
