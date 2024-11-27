package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	pb "Tarea1/protofiles"

	"google.golang.org/grpc"
)

const (
	costoReintento         = 100
	porcentajePrioritario  = 0.30
	maxIntentosOstronitas  = 3
	maxIntentosGrineer     = 2
	porcentajeExito        = 0.85
	maxPaquetesPorCaravana = 2
	tiempoEntrega          = 3
	tiempoEspera           = 6
)

const (
	serverLogistica = "dist018.inf.santiago.usm.cl:50052"
	serverCaravana  = ":50051"
)

// Caravana
type caravanServer struct {
	pb.UnimplementedCaravanServiceServer
	mu                *sync.Mutex     // Mutex para sincronizar los print en consola
	wg                *sync.WaitGroup // WaitGroup para esperar a que todas las caravanas terminen
	done              chan bool       // Canal para señalizar cuando todas las caravanas terminen
	paquetesQueue     []*pb.PackageRequest
	paquetesRecibidos int
	contadorRonda     map[int]int // Contador de rondas para cada caravana
	resultados        []*pb.DeliveryResult
}

// Manejador para recibir instrucciones de reparto (CaravanService)
func (s *caravanServer) SendDeliveryInstructions(ctx context.Context, req *pb.DeliveryInstructionRequest) (*pb.DeliveryInstructionResponse, error) {
	// Convertir el timestamp a un formato legible

	for _, p := range req.Packages {
		tiempoFormateado := time.Unix(p.Timestamp, 0).Format(time.RFC3339)
		log.Printf("Recibido pedido %s Tipo: %s, Time: %s, Destino: %s, Warframe: %s, Recurso: %s, Valor: %d", p.IdPaquete, p.Tipo, tiempoFormateado, p.Destino, p.Escolta, p.Nombre, p.Valor)
		// Añadir los paquetes a la cola de espera
		s.paquetesQueue = append(s.paquetesQueue, p)
	}

	// Esperar a segundo paquete antes de simular la entrega
	time.Sleep(tiempoEspera * time.Second)

	s.mu.Lock()
	defer s.mu.Unlock()

	// Simular la entrega de los paquetes en las caravanas y acumular resultados
	if len(s.paquetesQueue) != 0 {
		s.simularEntrega(s.paquetesQueue)
		return &pb.DeliveryInstructionResponse{
			Results: s.resultados, // Retornar los resultados de las entregas
		}, nil
	}

	return &pb.DeliveryInstructionResponse{
		Results: nil,
	}, nil
}

// Simular la entrega de paquetes dividiendo el trabajo entre tres caravanas y acumulando resultados
func (s *caravanServer) simularEntrega(paquetes []*pb.PackageRequest) {
	var wg sync.WaitGroup

	// Filtrar paquetes de Ostronitas y Prioritarios
	var paquetesOstronitas []*pb.PackageRequest
	var paquetesPrioritarios []*pb.PackageRequest
	var paquetesNormales []*pb.PackageRequest

	for _, p := range paquetes {
		if p.Tipo == "Ostronitas" {
			paquetesOstronitas = append(paquetesOstronitas, p)
		} else if p.Tipo == "Prioritario" {
			paquetesPrioritarios = append(paquetesPrioritarios, p)
		} else {
			paquetesNormales = append(paquetesNormales, p)
		}
	}

	// Ordenar paquetes por valor (prioridad para mayores ingresos)
	sort.SliceStable(paquetesOstronitas, func(i, j int) bool {
		return paquetesOstronitas[i].Valor > paquetesOstronitas[j].Valor
	})
	sort.SliceStable(paquetesPrioritarios, func(i, j int) bool {
		return paquetesPrioritarios[i].Valor > paquetesPrioritarios[j].Valor
	})
	sort.SliceStable(paquetesNormales, func(i, j int) bool {
		return paquetesNormales[i].Valor > paquetesNormales[j].Valor
	})

	for len(paquetesOstronitas) > 0 || len(paquetesPrioritarios) > 0 || len(paquetesNormales) > 0 {
		wg.Add(3)

		// Asignar paquetes a las caravanas prioritarias
		for j := 0; j < 2; j++ {
			var paquetesAsignados []*pb.PackageRequest

			for len(paquetesAsignados) < maxPaquetesPorCaravana && len(paquetesOstronitas) > 0 {
				paquetesAsignados = append(paquetesAsignados, paquetesOstronitas[0])
				paquetesOstronitas = paquetesOstronitas[1:]
			}
			for len(paquetesAsignados) < maxPaquetesPorCaravana && len(paquetesPrioritarios) > 0 {
				paquetesAsignados = append(paquetesAsignados, paquetesPrioritarios[0])
				paquetesPrioritarios = paquetesPrioritarios[1:]
			}

			if len(paquetesAsignados) > 0 {
				s.simularCaravanaPrioritaria(j+1, paquetesAsignados, &wg)
			} else {
				wg.Done()
			}
		}

		// La caravana general se encarga del resto de los suministros (incluye paquetes prioritarios y normales)
		if len(paquetesPrioritarios) > 0 || len(paquetesNormales) > 0 {
			var paquetesAsignados []*pb.PackageRequest
			for len(paquetesAsignados) < maxPaquetesPorCaravana && len(paquetesPrioritarios) > 0 {
				paquetesAsignados = append(paquetesAsignados, paquetesPrioritarios[0])
				paquetesPrioritarios = paquetesPrioritarios[1:]
			}
			for len(paquetesAsignados) < maxPaquetesPorCaravana && len(paquetesNormales) > 0 {
				paquetesAsignados = append(paquetesAsignados, paquetesNormales[0])
				paquetesNormales = paquetesNormales[1:]
			}
			if len(paquetesAsignados) > 0 {
				s.simularCaravanaGeneral(3, paquetesAsignados, &wg)
			} else {
				wg.Done()
			}
		} else {
			wg.Done()
		}

		wg.Wait()
	}

}

// Simulación de las caravanas prioritarias entregando paquetes
func (s *caravanServer) simularCaravanaPrioritaria(caravanID int, paquetes []*pb.PackageRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	// Intentar entregar los paquetes priorizando los de mayor valor
	s.contadorRonda[caravanID]++
	s.entregarPaquetes(caravanID, paquetes, true)
}

// Simulación de la caravana general entregando paquetes
func (s *caravanServer) simularCaravanaGeneral(caravanID int, paquetes []*pb.PackageRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	// Priorizamos los paquetes prioritarios sobre los normales dentro del resto de los paquetes
	s.contadorRonda[caravanID]++
	sort.SliceStable(paquetes, func(i, j int) bool {
		if paquetes[i].Tipo == "Prioritario" && paquetes[j].Tipo != "Prioritario" {
			return true
		}
		return paquetes[i].Valor > paquetes[j].Valor
	})

	// Intentar entregar los paquetes
	s.entregarPaquetes(caravanID, paquetes, false)
}

// Función para simular la entrega de paquetes con reintentos y manejar los fallos correctamente
func (s *caravanServer) entregarPaquetes(caravanID int, paquetes []*pb.PackageRequest, esPrioritaria bool) {

	caravanaTipo := "General"
	if esPrioritaria {
		caravanaTipo = "Prioritaria"
		//log.Printf("Caravana Prioritaria %d esta comenzando la entrega de %d paquetes", caravanID, len(paquetes))
	} else {
		//log.Printf("Caravana General %d esta comenzando la entrega de %d paquetes", caravanID, len(paquetes))
	}

	s.paquetesRecibidos += len(paquetes)

	client := s.connectToLogistica()
	intentosPaquetes := make(map[string]int32)
	paquetesEntregados := make(map[string]bool)

	// Inicializar todos los intentos en 0 y los paquetes como no entregados
	for _, paquete := range paquetes {
		intentosPaquetes[paquete.IdPaquete] = 0
		paquetesEntregados[paquete.IdPaquete] = false
	}

	// Seguir entregando hasta que todos los paquetes se entreguen o se agoten los intentos
	quedanPendientes := true
	for quedanPendientes {
		quedanPendientes = false

		for i, paquete := range paquetes {
			indexPaquete := i + 1
			if paquetesEntregados[paquete.IdPaquete] {
				continue // Si el paquete ya fue entregado, pasamos al siguiente
			}

			intentos := intentosPaquetes[paquete.IdPaquete]

			// Determinar los intentos maximos segun la facción
			maxIntentos := maxIntentosGrineer
			if paquete.Tipo == "Ostronitas" {
				maxIntentos = maxIntentosOstronitas
			}

			// Si el paquete ha alcanzado su límite de intentos, dejar de intentar entregarlo
			if intentos >= int32(maxIntentos) {
				// Retorna a logistica.go que el paquete fue no entregado
				_, err := client.UpdateDelivery(context.Background(), &pb.DeliveryUpdateRequest{
					IdPaquete:   paquete.IdPaquete,
					Seguimiento: "TC" + paquete.IdPaquete,
					Tipo:        paquete.Tipo,
					Valor:       paquete.Valor,
					Intentos:    int32(intentos),
					Estado:      "No Entregado",
				})
				if err != nil {
					log.Printf("error al enviar actualizacion a logistica: %v", err)
				}
				fileName := fmt.Sprintf("Caravana_%s_%d.txt", caravanaTipo, caravanID)
				file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println("Error al abrir o crear el archivo:", err)
					return
				}
				defer file.Close() // Asegurarse de cerrar el archivo

				// Crear un buffer para escribir en el archivo
				writer := bufio.NewWriter(file)
				new_linea := fmt.Sprintf("Id: %s, Tipo: %s , Valor: %d, Escolta: %s, Destino: %s, Intentos: %d, Fecha entrega: %d \n",
					paquete.IdPaquete, paquete.Tipo, paquete.Valor, paquete.Escolta, paquete.Destino, intentos+1, 0)

				// Escribir en el archivo
				_, err = writer.WriteString(new_linea)
				if err != nil {
					fmt.Println("Error al escribir en el archivo:", err)
					return
				}
				writer.Flush()
				log.Printf("Caravana %s %d: El paquete %s (Tipo: %s) sera devuelto a Cetus tras %d intentos fallidos (paquete %d/%d total, ronda caravana: %d)", caravanaTipo, caravanID, paquete.IdPaquete, paquete.Tipo, intentos, indexPaquete, len(paquetes), s.contadorRonda[caravanID])
				s.evaluarResultado(paquete, int(intentos), false)
				paquetesEntregados[paquete.IdPaquete] = true
				continue // Pasar al siguiente paquete
			}

			// Intentar entregar el paquete
			entregado := intentarEntrega()

			// solamente si es primera vez que se intenta entregar el paquete, se envia la actualizacion a logistica que fue Enviado
			if intentos == 0 {
				_, err := client.UpdateDelivery(context.Background(), &pb.DeliveryUpdateRequest{
					IdPaquete:   paquete.IdPaquete,
					Seguimiento: "TC" + paquete.IdPaquete,
					Tipo:        paquete.Tipo,
					Valor:       paquete.Valor,
					Intentos:    int32(intentos + 1),
					Estado:      "Enviado",
				})
				if err != nil {
					log.Printf("error al enviar actualizacion a logistica: %v", err)
				}
			}

			// Simular el tiempo de entrega
			time.Sleep(tiempoEntrega * time.Second)

			if entregado {
				// Si el paquete ha sido entregado, actualizamos su estado a "Entregado"
				_, err := client.UpdateDelivery(context.Background(), &pb.DeliveryUpdateRequest{
					IdPaquete:   paquete.IdPaquete,
					Seguimiento: "TC" + paquete.IdPaquete,
					Tipo:        paquete.Tipo,
					Valor:       paquete.Valor,
					Intentos:    int32(intentos + 1),
					Estado:      "Entregado",
				})
				if err != nil {
					log.Printf("error al enviar actualizacion a logistica: %v", err)
				}
				fileName := fmt.Sprintf("Caravana_%s_%d.txt", caravanaTipo, caravanID)
				file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println("Error al abrir o crear el archivo:", err)
					return
				}
				defer file.Close() // Asegurarse de cerrar el archivo

				// Crear un buffer para escribir en el archivo
				writer := bufio.NewWriter(file)
				tiempoFormateado := time.Now().Format(time.RFC3339)
				linea := fmt.Sprintf("Id: %s, Tipo: %s , Valor: %d, Escolta: %s, Destino: %s, Intentos: %d, Fecha entrega: %s \n",
					paquete.IdPaquete, paquete.Tipo, paquete.Valor, paquete.Escolta, paquete.Destino, intentos+1, tiempoFormateado)

				// Escribir en el archivo
				_, err = writer.WriteString(linea)
				if err != nil {
					fmt.Println("Error al escribir en el archivo:", err)
					return
				}
				writer.Flush()
				log.Printf("Caravana %s %d entrego exitosamente el paquete %s (Tipo: %s) a %s en %d intento(s) (paquete %d/%d total, ronda caravana: %d)", caravanaTipo, caravanID, paquete.IdPaquete, paquete.Tipo, paquete.Destino, intentos+1, indexPaquete, len(paquetes), s.contadorRonda[caravanID])
				s.evaluarResultado(paquete, int(intentos)+1, true)
				paquetesEntregados[paquete.IdPaquete] = true // Marcamos el paquete como entregado
			} else {
				// Si la entrega falla, actualizamos el estado del paquete a "Enviado" y aumentamos el número de intentos
				if intentos != 0 {
					_, err := client.UpdateDelivery(context.Background(), &pb.DeliveryUpdateRequest{
						IdPaquete:   paquete.IdPaquete,
						Seguimiento: "TC" + paquete.IdPaquete,
						Tipo:        paquete.Tipo,
						Valor:       paquete.Valor,
						Intentos:    int32(intentos + 1),
						Estado:      "Enviado",
					})
					if err != nil {
						log.Printf("error al enviar actualizacion a logistica: %v", err)
					}
				}
				intentosPaquetes[paquete.IdPaquete] = intentos + 1 // Incrementamos el número de intentos
				log.Printf("Caravana %s %d fallo el intento %d para el paquete %s (Tipo: %s) (paquete %d/%d total, ronda caravana: %d)", caravanaTipo, caravanID, intentos+1, paquete.IdPaquete, paquete.Tipo, indexPaquete, len(paquetes), s.contadorRonda[caravanID])
				quedanPendientes = true // Aun quedan paquetes por entregar
			}
		}
	}

	// Se eliminan de la cola los paquetes procesados
	for _, paquete := range paquetes {
		for i, p := range s.paquetesQueue {
			if p.IdPaquete == paquete.IdPaquete {
				s.paquetesQueue = append(s.paquetesQueue[:i], s.paquetesQueue[i+1:]...)
				break
			}
		}
	}

	log.Printf("Caravana %d ha completado la entrega de todos los paquetes asignados esta ronda", caravanID)

	// Leer la primera línea
	file, err := os.Open("solicitudes.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	paquetesEsperados := 0
	// Leer línea por línea
	if scanner.Scan() {
		linea := scanner.Text()
		parametros := strings.Split(linea, ",")
		fmt.Sscanf(parametros[0], "%d", &paquetesEsperados)

	}
	if s.paquetesRecibidos >= paquetesEsperados {
		log.Println("Todos los paquetes han sido procesados. Cerrando el servidor...")

		s.done <- true
	}
}

// Evaluar el resultado de la entrega y calcular ganancias/pérdidas
func (s *caravanServer) evaluarResultado(paquete *pb.PackageRequest, intentos int, entregado bool) {
	var valorGanancia int32 = 0

	if entregado {
		valorGanancia = paquete.Valor
	} else {
		if paquete.Tipo == "Ostronitas" {
			valorGanancia = paquete.Valor // Se queda con el suministro
		} else if paquete.Tipo == "Prioritario" {
			valorGanancia = int32(float32(paquete.Valor) * porcentajePrioritario) // 30% de ganancia
		} else if paquete.Tipo == "Normal" {
			valorGanancia = 0 // Se pierde el suministro
		}
	}

	// Calcular los créditos restantes basados en los intentos realizados
	valorGanancia = valorGanancia - int32(intentos)*100

	s.resultados = append(s.resultados, &pb.DeliveryResult{
		IdPaquete: paquete.IdPaquete,
		Exito:     entregado,
		Intentos:  int32(intentos),
		Valor:     valorGanancia,
	})

}

// Función auxiliar para simular una entrega con una probabilidad de éxito del 85%
func intentarEntrega() bool {
	return rand.Float32() < porcentajeExito
}

// Inicializa la conexión con logistica
func (s *caravanServer) connectToLogistica() pb.CaravanServiceClient {
	conn, err := grpc.Dial(serverLogistica, grpc.WithInsecure()) // Asegúrate de usar la dirección y puerto correctos del servidor logístico
	if err != nil {
		log.Fatalf("error al conectar con el servidor de logistica: %v", err)
	}

	return pb.NewCaravanServiceClient(conn)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Inicializar el generador de números aleatorios

	var mu sync.Mutex     // Crear un mutex para sincronizar los print
	var wg sync.WaitGroup // Crear un WaitGroup para esperar a que todas las caravanas terminen

	// Canal para señalizar el cierre del servidor
	done := make(chan bool)

	lis, err := net.Listen("tcp", serverCaravana)
	if err != nil {
		log.Fatalf("error al iniciar el servidor de caravana: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCaravanServiceServer(grpcServer, &caravanServer{
		mu:            &mu,
		done:          done,
		wg:            &wg,
		contadorRonda: make(map[int]int),
		resultados:    nil,
	})

	// Iniciar el servidor en una goroutine
	go func() {
		log.Println("Servidor de caravana en ejecución en el puerto 50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("error al ejecutar el servidor de caravana: %v", err)
		}
	}()

	// Esperar a que todas las caravanas terminen su entrega y luego cerrar el servidor
	<-done
	log.Println("Todas las caravanas han terminado. Cerrando caravanas...")
	grpcServer.GracefulStop()
}
