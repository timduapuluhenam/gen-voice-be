package main

import (
	"fmt"
	"log"

	"genVoice/helper/rimender"

	_driverFactory "genVoice/drivers"

	_userService "genVoice/business/users"
	_userController "genVoice/controllers/users"
	_userRepo "genVoice/drivers/databases/users"

	_activityRepo "genVoice/drivers/databases/activities"

	_invoiceService "genVoice/business/invoices"
	_invoiceController "genVoice/controllers/invoices"
	_invoiceRepo "genVoice/drivers/databases/invoices"

	_notifService "genVoice/business/notifications"
	_notifController "genVoice/controllers/notifications"

	_dbDriver "genVoice/drivers/postgres"

	_middleware "genVoice/app/middlewares"
	_routes "genVoice/app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	midtrans "github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./app/configs/")
	viper.AddConfigPath(".")
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
		&_userRepo.Users{},
		&_invoiceRepo.Invoices{},
		&_invoiceRepo.InvoiceDetail{},
		&_activityRepo.Activities{},
	)
}

var midclient midtrans.Client
var coreGateway midtrans.CoreGateway
var snapGateway midtrans.SnapGateway

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

	midclient = midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-sYHf9k6xSdZJa780ILj-MYXB"
	midclient.ClientKey = "SB-Mid-client-SwacphxrChYFYsTR"
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway = midtrans.CoreGateway{
		Client: midclient,
	}

	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	e := echo.New()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo, 240, &configJWT)
	userCtrl := _userController.NewUserController(userService)

	activityRepo := _driverFactory.NewActivityRepository(db)

	invoiceRepo := _driverFactory.NewInvoiceRepository(db)
	invoiceService := _invoiceService.NewInvoiceService(invoiceRepo, activityRepo, 240, &configJWT)
	invoiceCtrl := _invoiceController.NewInvoiceController(invoiceService)

	notifRepo := _driverFactory.NewNotifRepository(db)
	notifService := _notifService.NewNotifService(notifRepo, activityRepo, 240, &configJWT)
	notifCtrl := _notifController.NewNotifController(notifService)

	routesInit := _routes.ControllerList{
		JwtConfig:         configJWT.Init(),
		UserController:    *userCtrl,
		InvoiceController: *invoiceCtrl,
		NotifController:   *notifCtrl,
	}

	routesInit.RouteRegister(e)
	godotenv.Load()
	port := viper.GetString(`server.address`)
	fmt.Print(port)
	address := fmt.Sprintf("%s%s", "0.0.0.0", port)
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	go rimender.Rimender(db)
	log.Fatal(e.Start(address))
}
