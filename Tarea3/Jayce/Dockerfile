FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download

CMD ["go", "run", "Jayce/main.go", "dist020.inf.santiago.usm.cl:4000"]
