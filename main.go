package main

import (
	"log"

	_routes "minpro_arya/routes"

	_adminService "minpro_arya/features/admins/bussiness"
	_adminRepo "minpro_arya/features/admins/data"
	_adminController "minpro_arya/features/admins/presentation"

	_dbDriver "minpro_arya/config"

	_driverFactory "minpro_arya/drivers"

	_middleware "minpro_arya/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_adminRepo.Admins{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminService := _adminService.NewServiceAdmin(adminRepo, 10, &configJWT)
	adminCtrl := _adminController.NewHandlerAdmin(adminService)

	routesInit := _routes.RouteList{
		JWTMiddleware: configJWT.Init(),
		AdminRouter:   *adminCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
