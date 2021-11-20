package admin

import (
	"minpro_arya/bussiness/admin"

	"gorm.io/gorm"
)

type mysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) admin.Repository {
	return &mysqlAdminRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlAdminRepository) Register(userData *admin.Domain) (admin.Domain, error) {
	recUser := fromDomain(*userData)
	err := mysqlRepo.Conn.Create(&recUser).Error
	if err != nil {
		return admin.Domain{}, err
	}
	return recUser.toDomain(), nil
}

func (mysqlRepo *mysqlAdminRepository) GetByUsername(username string) (admin.Domain, error) {
	rec := Admins{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).Error
	if err != nil {
		return admin.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlAdminRepository) GetByID(id uint) (admin.Domain, error) {
	rec := Admins{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return admin.Domain{}, err
	}
	return rec.toDomain(), nil
}
