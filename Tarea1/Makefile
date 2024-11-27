# Nombre de las imágenes y los contenedores
IMAGE_NAME_1 = caravanas-app
IMAGE_NAME_2 = logistica-app
IMAGE_NAME_3 = finanzas-app
IMAGE_NAME_4 = clientes-app

CONTAINER_NAME_1 = caravanas_container
CONTAINER_NAME_2 = logistica_container
CONTAINER_NAME_3 = finanzas_container
CONTAINER_NAME_4 = clientes_container

PORT_CARAVANAS = 50051
PORT_LOGISTICA = 50052
PORT_FINANZAS = 50053
PORT_CLIENTES = 50054

docker-caravanas:
	@echo "Construyendo y ejecutando la aplicación de caravanas..."
	docker build -f DockerfileCaravanas -t $(IMAGE_NAME_1) .
	docker run --rm --name $(CONTAINER_NAME_1) -p $(PORT_CARAVANAS):$(PORT_CARAVANAS) $(IMAGE_NAME_1)

docker-logistica:
	@echo "Construyendo y ejecutando la aplicación de logística..."
	docker build -f DockerfileLogistica -t $(IMAGE_NAME_2) .
	docker run --rm --name $(CONTAINER_NAME_2) -p $(PORT_LOGISTICA):$(PORT_LOGISTICA) $(IMAGE_NAME_2)

docker-finanzas:
	@echo "Construyendo y ejecutando la aplicación de finanzas..."
	docker build -f DockerfileFinanzas -t $(IMAGE_NAME_3) .
	docker run --rm --name $(CONTAINER_NAME_3) -p $(PORT_FINANZAS):$(PORT_FINANZAS) $(IMAGE_NAME_3)

docker-clientes:
	@echo "Construyendo y ejecutando la aplicación de clientes..."
	docker build -f DockerfileClientes -t $(IMAGE_NAME_4) .
	docker run -it --rm --name $(CONTAINER_NAME_4) -p $(PORT_CLIENTES):$(PORT_CLIENTES) $(IMAGE_NAME_4)

clean:
	@echo "Eliminando contenedores, imágenes y recursos no utilizados..."
	docker system prune -f

stop-all:
	@echo "Deteniendo todos los contenedores..."
	docker stop $(shell docker ps -q)

remove-all:
	@echo "Eliminando todos los contenedores e imágenes..."
	docker ps -aq | xargs -r docker rm -f
	docker images -q | xargs -r docker rmi -f

help:
	@echo "Comandos disponibles:"
	@echo "  docker-logistica   - Construir y ejecutar la aplicación de logística"
	@echo "  docker-caravanas   - Construir y ejecutar la aplicación de caravanas"
	@echo "  docker-finanzas    - Construir y ejecutar la aplicación de finanzas"
	@echo "  docker-clientes    - Construir y ejecutar la aplicación de clientes"
	@echo "  clean              - Limpiar contenedores e imágenes no utilizados"
	@echo "  stop-all           - Detener todos los contenedores en ejecución"
	@echo "  remove-all         - Eliminar todos los contenedores e imágenes, incluso los que están en uso"
