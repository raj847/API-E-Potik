package data

import (
	"minpro_arya/features/product"

	"minpro_arya/features/product/bussiness"

	"gorm.io/gorm"
)

type MysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductRepository(conn *gorm.DB) product.Repository {
	return &MysqlProductRepository{
		Conn: conn,
	}
}

func (rep *MysqlProductRepository) Create(orgID int, domain *product.Domain) (product.Domain, error) {

	pr := fromDomain(*domain)

	pr.CompanyID = orgID

	result := rep.Conn.Create(&pr)

	if result.Error != nil {
		return product.Domain{}, result.Error
	}

	return toDomain(pr), nil

}

func (rep *MysqlProductRepository) AllProduct() ([]product.Domain, error) {

	var prod []Product

	result := rep.Conn.Find(&prod)

	if result.Error != nil {
		return []product.Domain{}, result.Error
	}

	return toDomainList(prod), nil

}

func (rep *MysqlProductRepository) Update(orgID int, prID int, domain *product.Domain) (product.Domain, error) {
	productUpdate := fromDomain(*domain)

	productUpdate.ID = prID
	result := rep.Conn.Where("id = ?", prID).Updates(&productUpdate)

	if result.Error != nil {
		return product.Domain{}, bussiness.ErrNotFound
	}

	return toDomainUpdate(productUpdate), nil
}

func (rep *MysqlProductRepository) Delete(orgID int, id int) (string, error) {
	rec := Product{}

	find := rep.Conn.Where("id = ?", id).First(&rec)

	if find.Error != nil {
		return "", bussiness.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", id).Error

	if err != nil {
		return "", bussiness.ErrNotFound
	}

	return "Product has been delete", nil

}

func (rep *MysqlProductRepository) MyProductByCompany(orgID int) ([]product.Domain, error) {
	var prod []Product

	result := rep.Conn.Where("company_id = ?", orgID).Find(&prod)

	if result.Error != nil {
		return []product.Domain{}, result.Error
	}

	return toDomainList(prod), nil
}

func (rep *MysqlProductRepository) ProductByID(id int) (product.Domain, error) {

	var prod Product

	result := rep.Conn.Where("id = ?", id).First(&prod)

	if result.Error != nil {
		return product.Domain{}, result.Error
	}

	return toDomain(prod), nil
}
func (rep *MysqlProductRepository) ProductByIdCompany(orgzID int) ([]product.Domain, error) {

	var prod []Product

	result := rep.Conn.Where("Company_id = ?", orgzID).Find(&prod)

	if result.Error != nil {
		return []product.Domain{}, result.Error
	}

	return toDomainList(prod), nil
}
