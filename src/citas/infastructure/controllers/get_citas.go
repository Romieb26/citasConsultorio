package controllers

import (
	"apiHospital/src/citas/domain/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCitasController(repo repositories.ICita) gin.HandlerFunc {
	return func(c *gin.Context) {
		citas, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, citas)
	}
}
