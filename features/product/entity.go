package product

import "time"

type Domain struct {
	ID          int
	CompanyID   int
	Name        string
	Description string
	NetWeight   int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	AllProduct() ([]Domain, error)
	Create(orgID int, domain *Domain) (Domain, error)
	Update(orgID int, prodID int, domain *Domain) (Domain, error)
	Delete(orgID, id int) (string, error)
	MyProductByCompany(orgID int) ([]Domain, error)
	ProductByID(id int) (Domain, error)
	ProductByIdCompany(orgzID int) ([]Domain, error)
}

type Repository interface {
	AllProduct() ([]Domain, error)
	Create(orgID int, domain *Domain) (Domain, error)
	Update(orgID int, prodID int, domain *Domain) (Domain, error)
	Delete(orgID, id int) (string, error)
	MyProductByCompany(orgID int) ([]Domain, error)
	ProductByID(id int) (Domain, error)
	ProductByIdCompany(orgzID int) ([]Domain, error)
}
