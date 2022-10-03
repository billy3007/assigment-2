package main

import (
	"assigment-2/config"
	"assigment-2/server/controller"
	"assigment-2/server/repositories/gorm"
	"assigment-2/server/routers"
	"assigment-2/server/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectMysqlGORM()

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	itemRepo := gorm.NewItemRepository(db)
	orderRepo := gorm.NewOrderRepository(db)

	orderService := services.NewOrderService(orderRepo, itemRepo)

	orderController := controller.NewOrderController(orderService)

	app := routers.NewRouter(router, orderController)
	app.SetupRouter(":3001")
}
