package services

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/YamilBracho/server/models"
	"github.com/YamilBracho/server/repository"
)

// --------------------------------------------------------------------
// GetProvincias
// Get list of Provincias
// --------------------------------------------------------------------
func GetProvincias(c *gin.Context) {
	if rows, err := repository.GetProvincias(); err != nil {
		c.JSON(http.StatusOK, gin.H{"Result": "ERROR", "Message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Result": "OK", "Records": rows})
	}
}

// --------------------------------------------------------------------
// GetDistritosByProvincia
// Gets list of Distritos by Provincia
// --------------------------------------------------------------------
func GetDistritosByProvincia(c *gin.Context) {

	var request models.DistritoRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"Result": "ERROR", "Message": err.Error()})
		return
	}

	if rows, err := repository.GetDistritosByProvincia(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"Result": "ERROR", "Message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Result": "OK", "Records": rows})
	}

}

// --------------------------------------------------------------------
// GetCorregimientosByDistrito
// Gets list of Corregimientos By Distrito
// --------------------------------------------------------------------
func GetCorregimientosByDistrito(c *gin.Context) {
	var request models.CorregimientoRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"Result": "ERROR", "Message": err.Error()})
		return
	}

	if rows, err := repository.GetCorregimientosByDistrito(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"Result": "ERROR", "Message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Result": "OK", "Records": rows})
	}
}

// --------------------------------------------------------------------
// GetBarriosByCorregimiento
// Gets list of Barrios by Corregimiento
// --------------------------------------------------------------------
func GetBarriosByCorregimiento(c *gin.Context) {
	var request models.BarrioRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{"Result": "ERROR", "Message": err.Error()})
		return
	}

	if rows, err := repository.GetBarriosByCorregimiento(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"Result": "ERROR", "Message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"Result": "OK", "Records": rows})
	}

}
