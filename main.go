package main

import (
	"os"

	"restaurant-management/database"
	middleware "restaurant-management/middleware"
	routes "restaurant-management/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)

	router.Run(":", port)

}
