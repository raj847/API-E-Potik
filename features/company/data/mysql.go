package data

import (
	"minpro_arya/features/company"
	"minpro_arya/features/company/bussiness"

	"gorm.io/gorm"
)

type MysqlCompanyRepository struct {
	Conn *gorm.DB
}

func NewMysqlCompanyRepository(conn *gorm.DB) company.Repository {
	return &MysqlCompanyRepository{
		Conn: conn,
	}
}

func (rep *MysqlCompanyRepository) Register(domain *company.Domain) (company.Domain, error) {

	org := fromDomain(*domain)

	result := rep.Conn.Create(&org)

	if result.Error != nil {
		return company.Domain{}, result.Error
	}

	return toDomain(org), nil
}

func (rep *MysqlCompanyRepository) Login(email, password string) (company.Domain, error) {
	var org Company
	err := rep.Conn.First(&org, "email = ?", email).Error

	if err != nil {
		return company.Domain{}, bussiness.ErrEmailorPass
	}

	return toDomain(org), nil
}
