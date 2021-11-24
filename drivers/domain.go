package drivers

import (
	"minpro_arya/features/admins"
	adminDB "minpro_arya/features/admins/data"
	"minpro_arya/features/company"
	companyDB "minpro_arya/features/company/data"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) admins.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
func NewCompanyRepository(conn *gorm.DB) company.Repository {
	return companyDB.NewMysqlCompanyRepository(conn)
}
