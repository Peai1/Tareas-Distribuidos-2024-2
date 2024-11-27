package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/streadway/amqp"
)

// Estructura del mensaje recibido desde RabbitMQ
type DeliveryResult struct {
	IdPaquete string `json:"id_paquete"`
	Exito     bool   `json:"exito"`
	Valor     int32  `json:"valor"`
	Intentos  int32  `json:"intentos"`
	Creditos  int32  `json:"creditos"`
	Tipo      string `json:"tipo"`
}

func main() {
	// Conectar con RabbitMQ
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%s/", os.Getenv("RABBITMQ_SERVER"), os.Getenv("RABBITMQ_PORT")))
	if err != nil {
		log.Fatalf("error al conectar con RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("error al abrir canal en RabbitMQ: %v", err)
	}
	defer ch.Close()

	// Declarar la cola desde la que se recibirán los mensajes
	q, err := ch.QueueDeclare(
		"logistica_finanzas", // Nombre de la cola
		false,                // Durable
		false,                // Delete when unused
		false,                // Exclusive
		false,                // No-wait
		nil,                  // Arguments
	)
	if err != nil {
		log.Fatalf("error al declarar la cola: %v", err)
	}
	file, err := os.Open("solicitudes.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalDeliveries := 0
	// Leer línea por línea
	if scanner.Scan() {
		linea := scanner.Text()
		parametros := strings.Split(linea, ",")
		fmt.Sscanf(parametros[0], "%d", &totalDeliveries)

	}

	// Verifica si hubo algún error al leer el archivo
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
	// Número total de entregas esperadas (ajusta esto según tus necesidades)
	var deliveriesProcessed int
	var mu sync.Mutex
	var balanceFinal int32

	// Suscribirse a los mensajes de la cola
	msgs, err := ch.Consume(
		q.Name, // Nombre de la cola
		"",     // Consumer
		true,   // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("error al suscribirse a la cola: %v", err)
	}

	// Canal para manejar la señal de finalización
	done := make(chan bool)

	go func() {
		for d := range msgs {
			// Procesar cada mensaje recibido
			var result DeliveryResult
			err := json.Unmarshal(d.Body, &result)
			if err != nil {
				log.Printf("error al deserializar mensaje: %v", err)
				continue
			}

			// Procesar el resultado de la entrega
			procesarResultado(result, &balanceFinal)

			// Incrementar el contador de entregas procesadas
			mu.Lock()
			deliveriesProcessed++
			if deliveriesProcessed >= totalDeliveries {
				done <- true // Enviar señal de finalización cuando se procesen todas las entregas
			}
			mu.Unlock()
		}
	}()

	log.Println("Esperando mensajes de Konzu...")
	<-done

	log.Printf("------------------------------------------------")
	log.Printf("Se han procesado todas las entregas. Balance final: %d creditos. Cerrando finanzas...", balanceFinal)
	log.Printf("------------------------------------------------")
}

func procesarResultado(result DeliveryResult, balanceFinal *int32) {
	const (
		colorReset = "\033[0m"
		colorRed   = "\033[31m"
		colorGreen = "\033[32m"
	)

	if result.Exito {
		log.Printf("Entrega %s%s%s para el paquete %s (tipo %s): balance = %d creditos | intentos = %d", colorGreen, "exitosa", colorReset, result.IdPaquete, result.Tipo, result.Creditos, result.Intentos)
	} else {
		log.Printf("Entrega %s%s%s para el paquete %s (tipo %s): balance = %d creditos | intentos = %d", colorRed, "fallida", colorReset, result.IdPaquete, result.Tipo, result.Creditos, result.Intentos)
	}

	// Actualizar el balance final
	*balanceFinal += result.Creditos

}
