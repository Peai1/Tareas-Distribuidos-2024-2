package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	mrand "math/rand"
	"encoding/hex"
	"io"
	"bytes"

   
	pb "Tarea2/protofiles"

	"google.golang.org/grpc"
)

type servidorRegional struct {
	ps float64 // Probabilidad de sacrificio
	te float64     // Tiempo de espera entre envíos
}

const (
	servidorPrimaryNode = "dist020.inf.santiago.usm.cl:50051"
)

var aesKey = []byte("32-byte-long-key-for-AES-256!!!!")

// PKCS7 padding
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Funcion para cifrar un mensaje usando AES
func encriptarMensaje(message string) (string, error) {
	plaintext := []byte(message)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	// Rellenar el mensaje usando PKCS7
	plaintext = pkcs7Padding(plaintext, aes.BlockSize)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// Devolver el mensaje cifrado como hexadecimal
	return hex.EncodeToString(ciphertext), nil
}

// Funcion para enviar el estado de un Digimon al Primary Node
func enviarEstadoDigimon(client pb.ServicioRegionalesClient, nombre, tipo string, sacrificado bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Preparar el mensaje a enviar (nombre, tipo y sacrificado)
	estado := fmt.Sprintf("%s,%s,%v", nombre, tipo, sacrificado)

	// Cifrar el mensaje con AES
	ciphertext, err := encriptarMensaje(estado)
	if err != nil {
		log.Fatalf("Error al cifrar el mensaje: %v", err)
	}

	// Enviar el mensaje cifrado al Primary Node
	req := &pb.DigimonRequest{
		MensajeCifrado: ciphertext, // Enviar el mensaje cifrado
	}

	_, err = client.EnviarEstadoDigimon(ctx, req)
	if err != nil {
		log.Fatalf("Error al enviar estado del Digimon: %v", err)
	}

	if sacrificado {
		log.Printf("[Continente Folder] Estado enviado: %s SACRIFICADO", nombre)
	} else {
		log.Printf("[Continente Folder] Estado enviado: %s NO-SACRIFICADO", nombre)
	}
}

// Funcion para consultar el estado de terminación al Primary Node
func consultarEstadoTerminacion(client pb.ServicioRegionalesClient) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.EstadoRequest{}
	res, err := client.ConsultarEstadoTerminacion(ctx, req)
	if err != nil {
		log.Printf("[Continente Folder] Error al consultar estado de terminación: %v", err)
		return false
	}

	if res.Activo {
		log.Println("[Continente Folder] Señal de terminación recibida. Finalizando ejecución.")
		return true
	}
	return false
}


// Funciones auxiliares

// Función para leer las variables del sistema desde el archivo input.txt
func (s *servidorRegional) leerInputFile() error {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		// Leer la línea del archivo y separar los valores por comas
		line := scanner.Text()
		values := strings.Split(line, ",")
		if len(values) != 5 {
			return fmt.Errorf("el archivo INPUT.txt debe contener exactamente 5 valores separados por comas")
		}

		// Extraer PS y TE (los únicos que usamos en este servidor)
		s.ps, err = strconv.ParseFloat(values[0], 64)
		if err != nil {
			return fmt.Errorf("error al parsear PS: %v", err)
		}

		s.te, err = strconv.ParseFloat(values[1], 64) // Ahora TE puede manejar decimales
		if err != nil {
			return fmt.Errorf("error al parsear TE: %v", err)
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

func leerDigimonsDesdeArchivo(filename string) ([][2]string, error) {
    var digimons [][2]string

    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("error al abrir el archivo: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        partes := strings.Split(line, ",")
        if len(partes) != 2 {
            continue // Si no tiene exactamente 2 partes, saltar la línea
        }

        // Guardar el nombre y tipo de Digimon
        nombre := strings.TrimSpace(partes[0])
        tipo := strings.TrimSpace(partes[1])
        digimons = append(digimons, [2]string{nombre, tipo})
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error al leer el archivo: %v", err)
    }

    return digimons, nil
}

// Funcion para seleccionar un Digimon al azar
func seleccionarDigimon(digimons [][2]string) (string, string) {
    // Seleccionar un Digimon aleatorio de la lista
    mrand.Seed(time.Now().UnixNano())
    seleccion := digimons[mrand.Intn(len(digimons))]
    return seleccion[0], seleccion[1]
}

// Funcion para determinar si el Digimon debe ser sacrificado
func debeSerSacrificado(ps float64) bool {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Float64() < ps
}

// Ejecutar el ciclo de verificacion y envio de datos
func ejecutarCiclo(client pb.ServicioRegionalesClient, s *servidorRegional, digimons [][2]string) {
    lastSentTime := time.Now()

    for {
        // Verificar si se debe terminar la ejecución
        if consultarEstadoTerminacion(client) {
            log.Println("[Continente Folder] Finalizando ejecucion...")
            break
        }

        // Enviar datos cada TE segundos
        if time.Since(lastSentTime).Seconds() >= s.te {
            nombre, tipo := seleccionarDigimon(digimons) // Seleccionar Digimon aleatorio
            sacrificado := debeSerSacrificado(s.ps)
            enviarEstadoDigimon(client, nombre, tipo, sacrificado)
            lastSentTime = time.Now()
        }

        time.Sleep(100 * time.Millisecond) // Evitar el bucle intensivo
    }
}

func main() {
	// Leer los Digimons desde el archivo DIGIDATA.txt
    digimons, err := leerDigimonsDesdeArchivo("DIGIDATA.txt")
    if err != nil {
        log.Fatalf("Error al leer los Digimons desde el archivo: %v", err)
    }

	// Conectar al Primary Node
	conn, err := grpc.Dial(servidorPrimaryNode, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar al Primary Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewServicioRegionalesClient(conn)

	// Crear un servidor regional y leer las variables del archivo INPUT.txt
	s := &servidorRegional{}
	if err := s.leerInputFile(); err != nil {
		log.Fatalf("Error al leer el archivo INPUT.txt: %v", err)
	}

	log.Println("[Continente Folder] Esperando a que Diaboromon esté activo para empezar a enviar datos de Digimon...")
	for {
		// Consultar el estado de Diaboromon
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		req := &pb.EstadoRequest{}
		res, err := client.ConsultarEstadoDiaboromon(ctx, req)
		if err != nil {
			log.Printf("[Continente Folder] Error al consultar el estado de Diaboromon: %v", err)
		}

		// Verificar si Diaboromon está activo
		if res.Activo {
			log.Println("[Continente Folder] Diaboromon está activo. Comenzando a enviar datos.")
			break
		}
	}

	// Subir 6 datos de Digimon al iniciar
	for i := 0; i < 6; i++ {
		nombre, tipo := seleccionarDigimon(digimons)
		sacrificado := debeSerSacrificado(s.ps)
		enviarEstadoDigimon(client, nombre, tipo, sacrificado)
	}

	// Ejecutar la verificación de terminación y el envío de datos en paralelo
	ejecutarCiclo(client, s, digimons)
}
