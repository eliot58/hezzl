package main

import (
	"github.com/eliot58/hezzl/internal/database"
	"github.com/eliot58/hezzl/internal/handlers"
	"github.com/gin-gonic/gin"
)
 
 func main() {
	route := gin.Default()
	
	database.ConnectDatabase()

	route.POST("/good/create", handlers.CreateGood)

	route.PATCH("/good/update", handlers.UpdateGood)

	route.DELETE("/good/remove", handlers.DeleteGood)

	route.GET("/goods/list", handlers.GetGoods)

	route.PATCH("/good/reprioritize", handlers.Reprioritize)

	err := route.Run(":8080")
	if err != nil {
	   panic(err)
	}
 
 }