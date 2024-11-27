package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	pb "Tarea1/protofiles"

	"google.golang.org/grpc"
)

const (
	serverLogistica = "dist018.inf.santiago.usm.cl:50052"
	intervaloEnvio  = 0  // Intervalo entre envíos
)

func main() {
	// Conectar con el servidor gRPC
	conn, err := grpc.Dial(serverLogistica, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error conectando al servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewClientServiceClient(conn)

	// Contador para las órdenes
	contador := 0

	// Crear un WaitGroup para esperar a que todas las goroutines terminen
	var wg sync.WaitGroup
	file, err := os.Open("solicitudes.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cont := 0
	// Leer línea por línea
	for scanner.Scan() {
		if cont == 0 {
			cont++
		} else {
			contador++
			linea := scanner.Text()
			parametros := strings.Split(linea, ",")
			tipo := parametros[0]
			puesto := parametros[1]
			unidad := parametros[2]
			item := parametros[3]
			cantidad := 0
			fmt.Sscanf(parametros[4], "%d", &cantidad)
			wg.Add(1)
			go enviarOrden(&wg, client, tipo, fmt.Sprintf("%d", contador), puesto, unidad, item, int32(cantidad))
			time.Sleep(intervaloEnvio * time.Second)
		}
	}
	// Espera a que todas las goroutines terminen
	wg.Wait()

	// Simula la consulta del estado del paquete después de que todas las órdenes han sido procesadas
	consultarEstado(client)
}

func enviarOrden(wg *sync.WaitGroup, client pb.ClientServiceClient, tipo string, paqueteID string, destino string, escolta string, nombre string, valor int32) {
	defer wg.Done() // Decrementa el contador del WaitGroup cuando la goroutine termina

	// Crea una nueva solicitud de paquete para la facción especificada
	req := &pb.PackageRequest{
		IdPaquete:   paqueteID,
		Timestamp:   time.Now().Unix(),
		Tipo:        tipo,
		Nombre:      nombre,
		Valor:       valor,
		Escolta:     escolta,
		Destino:     destino,
		Seguimiento: "0", // 0 en el caso que no tenga id de seguimiento
	}

	// Envía la solicitud al servidor logístico
	res, err := client.SendPackage(context.Background(), req)
	if err != nil {
		log.Fatalf("Error al enviar paquete (%s): %v", tipo, err)
	}

	log.Printf("Se envió y fue recibido en logistica/konzu el paquete con ID: %s, Tipo %s, Nombre %s, Valor %d, Escolta %s, Destino %s, \nSe obtuvo el código de seguimiento: %s | Estado: %s\n",
	paqueteID, tipo, nombre, valor, escolta, destino, res.Seguimiento, res.Estado)
}

func consultarEstado(client pb.ClientServiceClient) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Solicita por terminal el código de seguimiento
		fmt.Printf("Ingrese el código de seguimiento del paquete (o presione 'q' para salir): ")
		codigoSeguimiento, _ := reader.ReadString('\n')
		codigoSeguimiento = codigoSeguimiento[:len(codigoSeguimiento)-1] // Remueve el carácter de nueva línea

		if codigoSeguimiento == "q" {
			fmt.Println("Saliendo del ciclo de consultas.")
			break
		}

		// Crea una solicitud para consultar el estado del paquete
		req := &pb.StatusRequest{
			Seguimiento: codigoSeguimiento,
		}

		// Consulta el estado del paquete
		res, err := client.GetPackageStatus(context.Background(), req)
		if err != nil {
			log.Printf("Error al consultar estado del paquete con seguimiento (%s): %v", codigoSeguimiento, err)
			continue
		}

		// Imprime la respuesta
		tiempoFormateado := time.Unix(res.Timestamp, 0).Format(time.RFC3339)
		log.Println("------------------------------------------------")
		log.Printf("Estado del paquete con codigo de seguimiento %s:\n", codigoSeguimiento)
		log.Printf("  ID del paquete: %s\n", res.IdPaquete)
		log.Printf("  TimeStamp: %s\n", tiempoFormateado)
		log.Printf("  Tipo: %s\n", res.Tipo)
		log.Printf("  Destino: %s\n", res.Destino)
		log.Printf("  Valor: %d\n", res.Valor)
		log.Printf("  Escolta: %s\n", res.Escolta)
		log.Printf("  Estado: %s\n", res.Estado)
		log.Printf("  Intentos: %d\n", res.Intentos)
		log.Println("------------------------------------------------")
	}

	client.TerminarConexion(context.Background(), &pb.TerminarConexionMensaje{Mensaje: "Terminar conexión"})
}
