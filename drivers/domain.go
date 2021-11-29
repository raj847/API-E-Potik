package drivers

import (
	"minpro_arya/features/admins"
	adminDB "minpro_arya/features/admins/data"
	"minpro_arya/features/company"
	companyDB "minpro_arya/features/company/data"
	"minpro_arya/features/customer"
	customerDB "minpro_arya/features/customer/data"
	"minpro_arya/features/product"
	productDB "minpro_arya/features/product/data"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) admins.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
func NewCompanyRepository(conn *gorm.DB) company.Repository {
	return companyDB.NewMysqlCompanyRepository(conn)
}
func NewCustomerRepository(conn *gorm.DB) customer.Repository {
	return customerDB.NewMysqlCustomerRepository(conn)
}
func NewProductRepository(conn *gorm.DB) product.Repository {
	return productDB.NewMysqlProductRepository(conn)
}
