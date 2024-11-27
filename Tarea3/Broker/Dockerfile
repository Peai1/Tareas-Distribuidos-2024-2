FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download

CMD ["go", "run", "Broker/main.go", ":4000", "dist017.inf.santiago.usm.cl:5000", "dist018.inf.santiago.usm.cl:5001", "dist019.inf.santiago.usm.cl:5002"]
