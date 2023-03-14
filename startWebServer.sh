#!/bin/bash

echo "Atualizando as dependências..."
docker exec -it golang go mod tidy
echo "Iniciando o webserver..."
docker exec -it golang go run main.go http
