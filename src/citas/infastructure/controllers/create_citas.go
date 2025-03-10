package controllers

import (
	usecases "apiHospital/src/citas/application/use_cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCitaController struct {
	useCase *usecases.CreateCita
}

func NewCreateCitaController(useCase *usecases.CreateCita) *CreateCitaController {
	return &CreateCitaController{useCase: useCase}
}

func (cc *CreateCitaController) Run(c *gin.Context) {
	var input struct {
		NombrePaciente   string `json:"nombrePaciente"`
		ApellidoPaciente string `json:"apellidoPaciente"`
		NumeroContacto   string `json:"numeroContacto"`
		AreaCita         string `json:"areaCita"`
		Fecha            string `json:"fecha"`
		Hora             string `json:"hora"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cc.useCase.Execute(input.NombrePaciente, input.ApellidoPaciente, input.NumeroContacto, input.AreaCita, input.Fecha, input.Hora)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Cita creada exitosamente"})
}
