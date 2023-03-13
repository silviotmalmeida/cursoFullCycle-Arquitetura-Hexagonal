#!/bin/bash

echo "Iniciando a aplicação WEB..."

# criando product
docker exec -it golang bash -c "./testScriptWeb.sh"
