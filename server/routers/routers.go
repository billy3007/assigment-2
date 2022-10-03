package routers

import (
	docs "assigment-2/docs"
	"assigment-2/server/controller"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	router *gin.Engine
	order  *controller.OrderController
}

func NewRouter(r *gin.Engine, order *controller.OrderController) *router {
	return &router{
		router: r,
		order:  order,
	}
}

func (r *router) SetupRouter(port string) {
	docs.SwaggerInfo.BasePath = ""
	r.router.GET("/orders", r.order.GetAll)
	r.router.POST("/orders", r.order.CreateOrder)
	r.router.PUT("/orders/:id", r.order.UpdateOrder)
	r.router.DELETE("/orders/:id", r.order.DeleteOrder)
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.router.Run(port)
}
