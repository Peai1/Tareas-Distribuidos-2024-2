package main

import (
	"bufio"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "Tarea2/protofiles"

	"google.golang.org/grpc"
)

type primaryNodeServer struct {
	pb.UnimplementedServicioRegionalesServer
	pb.UnimplementedServicioTaiServer
	mu               sync.Mutex
	idCounter        int
	diaboromonActivo bool
	terminado        bool
}

var grpcServer *grpc.Server
var aesKey = []byte("32-byte-long-key-for-AES-256!!!!")

const (
	dataNode1               = "dist017.inf.santiago.usm.cl:50052"
	dataNode2               = "dist019.inf.santiago.usm.cl:50053"
)

// PKCS7 unpadding
func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// Funcion para desencriptar un mensaje usando AES
func desencriptarMensaje(cipherText string) (string, error) {
	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	if len(cipherBytes) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := cipherBytes[:aes.BlockSize]
	cipherBytes = cipherBytes[aes.BlockSize:]

	if len(cipherBytes)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherBytes, cipherBytes)

	// Quitar el padding después de descifrar
	plainText := pkcs7UnPadding(cipherBytes)
	return string(plainText), nil
}

// Funcion para recibir el estado del Digimon desde los servidores regionales
func (s *primaryNodeServer) EnviarEstadoDigimon(ctx context.Context, req *pb.DigimonRequest) (*pb.DigimonResponse, error) {
	// Desencriptar el mensaje cifrado
	mensajeDesencriptado, err := desencriptarMensaje(req.MensajeCifrado)
	if err != nil {
		return nil, fmt.Errorf("error al desencriptar el mensaje cifrado: %v", err)
	}

	// Separar el mensaje desencriptado en sus partes
	partes := strings.Split(mensajeDesencriptado, ",")
	if len(partes) != 3 {
		return nil, fmt.Errorf("mensaje desencriptado con formato incorrecto")
	}

	nombreDigimon := partes[0]
	tipoDigimon := partes[1]
	sacrificado := partes[2] == "true"

	// Generar un ID único
	s.mu.Lock()
	s.idCounter++
	id := s.idCounter
	s.mu.Unlock()

	// Determinar el Data Node basado en la inicial del nombre del Digimon
	var numDataNode int
	var dataNode string
	if nombreDigimon[0] >= 'A' && nombreDigimon[0] <= 'M' {
		numDataNode = 1
		dataNode = dataNode1
	} else {
		numDataNode = 2
		dataNode = dataNode2
	}

	// Determinar el estado (Sacrificado o No Sacrificado)
	estado := "No Sacrificado"
	if sacrificado {
		estado = "Sacrificado"
	}

	log.Printf("[Primary Node] Digimon recibido: %s, Id: %d, Tipo: %s, Sacrificado: %v", nombreDigimon, s.idCounter, tipoDigimon, sacrificado)

	// Almacenar la información en el archivo INFO.txt
	if err := s.almacenarDatosEnArchivo(id, numDataNode, nombreDigimon, estado); err != nil {
		return nil, err
	}

	// Conectar al Data Node correspondiente
	conn, err := grpc.Dial(dataNode, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al Data Node %d: %v", numDataNode, err)
	}
	defer conn.Close()

	client := pb.NewServicioDataNodesClient(conn)
	dataReq := &pb.DataNodeRequest{
		Id:          strconv.Itoa(id), // ID generado
		TipoDigimon: tipoDigimon,      // Atributo del Digimon
	}

	// Enviar datos al Data Node
	_, err = client.GuardarData(context.Background(), dataReq)
	if err != nil {
		log.Fatalf("Error enviando datos al Data Node %d: %v", numDataNode, err)
	}

	return &pb.DigimonResponse{Mensaje: "Estado procesado por el Primary Node"}, nil
}

// ---------------------------- Funciones para obtener cantidad de datos ----------------------------

// Leer los IDs de los Digimon sacrificados desde INFO.txt
func (s *primaryNodeServer) leerIDSacrificados() ([]string, error) {
	file, err := os.Open("INFO.txt")
	if err != nil {
		return nil, fmt.Errorf("no se pudo abrir el archivo INFO.txt: %v", err)
	}
	defer file.Close()

	var idsSacrificados []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 4 {
			continue
		}

		id := parts[0]
		estado := parts[3]

		if estado == "Sacrificado" {
			idsSacrificados = append(idsSacrificados, id)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error al leer el archivo INFO.txt: %v", err)
	}

	log.Printf("[Primary Node] IDs de Digimon sacrificados: %v", idsSacrificados)
	return idsSacrificados, nil
}

// Solicitar atributos a un Data Node enviando los IDs
// Solicitar atributos a un Data Node enviando los IDs
func (s *primaryNodeServer) solicitarAtributosDataNode(numDataNode int, idsSacrificados []string) ([]string, error) {
	var dataNode string
	if numDataNode == 1 {
		dataNode = dataNode1
	} else {
		dataNode = dataNode2
	}

	conn, err := grpc.Dial(dataNode, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar al Data Node %d: %v", numDataNode, err)
	}
	defer conn.Close()

	client := pb.NewServicioDataNodesClient(conn)
	req := &pb.DataNodeQuery{Ids: idsSacrificados}

	// Solicitar atributos al Data Node
	res, err := client.ConsultarAtributos(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error al solicitar atributos al Data Node %d: %v", numDataNode, err)
	}

	return res.Atributos, nil
}

// Funcion para que el Nodo Tai solicite los datos acumulados
func (s *primaryNodeServer) SolicitarDatos(ctx context.Context, req *pb.TaiRequest) (*pb.TaiResponse, error) {
	// Leer el archivo INFO.txt para identificar los IDs de Digimon sacrificados
	idsSacrificados, err := s.leerIDSacrificados()
	if err != nil {
		return nil, err
	}

	if len(idsSacrificados) == 0 {
		log.Println("[Primary Node] Solicitud de [Nodo Tai] , mensaje enviado: No hay Digimon sacrificados.")
		return &pb.TaiResponse{CantidadDatos: 0, Mensaje: "No hay datos disponibles."}, nil
	}

	// Solicitar atributos a Data Node 1
	atributosNode1, err := s.solicitarAtributosDataNode(1, idsSacrificados)
	if err != nil {
		return nil, err
	}

	// Solicitar atributos a Data Node 2
	atributosNode2, err := s.solicitarAtributosDataNode(2, idsSacrificados)
	if err != nil {
		return nil, err
	}

	// Calcular la cantidad de datos acumulados según los atributos recibidos
	totalDatos := s.procesarDatos(atributosNode1) + s.procesarDatos(atributosNode2)

	log.Printf("[Primary Node] Solicitud de [Nodo Tai] , mensaje enviado, Datos acumulados: Total=%f", totalDatos)

	return &pb.TaiResponse{CantidadDatos: totalDatos, Mensaje: "Datos transferidos al Nodo Tai"}, nil
}

func (s *primaryNodeServer) procesarDatos(atributos []string) float64 {
	var totalDatos float64

	for _, atributo := range atributos {
		// Acumular puntos dependiendo del tipo de Digimon
		switch atributo {
		case "Vaccine":
			totalDatos += 3.0
		case "Data":
			totalDatos += 1.5
		case "Virus":
			totalDatos += 0.8
		}
	}

	return totalDatos
}

// --------------------------------------------------------------------------------------------------------------------

// Funcion para que los servidores regionales consulten si Diaboromon esta activo
func (s *primaryNodeServer) ConsultarEstadoDiaboromon(ctx context.Context, req *pb.EstadoRequest) (*pb.EstadoResponse, error) {
	// Devolver el estado de Diaboromon
	return &pb.EstadoResponse{
		Activo:  s.diaboromonActivo,
		Mensaje: "Estado de Diaboromon consultado",
	}, nil
}

// Funcion para recibir el estado de Diaboromon desde el Nodo Tai
func (s *primaryNodeServer) EnviarEstadoDiaboromon(ctx context.Context, req *pb.EstadoRequest) (*pb.EstadoResponse, error) {
	s.diaboromonActivo = true
	return &pb.EstadoResponse{
		Activo:  s.diaboromonActivo,
		Mensaje: "Estado de Diaboromon consultado",
	}, nil
}

// Funcion para recibir la señal de terminación desde el Nodo Tai
func (s *primaryNodeServer) NotificarTerminacion(ctx context.Context, req *pb.TerminarRequest) (*pb.TerminarResponse, error) {

	// Bloqueo para actualizar el estado de terminación
	s.mu.Lock()
	s.terminado = true
	s.mu.Unlock()

	// Enviar señal de terminación a los servidores data_node
	conn1, err := grpc.Dial(dataNode1, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al Data Node 1: %v", err)
	}
	defer conn1.Close()

	client1 := pb.NewServicioDataNodesClient(conn1)
	_, err = client1.NotificarTerminacion(context.Background(), &pb.TerminarRequest{})
	if err != nil {
		log.Fatalf("Error al notificar terminación al Data Node 1: %v", err)
	}

	conn2, err := grpc.Dial(dataNode2, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al Data Node 2: %v", err)
	}
	defer conn2.Close()

	client2 := pb.NewServicioDataNodesClient(conn2)
	_, err = client2.NotificarTerminacion(context.Background(), &pb.TerminarRequest{})
	if err != nil {
		log.Fatalf("Error al notificar terminación al Data Node 2: %v", err)
	}

	// Después de enviar la señal de terminación, detener el servidor de manera controlada
	go func() {
		time.Sleep(2 * time.Second) // Dar tiempo para que las señales se envíen
		log.Println("[Primary Node] Finalizando primary node...")
		grpcServer.GracefulStop() // Asegúrate de que grpcServer esté correctamente inicializado
	}()

	return &pb.TerminarResponse{Mensaje: "Señal de terminación recibida. Finalizando Primary Node."}, nil
}

// Funcion que los servidores regionales consultan para saber si deben finalizar
func (s *primaryNodeServer) ConsultarEstadoTerminacion(ctx context.Context, req *pb.EstadoRequest) (*pb.EstadoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Responder con el estado actual de 'terminado'
	return &pb.EstadoResponse{
		Activo:  s.terminado,
		Mensaje: "Consulta de terminación respondida.",
	}, nil
}

// Funciones auxiliares

// Funcion para almacenar los datos en INFO.txt
func (s *primaryNodeServer) almacenarDatosEnArchivo(id int, numDataNode int, nombre string, estado string) error {
	file, err := os.OpenFile("INFO.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escribir en el archivo INFO.txt
	entry := fmt.Sprintf("%d,%d,%s,%s\n", id, numDataNode, nombre, estado)
	_, err = file.WriteString(entry)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Fallo al escuchar: %v", err)
	}
	file, err := os.Create("INFO.txt") // Esto vacía el archivo si ya existe
	if err != nil {
		log.Fatalf("Error al reiniciar el archivo: %v", err)
	}
	file.Close()

	grpcServer = grpc.NewServer()
	s := &primaryNodeServer{idCounter: 0, diaboromonActivo: false, terminado: false} // Inicia el ID en un valor específico

	pb.RegisterServicioRegionalesServer(grpcServer, s)
	pb.RegisterServicioTaiServer(grpcServer, s)

	log.Println("Primary Node en ejecución en el puerto 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al servir: %v", err)
	}
}
