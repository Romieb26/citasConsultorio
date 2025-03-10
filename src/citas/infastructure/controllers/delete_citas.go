package controllers

import (
	"net/http"
	"strconv"

	usecases "apiHospital/src/citas/application/use_cases"

	"github.com/gin-gonic/gin"
)

type DeleteCitaController struct {
	useCase *usecases.DeleteCita
}

func NewDeleteCitaController(useCase *usecases.DeleteCita) *DeleteCitaController {
	return &DeleteCitaController{useCase: useCase}
}

func (dc *DeleteCitaController) DeleteCita(c *gin.Context) {
	// Obtener el ID de la cita desde la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cita inválido"})
		return
	}

	// Ejecutar el caso de uso
	err = dc.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cita eliminada exitosamente"})
}
