#!/bin/bash

API_URL="http://localhost:8080/api/metrics"

echo "📊 Consultando métricas em $API_URL..."

RESPONSE=$(curl -s -w "\nHTTP Status: %{http_code}\n" "$API_URL")

echo "$RESPONSE"
