FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download

CMD ["go", "run", "Supervisores/ReadYourWrites/main.go", "dist020.inf.santiago.usm.cl:4000", "dist017.inf.santiago.usm.cl:5000", "dist018.inf.santiago.usm.cl:5001", "dist019.inf.santiago.usm.cl:5002"]
