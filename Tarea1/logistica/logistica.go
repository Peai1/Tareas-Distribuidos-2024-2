package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	pb "Tarea1/protofiles"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

// Estructura para representar un paquete
type Package struct {
	IdPaquete   string
	TimeStamp   int64
	Tipo        string
	Nombre      string
	Valor       int32
	Escolta     string
	Destino     string
	Seguimiento string
	Intentos    int32
	Estado      string
}

// Estructura del mensaje que se envia a finanzas
type DeliveryResult struct {
	IdPaquete string `json:"id_paquete"`
	Exito     bool   `json:"exito"`
	Valor     int32  `json:"valor"`
	Intentos  int32  `json:"intentos"`
	Creditos  int32  `json:"creditos"`
	Tipo      string `json:"tipo"`
}

const (
	serverCaravana  = "dist017.inf.santiago.usm.cl:50051"
	serverLogistica = ":50052"
)

// Servidor de logistica
type server struct {
	pb.UnimplementedClientServiceServer
	pb.UnimplementedCaravanServiceServer
	packageStore    sync.Map // Almacenamiento en memoria para los paquetes
	packages        []*Package
	received        int
	expected        int
	mu              sync.Mutex              // Mutex
	caravanConn     pb.CaravanServiceClient // Conexion al servicio de caravanas
	rabbitConn      *amqp.Connection        // Conexion a RabbitMQ
	rabbitCh        *amqp.Channel           // RabbitMQ
	done            chan bool
	mensajeMostrado bool
}

// Inicializar el servidor con el número esperado de paquetes
func newServer(expectedPackages int, caravanConn pb.CaravanServiceClient, rabbitConn *amqp.Connection, rabbitCh *amqp.Channel) *server {
	return &server{
		packages:    make([]*Package, 0),
		received:    0,
		expected:    expectedPackages,
		caravanConn: caravanConn,
		rabbitConn:  rabbitConn,
		rabbitCh:    rabbitCh,
		done:        make(chan bool), // indicar que se han recibido todos los paquetes
	}
}

// envio de un nuevo paquete desde cliente a logistica (ClientService)
func (s *server) SendPackage(ctx context.Context, req *pb.PackageRequest) (*pb.PackageResponse, error) {
	trackingCode := "TC" + req.IdPaquete // codigo de seguimient

	// Crear un nuevo paquete y almacenarlo en la lista
	p := &Package{
		IdPaquete:   req.IdPaquete,
		TimeStamp:   req.Timestamp,
		Seguimiento: trackingCode,
		Destino:     req.Destino,
		Escolta:     req.Escolta,
		Nombre:      req.Nombre,
		Valor:       req.Valor,
		Intentos:    0,
		Estado:      "En Cetus",
		Tipo:        req.Tipo,
	}

	// timestamp legible
	tiempoFormateado := time.Unix(req.Timestamp, 0).Format(time.RFC3339)

	s.mu.Lock()
	log.Printf("Recibido Paquete %s Tipo: %s, TimeStamp: %s, Destino: %s, Warframe: %s, Recurso: %s, Valor: %d, Seguimiento: %s.\nSe ha creado su código de seguimiento: %s", 
    req.IdPaquete, req.Tipo, tiempoFormateado, req.Destino, req.Escolta, req.Nombre, req.Valor, req.Seguimiento, trackingCode)

	s.packages = append(s.packages, p)
	s.mu.Unlock()

	// Almacenar el paquete en el sistema
	s.packageStore.Store(trackingCode, p)

	// Verificar si todos los paquetes esperados han llegado
	go s.checkAndDisplayPackages(req)

	// Responder al cliente con el codigo de seguimiento y el estado inicial
	return &pb.PackageResponse{Seguimiento: trackingCode, Estado: "En Cetus"}, nil
}

func (s *server) TerminarConexion(ctx context.Context, req *pb.TerminarConexionMensaje) (*pb.TerminarConexionMensaje, error) {
	log.Printf("Cerrando conexión con el servidor logístico...")
	s.done <- true
	return &pb.TerminarConexionMensaje{Mensaje: "Conexión cerrada"}, nil
}

// Función que verifica si todos los paquetes han llegado y muestra la lista completa
func (s *server) checkAndDisplayPackages(req *pb.PackageRequest) {

	instruction := &pb.DeliveryInstructionRequest{
		Packages: []*pb.PackageRequest{},
	}

	// manda el parametro req en funcion SendDeliveryInstructions
	instruction.Packages = append(instruction.Packages, &pb.PackageRequest{
		IdPaquete:   req.IdPaquete,
		Timestamp:   req.Timestamp,
		Tipo:        req.Tipo,
		Nombre:      req.Nombre,
		Valor:       req.Valor,
		Escolta:     req.Escolta,
		Destino:     req.Destino,
		Seguimiento: req.Seguimiento,
	})

	log.Printf("Enviando paquete %s a caravanas", req.IdPaquete)

	// Llamar al metodo gRPC de las caravanas para enviar las instrucciones
	confirmation, err := s.caravanConn.SendDeliveryInstructions(context.Background(), instruction)
	if err != nil {
		log.Printf("error al enviar instrucciones a las caravanas: %v", err)
	} else {
		// Procesar los resultados de las entregas
		for _, result := range confirmation.Results {
			for _, p := range s.packages {
				if p.IdPaquete == result.IdPaquete {
					p.Intentos = result.Intentos
					estado := ""
					if result.Exito {
						estado = "Entregado"
					} else {
						estado = "No Entregado"
					}
					log.Printf("Resultado para el paquete %s %s (Tipo %s): Creditos restantes = %d | Intentos = %d", result.IdPaquete, estado, p.Tipo, result.Valor, result.Intentos)
					s.received++
				}
			}
		}
	}
	if s.received >= s.expected {
		s.mu.Lock()
		if !s.mensajeMostrado {
			log.Println("Todos los paquetes procesados, esperando solicitudes de seguimiento de cliente para cerrar conexion...")
			s.mensajeMostrado = true
		}
		s.mu.Unlock()
	}
}

// consultar el estado de un paquete (ClientService)
func (s *server) GetPackageStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	// Buscar el paquete en el almacenamiento en variable packages
	if p, ok := s.packageStore.Load(req.Seguimiento); ok {
		packageInfo := p.(*Package)
		log.Printf("Consultado estado del paquete %s con codigo de seguimiento %s", packageInfo.IdPaquete, req.Seguimiento)
		return &pb.StatusResponse{
			IdPaquete: packageInfo.IdPaquete,
			Timestamp: packageInfo.TimeStamp,
			Tipo:      packageInfo.Tipo,
			Destino:   packageInfo.Destino,
			Valor:     packageInfo.Valor,
			Escolta:   packageInfo.Escolta,
			Estado:    packageInfo.Estado,
			Intentos:  packageInfo.Intentos,
		}, nil
	}
	return nil, fmt.Errorf("paquete con codigo %s no encontrado", req.Seguimiento)
}

// actualizar el estado de la entrega desde las caravanas (CaravanService)
func (s *server) UpdateDelivery(ctx context.Context, req *pb.DeliveryUpdateRequest) (*pb.DeliveryUpdateResponse, error) {
	// Actualizar el estado del paquete correspondiente
	for _, p := range s.packages {
		if p.IdPaquete == req.IdPaquete {
			p.Estado = req.Estado
			p.Intentos = req.Intentos
			valorGanancia := int32(0)

			if req.Estado == "Entregado" {
				valorGanancia = p.Valor
			} else {
				if p.Tipo == "Ostronitas" {
					valorGanancia = p.Valor // Se queda con el suministro
				} else if p.Tipo == "Prioritario" {
					valorGanancia = int32(float32(p.Valor) * 0.30) // 30% de ganancia
				} else if p.Tipo == "Normal" {
					valorGanancia = 0 // Se pierde el suministro
				}
			}

			// creditos restantes basados en los intentos realizados
			remainingCredits := valorGanancia - int32(req.Intentos)*100

			// Crear el resultado de la entrega para enviar a RabbitMQ
			deliveryResult := DeliveryResult{
				IdPaquete: p.IdPaquete,
				Exito:     req.Estado == "Entregado", // exito true si el estado es Entregado
				Valor:     p.Valor,
				Intentos:  req.Intentos,
				Creditos:  remainingCredits,
				Tipo:      p.Tipo,
			}

			// Enviar el mensaje a RabbitMQ solo si el paquete ha sido entregado o no entregado
			if req.Estado == "Entregado" || req.Estado == "No Entregado" {
				// Registrar el estado en los logs
				if req.Estado == "Entregado" {
					log.Printf("Paquete %s con Seguimiento %s de Tipo %s fue entregado", req.IdPaquete, req.Seguimiento, req.Tipo)
				} else if req.Estado == "No Entregado" {
					log.Printf("Paquete %s con Seguimiento %s de Tipo %s fue no entregado", req.IdPaquete, req.Seguimiento, req.Tipo)
				}

				// Enviar el resultado a RabbitMQ
				err := s.enviarMensajeRabbitMQ(deliveryResult)
				if err != nil {
					log.Printf("Error al enviar el resultado a RabbitMQ para el paquete %s: %v", req.IdPaquete, err)
				}
			} else {
				log.Printf("Paquete %s con Seguimiento %s de Tipo %s fue Enviado y lleva %d intentos", req.IdPaquete, req.Seguimiento, req.Tipo, req.Intentos)
			}
		}
	}

	return &pb.DeliveryUpdateResponse{Mensaje: "ok"}, nil
}

// Funcion para enviar un mensaje a RabbitMQ
func (s *server) enviarMensajeRabbitMQ(result DeliveryResult) error {
	body, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error al serializar el resultado: %v", err)
	}

	err = s.rabbitCh.Publish(
		"",                   // Exchange
		"logistica_finanzas", // Nombre de la cola
		false,                // Mandatory
		false,                // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("error al publicar el mensaje en RabbitMQ: %v", err)
	}

	return nil
}
func main() {
	// Establecer conexion con RabbitMQ
	rabbitConn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%s/", os.Getenv("RABBITMQ_SERVER"), os.Getenv("RABBITMQ_PORT")))
	if err != nil {
		log.Fatalf("error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()

	rabbitCh, err := rabbitConn.Channel()
	if err != nil {
		log.Fatalf("error al abrir canal en RabbitMQ: %v", err)
	}
	defer rabbitCh.Close()

	// Declarar la cola de RabbitMQ
	_, err = rabbitCh.QueueDeclare(
		"logistica_finanzas", // Nombre de la cola
		false,                // Durable
		false,                // Delete when unused
		false,                // Exclusive
		false,                // No-wait
		nil,                  // Arguments
	)
	if err != nil {
		log.Fatalf("error al declarar la cola en RabbitMQ: %v", err)
	}

	// Establecer conexion con el servidor de caravanas
	conn, err := grpc.Dial(serverCaravana, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error al conectar con el servidor de caravanas: %v", err)
	}
	defer conn.Close()

	caravanConn := pb.NewCaravanServiceClient(conn)

	// numero esperado de paquetes
	file, err := os.Open("solicitudes.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	expectedPackages := 0
	if scanner.Scan() {
		linea := scanner.Text()
		parametros := strings.Split(linea, ",")
		expectedPackages := 0
		fmt.Sscanf(parametros[0], "%d", &expectedPackages)

	}
	s := newServer(expectedPackages, caravanConn, rabbitConn, rabbitCh)

	// servidor de logistica en una goroutine
	go func(){
		lis, err := net.Listen("tcp", serverLogistica)
		if err != nil {
			log.Fatalf("error al iniciar servidor: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterClientServiceServer(grpcServer, s)
		pb.RegisterCaravanServiceServer(grpcServer, s)
		log.Println("Servidor logistico en ejecucion en el puerto 50052...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("error al ejecutar el servidor: %v", err)
		}
	}()

	// Esperar a que se envien todos los paquetes
	<-s.done

	// Finalizar el servidor y cerrar conexiones
	log.Println("Cerrando logistica...")
	time.Sleep(2 * time.Second)
}
