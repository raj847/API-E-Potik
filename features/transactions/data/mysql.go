package transactions

import (
	"minpro_arya/features/transactions"

	"gorm.io/gorm"
)

type MysqlTransRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransRepository(conn *gorm.DB) transactions.Repository {
	return &MysqlTransRepository{
		Conn: conn,
	}
}

func (rep *MysqlTransRepository) Trans(customerID int, domain *transactions.Domain) (transactions.Domain, error) {

	tr := fromDomain(*domain)

	tr.CustomerID = customerID
	result := rep.Conn.Create(&tr)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return toDomain(tr), nil
}

func (rep *MysqlTransRepository) GetTransByID(id int) (transactions.Domain, error) {

	var trans Transactions

	result := rep.Conn.Where("id = ?", id).First(&trans)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return toDomain(trans), nil

}

func (rep *MysqlTransRepository) GetAllTrans() ([]transactions.Domain, error) {

	trans := []Transactions{}

	result := rep.Conn.Find(&trans)

	if result.Error != nil {
		return []transactions.Domain{}, result.Error
	}

	return toDomainList(trans), nil
}

func (rep *MysqlTransRepository) GetAllCustomerTrans(customerID int) ([]transactions.Domain, error) {

	trans := []Transactions{}

	result := rep.Conn.Where("customer_id = ?", customerID).Find(&trans)

	if result.Error != nil {
		return []transactions.Domain{}, result.Error
	}

	return toDomainList(trans), nil
}
