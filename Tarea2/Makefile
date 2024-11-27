# Nombres de las imágenes y los contenedores
IMAGE_NAME_1 = diaboromon-app
IMAGE_NAME_2 = primary-node-app
IMAGE_NAME_3 = nodo-tai-app
IMAGE_NAME_4 = data-node1-app
IMAGE_NAME_5 = data-node2-app
IMAGE_NAME_6 = continente-folder-app
IMAGE_NAME_7 = isla-file-app
IMAGE_NAME_8 = continente-server-app

CONTAINER_NAME_1 = diaboromon_container
CONTAINER_NAME_2 = primary_node_container
CONTAINER_NAME_3 = nodo_tai_container
CONTAINER_NAME_4 = data_node1_container
CONTAINER_NAME_5 = data_node2_container
CONTAINER_NAME_6 = continente_folder_container
CONTAINER_NAME_7 = isla_file_container
CONTAINER_NAME_8 = continente_server_container

# Puertos para cada contenedor
PORT_PRIMARY_NODE = 50051
PORT_DATANODE1 = 50052
PORT_DATANODE2 = 50053
PORT_NODO_TAI = 50054
PORT_DIABOROMON = 50055

# Construcción y ejecución de Diaboromon
docker-diaboromon:
	@echo "Construyendo y ejecutando la aplicación de Diaboromon..."
	docker build -f DockerfileDiaboromon -t $(IMAGE_NAME_1) .
	docker run --rm --name $(CONTAINER_NAME_1) -p $(PORT_DIABOROMON):$(PORT_DIABOROMON) $(IMAGE_NAME_1)

# Construcción y ejecución del Primary Node
docker-primary-node:
	@echo "Construyendo y ejecutando la aplicación del Primary Node..."
	docker build -f DockerfilePrimaryNode -t $(IMAGE_NAME_2) .
	docker run -it --rm --name $(CONTAINER_NAME_2) -p $(PORT_PRIMARY_NODE):$(PORT_PRIMARY_NODE) $(IMAGE_NAME_2)

# Construcción y ejecución del Nodo Tai
docker-nodo-tai:
	@echo "Construyendo y ejecutando la aplicación del Nodo Tai..."
	docker build -f DockerfileNodoTai -t $(IMAGE_NAME_3) .
	docker run -it --rm --name $(CONTAINER_NAME_3) -p $(PORT_NODO_TAI):$(PORT_NODO_TAI) $(IMAGE_NAME_3)

# Construcción y ejecución del Data Node 1
docker-datanode1:
	@echo "Construyendo y ejecutando la aplicación del Data Node 1..."
	docker build -f DockerfileDataNode1 -t $(IMAGE_NAME_4) .
	docker run -it --rm --name $(CONTAINER_NAME_4) -p $(PORT_DATANODE1):$(PORT_DATANODE1) $(IMAGE_NAME_4)

# Construcción y ejecución del Data Node 2
docker-datanode2:
	@echo "Construyendo y ejecutando la aplicación del Data Node 2..."
	docker build -f DockerfileDataNode2 -t $(IMAGE_NAME_5) .
	docker run -it --rm --name $(CONTAINER_NAME_5) -p $(PORT_DATANODE2):$(PORT_DATANODE2) $(IMAGE_NAME_5)

# Construcción y ejecución del Nodo Regional "Continente Folder"
docker-continente-folder:
	@echo "Construyendo y ejecutando la aplicación del Nodo Regional 'Continente Folder'..."
	docker build -f DockerfileContinenteFolder -t $(IMAGE_NAME_6) .
	docker run -it --rm --name $(CONTAINER_NAME_6) $(IMAGE_NAME_6)

# Construcción y ejecución del Nodo Regional "Isla File"
docker-isla-file:
	@echo "Construyendo y ejecutando la aplicación del Nodo Regional 'Isla File'..."
	docker build -f DockerfileIslaFile -t $(IMAGE_NAME_7) .
	docker run --rm --name $(CONTAINER_NAME_7) $(IMAGE_NAME_7)

# Construcción y ejecución del Nodo Regional "Continente Server"
docker-continente-server:
	@echo "Construyendo y ejecutando la aplicación del Nodo Regional 'Continente Server'..."
	docker build -f DockerfileContinenteServer -t $(IMAGE_NAME_8) .
	docker run --rm --name $(CONTAINER_NAME_8) $(IMAGE_NAME_8)

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
	@echo "  docker-diaboromon        - Construir y ejecutar la aplicación de Diaboromon"
	@echo "  docker-primary-node      - Construir y ejecutar la aplicación del Primary Node"
	@echo "  docker-nodo-tai          - Construir y ejecutar la aplicación del Nodo Tai"
	@echo "  docker-datanode1         - Construir y ejecutar la aplicación del Data Node 1"
	@echo "  docker-datanode2         - Construir y ejecutar la aplicación del Data Node 2"
	@echo "  docker-continente-folder - Construir y ejecutar la aplicación del Nodo Regional 'Continente Folder'"
	@echo "  docker-isla-file         - Construir y ejecutar la aplicación del Nodo Regional 'Isla File'"
	@echo "  docker-continente-server - Construir y ejecutar la aplicación del Nodo Regional 'Continente Server'"
	@echo "  clean                    - Limpiar contenedores e imágenes no utilizados"
	@echo "  stop-all                 - Detener todos los contenedores en ejecución"
	@echo "  remove-all               - Eliminar todos los contenedores e imágenes"
