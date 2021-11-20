package admin

import "time"

type Domain struct {
	Id        uint
	Username  string
	Password  string
	RoleID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(userData *Domain) (Domain, error)
	Login(username, password string) error
	// GetByID(id uint) (Domain, error)
}

type Repository interface {
	Register(userData *Domain) (Domain, error)
	GetByUsername(username string) (Domain, error)
	// GetByID(id uint) (Domain, error)
}
