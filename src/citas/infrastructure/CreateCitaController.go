// CreateCitaController.go
package infrastructure

import (
	"encoding/json"
	"net/http"

	"apiHospital/src/citas/application"
	"apiHospital/src/citas/domain/entities"
	"apiHospital/src/core"

	"github.com/gin-gonic/gin"
)

type CreateCitaController struct {
	createUseCase *application.CreateCitaUseCase
}

func NewCreateCitaController(createUseCase *application.CreateCitaUseCase) *CreateCitaController {
	return &CreateCitaController{
		createUseCase: createUseCase,
	}
}

func (ctrl *CreateCitaController) Run(c *gin.Context) {
	var citaRequest struct {
		NombrePaciente   string `json:"nombrePaciente"`
		ApellidoPaciente string `json:"apellidoPaciente"`
		NumeroContacto   string `json:"numeroContacto"`
		AreaCita         string `json:"areaCita"`
		Fecha            string `json:"fecha"`
		Hora             string `json:"hora"`
		Estado           string `json:"estado"`
	}

	if err := c.ShouldBindJSON(&citaRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	cita := entities.NewCita(
		citaRequest.NombrePaciente,
		citaRequest.ApellidoPaciente,
		citaRequest.NumeroContacto,
		citaRequest.AreaCita,
		citaRequest.Fecha,
		citaRequest.Hora,
		citaRequest.Estado,
	)

	createdCita, err := ctrl.createUseCase.Run(cita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo crear la cita",
			"error":   err.Error(),
		})
		return
	}

	// Publicar mensaje en RabbitMQ
	message, err := json.Marshal(createdCita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al serializar la cita",
			"error":   err.Error(),
		})
		return
	}

	err = core.PublishMessage("citas_creadas", message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al publicar mensaje en RabbitMQ",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdCita)
}
