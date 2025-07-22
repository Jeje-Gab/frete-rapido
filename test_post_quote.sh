#!/bin/bash

API_URL="http://localhost:8080/api/quote"

echo "🚚 Enviando requisição de cotação para $API_URL..."

RESPONSE=$(curl -s -w "\nHTTP Status: %{http_code}\n" -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "recipient": {
      "address": {
        "zipcode": "01311000"
      }
    },
    "volumes": [
      {
        "category": 7,
        "amount": 1,
        "unitary_weight": 5,
        "price": 349,
        "sku": "abc-teste-123",
        "height": 0.2,
        "width": 0.2,
        "length": 0.2
      }
    ]
  }')

echo "$RESPONSE"
