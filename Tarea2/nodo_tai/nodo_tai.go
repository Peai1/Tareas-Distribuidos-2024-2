package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	pb "Tarea2/protofiles"

	"google.golang.org/grpc"
)

type nodoTai struct {
	pb.UnimplementedServicioTaiServer
	vi            int     // Vida Inicial de Greymon/Garurumon
	cd            int     // Cantidad de Datos necesarios para evolucionar
	td            int     // Tiempo de ataque de Diaboromon
	puntosDeVida  int     // Puntos de vida actuales
	cantidadDatos float64 // Cantidad de datos obtenidos
}

const (
	servidorPrimaryNode = "dist020.inf.santiago.usm.cl:50051"
	servidorDiaboromon  = "dist018.inf.santiago.usm.cl:50055"
)

var grpcServer *grpc.Server
var done chan bool = make(chan bool)

// Funcion para solicitar datos al Primary Node
func solicitarDatos(n *nodoTai) error {
	conn, err := grpc.Dial(servidorPrimaryNode, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("no se pudo conectar al Primary Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewServicioTaiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TaiRequest{Accion: "Solicitar Datos"}
	res, err := client.SolicitarDatos(ctx, req)
	if err != nil {
		return err
	}

	n.cantidadDatos = res.CantidadDatos
	log.Printf("[Nodo Tai] Cantidad de datos obtenidos del Primary Node: %.2f", n.cantidadDatos)

	if n.cantidadDatos >= float64(n.cd) {
		fmt.Println("[Nodo Tai] ¡Greymon y Garurumon tienen suficientes datos para evolucionar a Omegamon! Diaboromon ha sido derrotado.")
		go n.notificarDiaboromon(-1)
		fmt.Println("[Nodo Tai] Fin del sistema.")
		n.enviarNotificacionTerminacion()
		go finalizarNodoTai()
	} else {
		fmt.Println("[Nodo Tai] No hay suficientes datos para derrotar a Diaboromon, se recibe ataque.")
		restantes := n.puntosDeVida - 10
		if restantes < 0 {
			restantes = 0
		}
		go n.notificarDiaboromon(restantes)
		n.recibirAtaqueDeDiaboromon()
	}

	return nil
}

// Funcion para manejar el ataque de Diaboromon
func (n *nodoTai) recibirAtaqueDeDiaboromon() {
	n.puntosDeVida -= 10
	log.Printf("[Nodo Tai] ¡Diaboromon atacó! Vida restante de Greymon/Garurumon: %d", n.puntosDeVida)

	if n.puntosDeVida <= 0 {
		fmt.Println("[Nodo Tai] Greymon y Garurumon han sido derrotados. Diaboromon no pudo ser vencido.")
		n.enviarNotificacionTerminacion()
		go finalizarNodoTai()
	}
}

// Funcion para recibir un ataque de Diaboromon (gRPC)
func (n *nodoTai) Atacar(ctx context.Context, req *pb.TaiRequest) (*pb.TaiResponse, error) {
	log.Println("[Nodo Tai] Recibiendo ataque de Diaboromon estipulado por variable TD.")

	n.enviarEstadoAlPrimaryNode()
	n.recibirAtaqueDeDiaboromon()
	puntosDeVida := n.puntosDeVida
	if puntosDeVida < 0 {
		puntosDeVida = 0
	}

	return &pb.TaiResponse{
		Mensaje: strconv.Itoa(puntosDeVida),
	}, nil
}

// Funcion para enviar el estado de Diaboromon al Primary Node
func (s *nodoTai) enviarEstadoAlPrimaryNode() {
	conn, err := grpc.Dial(servidorPrimaryNode, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error al conectar con Primary Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewServicioTaiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.EstadoRequest{}
	_, err = client.EnviarEstadoDiaboromon(ctx, req)
	if err != nil {
		log.Fatalf("Error al enviar estado de Diaboromon: %v", err)
	}
}

// Funcion para enviar señal de terminacion al Primary Node
func (s *nodoTai) enviarNotificacionTerminacion() {
	conn, err := grpc.Dial(servidorPrimaryNode, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Primary Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewServicioTaiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TerminarRequest{}
	_, err = client.NotificarTerminacion(ctx, req)
	if err != nil {
		log.Fatalf("Error al notificar terminación al Primary Node: %v", err)
	}

}

// Función para notificar el estado final a Diaboromon
func (s *nodoTai) NotificarEstado(ctx context.Context, req *pb.EstadoTaiRequest) (*pb.EstadoTaiResponse, error) {
	var estado string
	var mensaje string

	// Determinar si ganó o perdió
	if s.puntosDeVida <= 0 {
		estado = "perdido"
		mensaje = "Greymon y Garurumon han sido derrotados."
	} else if s.cantidadDatos >= float64(s.cd) {
		estado = "ganado"
		mensaje = "¡Diaboromon ha sido derrotado!"
	}

	return &pb.EstadoTaiResponse{
		EstadoJuego: estado,
		Mensaje:     mensaje,
	}, nil
}

// Funcion para notificar a Diaboromon sobre un ataque simulado
func (n *nodoTai) notificarDiaboromon(vidaRestante int) {
	conn, err := grpc.Dial(servidorDiaboromon, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar a Diaboromon: %v", err)
	}
	defer conn.Close()

	client := pb.NewServicioDiaboromonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DiaboromonRequest{
		VidaRestante: int32(vidaRestante),
		Mensaje:      "Ataque simulado por Diaboromon.",
	}

	_, err = client.RecibirNotificacion(ctx, req)
	if err != nil {
		log.Fatalf("error al notificar a Diaboromon: %v", err)
	}

}

// Funciones auxiliares

// Funcion para leer las variables del sistema desde el archivo INPUT.txt
func (n *nodoTai) leerInputFile() error {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		if len(values) != 5 {
			return fmt.Errorf("el archivo INPUT.txt debe contener exactamente 5 valores separados por comas")
		}

		n.td, err = strconv.Atoi(values[2])
		if err != nil {
			return fmt.Errorf("error al parsear TD: %v", err)
		}

		n.cd, err = strconv.Atoi(values[3])
		if err != nil {
			return fmt.Errorf("error al parsear CD: %v", err)
		}

		n.vi, err = strconv.Atoi(values[4])
		if err != nil {
			return fmt.Errorf("error al parsear VI: %v", err)
		}
	}

	n.puntosDeVida = n.vi

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

func finalizarNodoTai() {
	time.Sleep(4 * time.Second)
	log.Println("[Nodo Tai] Finalizando nodo tai...")
	grpcServer.GracefulStop()
	done <- true
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("error al escuchar en el puerto 50054: %v", err)
	}

	grpcServer = grpc.NewServer()

	nodo := &nodoTai{}
	if err := nodo.leerInputFile(); err != nil {
		log.Fatalf("error al leer el archivo INPUT.txt: %v", err)
	}

	pb.RegisterServicioTaiServer(grpcServer, nodo)

	go func() {
		log.Println("[Nodo Tai] En ejecucion en puerto 50054.")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al servir: %v", err)
		}
	}()
	
	fmt.Println("Ingresa 1 en cualquier momento para solicitar cantidad de datos al Primary Node")

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanner.Scan()
			opcion := scanner.Text()
			if opcion != "1" {
				fmt.Println("Opción no válida. Intenta de nuevo.")
				continue
			} else {
				if err := solicitarDatos(nodo); err != nil {
					log.Fatalf("Error al solicitar datos al Primary Node: %v", err)
				}
			}
		}
	}()

	<-done
}
