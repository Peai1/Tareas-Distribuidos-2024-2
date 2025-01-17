package main

import (
	"context"
	messages "lab3/proto"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverName string = "Broker"
var servePort string = ":4000"       // Puerto en el que este broker listenea
var addressFulcrum1 string = "dist017.inf.santiago.usm.cl:5000" // Direccion del fulcrum1
var addressFulcrum2 string = "dist018.inf.santiago.usm.cl:5001" // ...
var addressFulcrum3 string = "dist019.inf.santiago.usm.cl:5002" // ...

func aleatorio(menor, mayor int) int {
	return rand.Intn(mayor-menor) + menor
}

type Server struct {
	messages.UnimplementedMessageServiceServer
}

func (s *Server) AskAddress(ctx context.Context, msg *messages.Cmd) (*messages.Address, error) {
	// Implementar aca la seleccion aleatoria del fulcrum a devolver

	aleatorio_interno := aleatorio(1, 4)

	fulcrumAddress := ""
	if aleatorio_interno == 1 {
		fulcrumAddress = addressFulcrum1
	} else if aleatorio_interno == 2 {
		fulcrumAddress = addressFulcrum2
	} else {
		fulcrumAddress = addressFulcrum3
	}

	log.Printf("[%s] Direccion del servidor solicitada, enviando %s...", serverName, fulcrumAddress)

	response := &messages.Address{FulcrumAddress: fulcrumAddress}

	return response, nil
}

func mergeRequest() {
	aleatorio_interno := aleatorio(1, 4)
	fulcrumAddress := ""
	fulcrumName := ""
	if aleatorio_interno == 1 {
		fulcrumAddress = addressFulcrum1
		fulcrumName = "Servidor1"
	} else if aleatorio_interno == 2 {
		fulcrumAddress = addressFulcrum2
		fulcrumName = "Servidor2"
	} else {
		fulcrumAddress = addressFulcrum3
		fulcrumName = "Servidor3"
	}

	log.Printf("[%s] Solicitando merge al servidor %s", serverName, fulcrumName)

	conn, err := grpc.Dial(fulcrumAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("No se pudo conectar a %s: %s", fulcrumAddress, err)
	}
	defer conn.Close()

	c := messages.NewMessageServiceClient(conn)

	request := &messages.AskMerge{FulcrumName: fulcrumName}

	_, err = c.RequestMerge(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

}

func (s *Server) GetSoldados(ctx context.Context, msg *messages.Info) (*messages.ValueInfo, error) {
	mergeRequest()

	aleatorio_interno := aleatorio(1, 4)

	fulcrumAddress := ""
	if aleatorio_interno == 1 {
		fulcrumAddress = addressFulcrum1
	} else if aleatorio_interno == 2 {
		fulcrumAddress = addressFulcrum2
	} else {
		fulcrumAddress = addressFulcrum3
	}

	conn, err := grpc.Dial(fulcrumAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("No se pudo conectar a %s: %s", fulcrumAddress, err)
	}
	defer conn.Close()

	c := messages.NewMessageServiceClient(conn)

	responseFulcrum, err := c.Soldados(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[%s] Cantidad de mercancia solicitada", serverName)

	return responseFulcrum, nil
}

func serve(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	messages.RegisterMessageServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", port, err)
	}
}

func main() {
	servePort = os.Args[1]
	addressFulcrum1 = os.Args[2]
	addressFulcrum2 = os.Args[3]
	addressFulcrum3 = os.Args[4]

	log.Printf("[%s] Iniciando...", serverName)

	go serve(servePort)

	for {
		time.Sleep(60 * time.Second)
		mergeRequest()
	}
}
