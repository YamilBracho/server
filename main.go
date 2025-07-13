package main

import (
	"fmt"
	"log"

	"github.com/YamilBracho/server/repository"
	"github.com/YamilBracho/server/services"
	"github.com/gin-gonic/gin"
)

// --------------------------------------------------------------------
// handleRequests
// Set URL handling routines
// --------------------------------------------------------------------
func handleRequests(router *gin.Engine) {
	router.GET("/api/provincias", services.GetProvincias)
	router.POST("/api/distritos", services.GetDistritosByProvincia)
	router.POST("/api/corregimientos", services.GetCorregimientosByDistrito)
	router.POST("/api/barrios", services.GetBarriosByCorregimiento)
}

// --------------------------------------------------------------------
// Main
// Opens database connections, sets request handlers and start server
// --------------------------------------------------------------------
func main() {
	if err := repository.SetConnection(); err != nil {
		log.Println(err)
		return
	}
	defer repository.DB.Close()

	router := gin.Default()
	handleRequests(router)

	fmt.Println("Server started at localhost:8080...")
	router.Run("localhost:8080")

}
