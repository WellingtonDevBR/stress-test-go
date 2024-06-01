# Load Tester CLI em Go

Este projeto implementa um sistema CLI em Go para realizar testes de carga em um serviço web. Ele permite que o usuário forneça a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas. Ao final da execução dos testes, um relatório detalhado é gerado.

## Funcionalidades

- **Realização de Testes de Carga**: Envia requests HTTP para a URL especificada.
- **Concorrência**: Controla o número de chamadas simultâneas.
- **Relatório Detalhado**: Gera um relatório contendo:
  - Tempo total gasto na execução.
  - Quantidade total de requests realizados.
  - Quantidade de requests com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Requisitos

- Go 1.16+
- Docker (opcional, para execução em contêiner)

## Uso

### 1. Clonar o Repositório

```sh
git clone https://github.com/WellingtonDevBR/stress-test-go
go build -o load-tester main.go
./load-tester --url=http://google.com --requests=1000 --concurrency=10
```

### Executar com docker
```sh
docker build -t load-tester .
docker run --rm load-tester --url=http://google.com --requests=10 --concurrency=10
```

### Exemplo de saida
Tempo total gasto: 1.23s
Quantidade total de requests realizados: 1000
Quantidade de requests com status HTTP 200: 950
Distribuição de outros códigos de status HTTP:
HTTP 404: 30
HTTP 500: 20

Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.