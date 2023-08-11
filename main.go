package main

import (
	"Go-project/constants"
	"Go-project/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	routes.AppRoutes(router)
	fmt.Println("server running on port ",constants.Port)
	log.Fatal(router.Run(constants.Port))
}