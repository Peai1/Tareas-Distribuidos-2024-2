package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"bufio"
	"strings"

	pb "Tarea2/protofiles"

	"google.golang.org/grpc"
)

type dataNodeServer struct {
	pb.UnimplementedServicioDataNodesServer
}

var grpcServer *grpc.Server 

func (s *dataNodeServer) GuardarData(ctx context.Context, req *pb.DataNodeRequest) (*pb.DataNodeResponse, error) {
	file, err := os.OpenFile("INFO2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	entry := fmt.Sprintf("%s,%s\n", req.Id, req.TipoDigimon)
	_, err = file.WriteString(entry)
	if err != nil {
		return nil, err
	}

	log.Printf("[Data Node 2] Informacion almacenada: %s", entry)
	return &pb.DataNodeResponse{Mensaje: "Datos almacenados en Data Node 2"}, nil
}

// Funcion para consultar datos acumulados enviada desde primary node
func (s *dataNodeServer) ConsultarAtributos(ctx context.Context, req *pb.DataNodeQuery) (*pb.DataNodeQueryResponse, error) {
    file, err := os.Open("INFO2.txt")
    if err != nil {
        return nil, fmt.Errorf("no se pudo abrir el archivo INFO.txt: %v", err)
    }
    defer file.Close()

    var atributos []string
    idsSolicitados := req.Ids
    idSet := make(map[string]bool)

    for _, id := range idsSolicitados {
        idSet[id] = true
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ",")
        if len(parts) != 2 {
            continue 
        }

        id := parts[0]
        atributo := parts[1]

        if _, exists := idSet[id]; exists {
            atributos = append(atributos, atributo)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error al leer el archivo INFO.txt: %v", err)
    }

    log.Printf("[Data Node 2] Solicitud de Primary Node recibida, mensaje enviado: %s", atributos)
    return &pb.DataNodeQueryResponse{Atributos: atributos}, nil
}

// Funcion para recibir la señal de terminación
func (s *dataNodeServer) NotificarTerminacion(ctx context.Context, req *pb.TerminarRequest) (*pb.TerminarResponse, error) {
	go func() {
		log.Println("[Data Node] Cerrando Data Node 2...")
		time.Sleep(2 * time.Second)
		grpcServer.GracefulStop()
	}()
	return &pb.TerminarResponse{Mensaje: "Data Node finalizando ejecucion."}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Fallo al escuchar: %v", err)
	}

	file, err := os.Create("INFO2.txt") 
    if err != nil {
        log.Fatalf("Error al reiniciar el archivo: %v", err)
    }
    file.Close()

	grpcServer = grpc.NewServer()
	pb.RegisterServicioDataNodesServer(grpcServer, &dataNodeServer{})

	log.Println("[Data Node 2] Data Node 2 en ejecucion en el puerto 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al servir: %v", err)
	}
}
