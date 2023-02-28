#!/bin/bash

echo "Atualizando as dependências..."
docker exec -it golang go mod tidy
