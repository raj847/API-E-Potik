package data

import (
	"minpro_arya/features/customer"

	"minpro_arya/features/customer/bussiness"

	"gorm.io/gorm"
)

type MysqlCustomerRepository struct {
	Conn *gorm.DB
}

func NewMysqlCustomerRepository(conn *gorm.DB) customer.Repository {
	return &MysqlCustomerRepository{
		Conn: conn,
	}
}

func (rep *MysqlCustomerRepository) Register(domain *customer.Domain) (customer.Domain, error) {

	cus := fromDomain(*domain)

	result := rep.Conn.Create(&cus)
	if result.Error != nil {
		return customer.Domain{}, result.Error
	}

	return toDomain(cus), nil
}

func (rep *MysqlCustomerRepository) Login(email, password string) (customer.Domain, error) {
	var cus Customer
	err := rep.Conn.First(&cus, "email = ?", email).Error

	if err != nil {
		return customer.Domain{}, bussiness.ErrEmailorPass
	}

	return toDomain(cus), nil
}
