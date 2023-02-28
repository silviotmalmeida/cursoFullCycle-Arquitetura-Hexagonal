#!/bin/bash

echo "Iniciando os testes..."
docker exec -it golang go test ./...
