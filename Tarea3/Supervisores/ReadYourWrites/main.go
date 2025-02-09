// ReadYourWrites.main.go
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	messages "lab3/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverName string
var servePort string
var brokerAddress string = "dist020.inf.santiago.usm.cl:4000"

type Server struct {
	messages.UnimplementedMessageServiceServer
}

// Vars to keep modifiedSectors
var modifiedSectors map[string]*messages.VectorClock = make(map[string]*messages.VectorClock)
var serverSector map[string]string = make(map[string]string)

var fulcrumAddresses []string

// Aux function to get max
func max(a, b int32) int32 {
	if a >= b {
		return a
	}

	return b
}

func askAddres(cmd, sector, base, newBase string, value int32) string {
	conn, err := grpc.Dial(brokerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("No se pudo conectar a %s: %s", brokerAddress, err)
	}
	defer conn.Close()

	c := messages.NewMessageServiceClient(conn)

	request := &messages.Cmd{
		Cmd:       cmd,
		Sector:    sector,
		Base:      base,
		NuevaBase: newBase,
		Valor:     value,
	}

	responseBroker, err := c.AskAddress(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	return responseBroker.FulcrumAddress
}

func executeCommand(cmd string, sector string, base string, newBase string, value int32, address string) *messages.VectorClock {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("No se pudo conectar a %s: %s", address, err)
	}
	defer conn.Close()

	c := messages.NewMessageServiceClient(conn)

	request := &messages.Cmd{
		Cmd:       cmd,
		Sector:    sector,
		Base:      base,
		NuevaBase: newBase,
		Valor:     value,
	}

	responseBroker, err := c.Command(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	return responseBroker
}

func main() {

	// Inicializamos las direcciones de Fulcrum y broker
	fulcrumAddresses = append(fulcrumAddresses, os.Args[2], os.Args[3], os.Args[4])
	brokerAddress = os.Args[1]

	// Vars
	var input string

	var cmd string
	var sector string
	var base string
	newBase := ""
	value := int32(0)
	var tempValue int

	for {

		// Get input
		fmt.Println("Ingrese uno de los siguientes comandos:\n* AgregarProducto <nombre_region> <nombre_producto> <valor>\n* RenombrarProducto <nombre_region> <nombre_producto> <nuevo_nombre>\n* ActualizarValor <nombre_region> <nombre_producto> <nuevo_valor>\n* BorrarProducto <nombre_region> <nombre_producto>")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			input = scanner.Text()
		}
		if input == "" {
			break
		}

		// Extract data
		params := strings.Split(input, " ")

		if len(params) < 3 {
			fmt.Println("Ninguno de los parametros puede estar vacio.")
			continue
		}

		cmd = params[0]
		sector = params[1]
		base = params[2]

		if cmd == "AgregarProducto" {
			if len(params) == 4 {
				tempValue, _ = strconv.Atoi(params[3])
				value = int32(tempValue)
			} else {
				value = 0
			}
		} else if cmd == "RenombrarProducto" {
			if len(params) == 4 {
				newBase = params[3]
			} else {
				fmt.Println("Ninguno de los parametros puede estar vacio.")
				continue
			}
		} else if cmd == "ActualizarValor" {
			if len(params) == 4 {
				tempValue, _ = strconv.Atoi(params[3])
				value = int32(tempValue)
			} else {
				fmt.Println("Ninguno de los parametros puede estar vacio.")
				continue
			}
		}

		// If sector has prev address, use it. Else, ask for it and set new clock
		address := ""

		server, ok := serverSector[sector]
		if ok {
			address = server
		} else {
			address = askAddres(cmd, sector, base, newBase, value)
			serverSector[sector] = address
			modifiedSectors[sector] = &messages.VectorClock{X: 0, Y: 0, Z: 0}
		}

		clock := executeCommand(cmd, sector, base, newBase, value, address)

		modifiedSectors[sector].X = max(modifiedSectors[sector].X, clock.X)
		modifiedSectors[sector].Y = max(modifiedSectors[sector].Y, clock.Y)
		modifiedSectors[sector].Z = max(modifiedSectors[sector].Z, clock.Z)
	}
}
