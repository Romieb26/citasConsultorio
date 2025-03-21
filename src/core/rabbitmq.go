// core/rabbitmq.go
package core

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection

func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://romina:romina264@3.230.241.180:5672/")
	if err != nil {
		log.Fatalf("No se pudo conectar a RabbitMQ: %v", err)
	}
	RabbitMQConn = conn
	log.Println("Conectado a RabbitMQ")
}

func PublishMessage(queueName string, message []byte) error {
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declarar la cola (se crea automáticamente si no existe)
	_, err = ch.QueueDeclare(
		queueName, // nombre de la cola
		true,      // durable: la cola sobrevive a reinicios del servidor
		false,     // delete when unused: no se elimina cuando no hay consumidores
		false,     // exclusive: la cola no es exclusiva para esta conexión
		false,     // no-wait: no espera confirmación del servidor
		nil,       // arguments: sin argumentos adicionales
	)
	if err != nil {
		return err
	}

	// Publicar el mensaje en la cola
	err = ch.Publish(
		"",        // exchange: se usa el exchange por defecto
		queueName, // routing key: nombre de la cola
		false,     // mandatory: no es obligatorio que el mensaje sea enrutado
		false,     // immediate: no es necesario que el mensaje sea consumido inmediatamente
		amqp.Publishing{
			ContentType:  "application/json", // tipo de contenido del mensaje
			Body:         message,            // cuerpo del mensaje
			DeliveryMode: amqp.Persistent,    // hace que el mensaje sea duradero
		},
	)
	return err
}
