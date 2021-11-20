package main

import (
	//!handler

	_handlerAdmin "minpro_arya/app/controllers/admin"

	//!routes
	_routes "minpro_arya/app/routes"

	//!service
	_ServAdmin "minpro_arya/bussiness/admin"

	//!Repository
	_repoAdmin "minpro_arya/drivers/mysql/admin"

	//mysql
	"log"
	_dbDriver "minpro_arya/drivers/mysql"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_repoAdmin.Admins{},
		&_repoAdmin.Roles{},
	)
	roles := []_repoAdmin.Roles{{ID: 1, Name: "Owner"}, {ID: 2, Name: "Admin"}}
	db.Create(&roles)
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
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	//* factory of domain
	// ?admin
	adminRepo := _repoAdmin.NewMySQLRepository(db)
	adminServe := _ServAdmin.NewadminService(adminRepo)
	adminHandler := _handlerAdmin.NewUserController(adminServe)

	//* initial of routes
	routesInit := _routes.HandlerRoute{
		AdminController: *adminHandler,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(":8080"))
}
