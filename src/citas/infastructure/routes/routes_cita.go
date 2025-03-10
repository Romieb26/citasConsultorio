package routes

import (
	usecases "apiHospital/src/citas/application/use_cases"
	"apiHospital/src/citas/domain/repositories"
	"apiHospital/src/citas/infastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCitaRoutes(router *gin.Engine, repo repositories.ICita) {
	createCitaUseCase := usecases.NewCreateCita(repo)
	createCitaController := controllers.NewCreateCitaController(createCitaUseCase)
	deleteCitaUseCase := usecases.NewDeleteCita(repo)
	deleteCitaController := controllers.NewDeleteCitaController(deleteCitaUseCase)

	citaGroup := router.Group("/citas")
	{
		citaGroup.POST("", createCitaController.Run)
		citaGroup.GET("", controllers.GetAllCitasController(repo))
		citaGroup.DELETE("/:id", deleteCitaController.DeleteCita)
	}
}
