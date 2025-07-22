### ------------------------------------ Testes:

# 🚚 Desafio Frete Rápido - Backend com Go, Docker e PostgreSQL

Este projeto consiste em uma API RESTful desenvolvida em Go que simula cotações de frete (mockadas), armazena os dados em PostgreSQL, e oferece um endpoint para consulta de métricas. Abaixo está o passo a passo completo para executar o ambiente e realizar os testes localmente.

---

## ✅ 1. Pré-requisitos

Certifique-se de ter os seguintes componentes instalados:

- **Go** versão **1.24** ou superior
- **Docker** e **Docker Compose**
- Nenhum container PostgreSQL ativo usando a porta `5432`

### Verificações rápidas:

```bash
go version
docker --version
docker-compose --version
docker ps
```
## 🧹 2. Limpeza de ambiente (opcional)

```bash
docker-compose down -v
docker volume prune
```

## 🗂️ 3. Estrutura esperada do projeto
```bash
frete-rapido/
├── cmd/api/main.go
├── config/docker-compose.yml
├── Dockerfile
├── env/.env.development
├── migrations/01_up.sql
├── test_post_quote.sh
├── test_get_metrics.sh
```


## 4. Arquivo de variáveis de ambiente .env
no arquivo /env/.env.development ajustar as informações sensiveis
de acordo com o que foi passado via documentação do desafio "desafio-back-end-2.html",
substituindo as informaçôes em "XXXX" pelas reais!
```bash
FR_TOKEN=1d52XXXXXXXXXXXXXXXXXXXXXXXXXXX
FR_ENDPOINT=https://sp.freterapido.com/api/v3/quote/simulate
FR_CNPJ=25438XXXXXXXXX
FR_PLATFORM_CODE=5AKXXXXXX
FR_DISPATCHER_ZIP=29161376
```


## 🐳 5. Build e execução do projeto

Entre na pasta de configuração do projeto:
```bash
cd config
```
Execute o build completo com:
```bash
docker-compose up --build
```
Ao aparecer no terminal:
🚀 Servidor rodando em http://localhost:8080

o mesmo estará pronto para uso e teste


🧪 6. Testes com curl
Permitir execução dos scripts de teste, no diretorio raiz do projeto:
```bash
chmod +x test_post_quote.sh test_get_metrics.sh
```
e para executar os teste:
```bash
./test_post_quote.sh
./test_get_metrics.sh
```

OBS: sendo possivel utilizar uma ferramenta de testes como bruno ou
postman!
apenas direcionar as rotas para "http://localhost:8080/rota"



### Mais detalhes sobre a aplicação desenvolvida

# Desafio Frete Rápido - Backend em Golang

## Descrição do Projeto

Este projeto é uma API RESTful desenvolvida em Golang, que realiza cotações de frete simuladas (mockadas) a partir de dados enviados pelo usuário. Todas as cotações são salvas em um banco de dados PostgreSQL, permitindo posteriormente a consulta de métricas agregadas sobre as cotações realizadas.

## Tecnologias Utilizadas

- Golang
- Echo Framework
- PostgreSQL
- Docker e Docker Compose

## Justificativa da Arquitetura

### Mock de Transportadoras na Rota `/quote`

Como o ambiente de testes (sandbox) da API Frete Rápido não retorna valores reais de transportadoras de forma síncrona, a resposta do endpoint `/quote` é mockada com dois exemplos fixos de transportadoras, conforme o modelo do desafio técnico. Isso garante o funcionamento da aplicação, o salvamento dos dados e possibilita o uso da rota de métricas mesmo sem integração completa com a API real.

### Uso da Tabela `quote_requests`

Para estruturar corretamente o fluxo de cotação, foi criada a tabela `quote_requests`, que armazena cada requisição feita à rota `/quote`. Essa tabela contém informações como o CEP de destino e o horário da solicitação.

A tabela `quotes` armazena cada transportadora retornada em uma cotação, sempre vinculada ao campo `quote_request_id`. Isso permite identificar e agrupar corretamente todas as transportadoras ofertadas em uma mesma solicitação, além de permitir consultas precisas das últimas cotações, conforme solicitado pelo desafio.

## Estrutura do Banco de Dados

### quote_requests

- id (serial, PK)
- recipient_zipcode (varchar)
- created_at (timestamp)

### quotes

- id (serial, PK)
- quote_request_id (int, FK para quote_requests)
- carrier_name (text)
- service (text)
- deadline (int)
- price (numeric)
- created_at (timestamp)

## Endpoints

### POST `/quote`

- Recebe um JSON com os dados do destinatário e dos volumes.
- Realiza uma cotação simulada (mock).
- Salva a requisição e as opções de frete no banco de dados.

**Exemplo de resposta:**
```json
{
  "carrier": [
    {
      "name": "EXPRESSO FR",
      "service": "Rodoviário",
      "deadline": 3,
      "price": 17
    },
    {
      "name": "Correios",
      "service": "SEDEX",
      "deadline": 1,
      "price": 20.99
    }
  ]
}













