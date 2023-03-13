#!/bin/bash

echo "Iniciando o banco de dados..."
docker exec -it golang rm -rf db.sqlite
docker exec -it golang touch db.sqlite
docker exec -it golang bash -c "sqlite3 db.sqlite 'CREATE TABLE products (id string PRIMARY KEY,name string,price float,status string);'"
