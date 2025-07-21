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
