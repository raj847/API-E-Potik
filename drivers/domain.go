package drivers

import (
	"minpro_arya/features/admins"
	adminDB "minpro_arya/features/admins/data"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) admins.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
