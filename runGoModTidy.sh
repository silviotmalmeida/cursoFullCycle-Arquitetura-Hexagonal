#!/bin/bash

echo "Atualizando as dependĂȘncias..."
docker exec -it golang go mod tidy
