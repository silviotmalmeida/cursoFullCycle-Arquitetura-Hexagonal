#!/bin/bash

echo "Iniciando a aplicação CLI..."

# criando product
docker exec -it golang bash -c "./testScriptCli.sh"
