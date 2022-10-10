package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"CrudAPI/controllers"
	"CrudAPI/models"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())

	//InitPostgres()
	models.InitGormPostgres()
	defer models.MPosGORM.Close()

	// Set the router as the default one shipped with Gin
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")

	api.POST("/order/add", controllers.OrderAdd)
	api.POST("/order/delete", controllers.OrderDelete)
	api.POST("/order/edit", controllers.OrderEdit)
	api.POST("/order/show", controllers.OrderShowByDate)
	api.GET("/order/show", controllers.OrderShowByPhone)
	api.GET("/order/id/:orderid", controllers.OrderShowByID)

	// Start and run the server
	router.Run(":4000")
}
