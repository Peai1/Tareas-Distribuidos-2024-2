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

type diaboromonServer struct {
	pb.UnimplementedServicioDiaboromonServer    
	td                                       int
}

const (
	servidorTai = "dist020.inf.santiago.usm.cl:50054"
)

var grpcServer *grpc.Server
var done chan bool = make(chan bool) 

// Funcion para realizar un ataque al Nodo Tai
func (d *diaboromonServer) atacarNodoTai(client pb.ServicioTaiClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.TaiRequest{Accion: "Ataque Diaboromon"}
	res, err := client.Atacar(ctx, req)
	if err != nil {
		log.Fatalf("error al atacar al Nodo Tai: %v", err)
	}

	log.Printf("[Diaboromon] Ataque realizado al Nodo Tai. Vida restante: %s", res.Mensaje)
	if res.Mensaje == "0" {
		log.Println("[Diaboromon] Greymon y Garurumon han sido derrotados. Diaboromon no pudo ser vencido.")
		done <- true
	} 
}

// Funcion para recibir notificaciones desde Nodo Tai
func (d *diaboromonServer) RecibirNotificacion(ctx context.Context, req *pb.DiaboromonRequest) (*pb.DiaboromonResponse, error) {
	if req.VidaRestante == -1 {
		go func() {
			log.Println("[Diaboromon] Greymon y Garurumon consultan datos y tienen los puntos suficientes para evolucionar. Diaboromon ha sido derrotado.")
			time.Sleep(500 * time.Millisecond)
			done <- true
		}()
	} else if req.VidaRestante == 0 {
		go func() {
			log.Println("[Diaboromon] Nodo Tai consultan datos pero no tienen los puntos suficientes para evolucionar y quedan con 0 de vida. Diaboromon gana la batalla.")
			time.Sleep(500 * time.Millisecond)
			done <- true
		}()
	} else {
		log.Printf("[Diaboromon] Nodo Tai consultan datos pero no tiene los puntos suficientes. Recibe ataque y queda con %d de vida", req.VidaRestante)
	}
	return &pb.DiaboromonResponse{Mensaje: "Ataque recibido."}, nil
}

// Funciones auxiliares

// Funcion para leer las variables del sistema desde el archivo INPUT.txt
func (d *diaboromonServer) leerInputFile() error {
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

		d.td, err = strconv.Atoi(values[2])
		if err != nil {
			return fmt.Errorf("error al parsear TD: %v", err)
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	log.Printf("[Diaboromon] Se realiza ataque cada %d segundos", d.td)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50055") 
	if err != nil {
		log.Fatalf("error al escuchar en el puerto 50055: %v", err)
	}

	grpcServer = grpc.NewServer()
	diaboromon := &diaboromonServer{}

	if err := diaboromon.leerInputFile(); err != nil {
		log.Fatalf("error al leer el archivo INPUT.txt: %v", err)
	}

	pb.RegisterServicioDiaboromonServer(grpcServer, diaboromon)

	go func() {
		log.Println("[Diaboromon] En ejecucion en el puerto 50055.")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al servir: %v", err)
		}
	}()

	conn, err := grpc.Dial(servidorTai, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("no se pudo conectar al Nodo Tai: %v", err)
	}
	defer conn.Close()

	client := pb.NewServicioTaiClient(conn)

	go func() {
		for {
			select {
			case <-done:
				log.Println("[Diaboromon] Terminando los ataques y finalizando el servidor.")
				grpcServer.GracefulStop()
				return
			default:
			select {
			case <-done:
				log.Println("[Diaboromon] Señal de cierre recibida, cancelando ataque.")
				grpcServer.GracefulStop()
				return
			default:
				log.Println("[Diaboromon] Preparando para atacar al Nodo Tai...")
					diaboromon.atacarNodoTai(client)
					time.Sleep(time.Duration(diaboromon.td) * time.Second)
				}
			}
		}
	}()

	<-done
	log.Println("[Diaboromon] Cerrando Diaboromon..")
}
