package main

import (
	"apiHospital/src/citas/infastructure/adapters"
	"apiHospital/src/citas/infastructure/routes"
	"apiHospital/src/core"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := core.ConnectToDataBase()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	citaRepo := adapters.NewMySQLRepository(db)

	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Cambia esto según tu frontend
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupCitaRoutes(router, citaRepo)

	log.Println("Iniciando el Servidor en el puerto 8080...")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
