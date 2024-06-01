# Etapa de construção
FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .

# Instala as dependências
RUN go mod tidy

# Constrói o executável
RUN go build -o load-tester .

# Etapa final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/load-tester .
ENTRYPOINT ["./load-tester"]
