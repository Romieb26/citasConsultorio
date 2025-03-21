package main

import (
	"log"

	"apiHospital/src/citas/infrastructure"
	"apiHospital/src/core"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	core.InitDB()

	// Inicializar la conexión a RabbitMQ
	core.InitRabbitMQ()

	// Crear un router con Gin
	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Ajusta el puerto según sea necesario
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Inicializar dependencias
	citaRouter := infrastructure.NewCitaRouter(router)
	citaRouter.Run() // Agregar rutas

	// Iniciar el servidor
	log.Println("Servidor corriendo en http://localhost:8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
