# Base image
FROM golang:1.21

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos do projeto
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build do projeto
RUN go build -o main cmd/main.go

# Comando de execução
CMD ["./main"]
