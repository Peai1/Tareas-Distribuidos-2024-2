# Nombres de las imágenes y contenedores
IMAGE_NAME_BROKER = broker-app
IMAGE_NAME_SUPERVISOR1 = supervisor1-app
IMAGE_NAME_SUPERVISOR2 = supervisor2-app
IMAGE_NAME_JAYCE = jayce-app
IMAGE_NAME_SERVIDOR1 = servidor1-app
IMAGE_NAME_SERVIDOR2 = servidor2-app
IMAGE_NAME_SERVIDOR3 = servidor3-app

CONTAINER_NAME_BROKER = broker-container
CONTAINER_NAME_SUPERVISOR1 = supervisor1-container
CONTAINER_NAME_SUPERVISOR2 = supervisor2-container
CONTAINER_NAME_JAYCE = jayce-container
CONTAINER_NAME_SERVIDOR1 = servidor1-container
CONTAINER_NAME_SERVIDOR2 = servidor2-container
CONTAINER_NAME_SERVIDOR3 = servidor3-container

# Puertos de cada aplicación
PORT_BROKER = 4000
PORT_SUPERVISOR1 = 6000
PORT_SUPERVISOR2 = 6001
PORT_JAYCE = 6002
PORT_SERVIDOR1 = 5000
PORT_SERVIDOR2 = 5001
PORT_SERVIDOR3 = 5002

# Construcción y ejecución del Broker
docker-broker:
	@echo "Construyendo y ejecutando la aplicación del Broker..."
	docker build -f DockerfileBroker -t $(IMAGE_NAME_BROKER) .
	docker run -it --rm --name $(CONTAINER_NAME_BROKER) -p $(PORT_BROKER):$(PORT_BROKER) $(IMAGE_NAME_BROKER)

# Construcción y ejecución del Supervisor 1
docker-supervisor1:
	@echo "Construyendo y ejecutando la aplicación del Supervisor 1..."
	docker build -f DockerfileSupervisor1 -t $(IMAGE_NAME_SUPERVISOR1) .
	docker run -it --rm --name $(CONTAINER_NAME_SUPERVISOR1) -p $(PORT_SUPERVISOR1):$(PORT_SUPERVISOR1) $(IMAGE_NAME_SUPERVISOR1)

# Construcción y ejecución del Supervisor 2
docker-supervisor2:
	@echo "Construyendo y ejecutando la aplicación del Supervisor 2..."
	docker build -f DockerfileSupervisor2 -t $(IMAGE_NAME_SUPERVISOR2) .
	docker run -it --rm --name $(CONTAINER_NAME_SUPERVISOR2) -p $(PORT_SUPERVISOR2):$(PORT_SUPERVISOR2) $(IMAGE_NAME_SUPERVISOR2)

# Construcción y ejecución de Jayce
docker-jayce:
	@echo "Construyendo y ejecutando la aplicación de Jayce..."
	docker build -f DockerfileJayce -t $(IMAGE_NAME_JAYCE) .
	docker run -it --rm --name $(CONTAINER_NAME_JAYCE) -p $(PORT_JAYCE):$(PORT_JAYCE) $(IMAGE_NAME_JAYCE)

# Construcción y ejecución del Servidor 1
docker-servidor1:
	@echo "Construyendo y ejecutando la aplicación del Servidor 1..."
	docker build -f DockerfileServidor1 -t $(IMAGE_NAME_SERVIDOR1) .
	docker run -it --rm --name $(CONTAINER_NAME_SERVIDOR1) -p $(PORT_SERVIDOR1):$(PORT_SERVIDOR1) $(IMAGE_NAME_SERVIDOR1)

# Construcción y ejecución del Servidor 2
docker-servidor2:
	@echo "Construyendo y ejecutando la aplicación del Servidor 2..."
	docker build -f DockerfileServidor2 -t $(IMAGE_NAME_SERVIDOR2) .
	docker run -it --rm --name $(CONTAINER_NAME_SERVIDOR2) -p $(PORT_SERVIDOR2):$(PORT_SERVIDOR2) $(IMAGE_NAME_SERVIDOR2)

# Construcción y ejecución del Servidor 3
docker-servidor3:
	@echo "Construyendo y ejecutando la aplicación del Servidor 3..."
	docker build -f DockerfileServidor3 -t $(IMAGE_NAME_SERVIDOR3) .
	docker run -it --rm --name $(CONTAINER_NAME_SERVIDOR3) -p $(PORT_SERVIDOR3):$(PORT_SERVIDOR3) $(IMAGE_NAME_SERVIDOR3)

# Limpiar contenedores e imágenes no utilizados
clean:
	@echo "Eliminando contenedores e imágenes no utilizados..."
	docker system prune -f

# Detener todos los contenedores en ejecución
stop-all:
	@echo "Deteniendo todos los contenedores en ejecución..."
	docker stop $(shell docker ps -q)

# Eliminar todos los contenedores e imágenes
remove-all:
	@echo "Eliminando todos los contenedores e imágenes..."
	docker ps -aq | xargs -r docker rm -f
	docker images -q | xargs -r docker rmi -f

# Mostrar ayuda
help:
	@echo "Comandos disponibles:"
	@echo "  docker-broker          - Construir y ejecutar la aplicación del Broker"
	@echo "  docker-supervisor1     - Construir y ejecutar la aplicación del Supervisor 1"
	@echo "  docker-supervisor2     - Construir y ejecutar la aplicación del Supervisor 2"
	@echo "  docker-jayce           - Construir y ejecutar la aplicación de Jayce"
	@echo "  docker-servidor1       - Construir y ejecutar la aplicación del Servidor 1"
	@echo "  docker-servidor2       - Construir y ejecutar la aplicación del Servidor 2"
	@echo "  docker-servidor3       - Construir y ejecutar la aplicación del Servidor 3"
	@echo "  clean                  - Limpiar contenedores e imágenes no utilizados"
	@echo "  stop-all               - Detener todos los contenedores en ejecución"
	@echo "  remove-all             - Eliminar todos los contenedores e imágenes"
