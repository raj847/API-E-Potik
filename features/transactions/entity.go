package transactions

import "time"

type Domain struct {
	ID               int
	UserID           int
	ProductID        int
	Transaction_code string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	Trans(userID int, domain *Domain) (Domain, error)
	GetTransByID(id int) (Domain, error)
	GetAllTrans() ([]Domain, error)
	GetAllUserTrans(userID int) ([]Domain, error)
}

type Repository interface {
	Trans(userID int, domain *Domain) (Domain, error)
	GetTransByID(id int) (Domain, error)
	GetAllTrans() ([]Domain, error)
	GetAllUserTrans(userID int) ([]Domain, error)
}
