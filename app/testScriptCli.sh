#!/bin/bash

echo " - Atualizando as dependências..."
go mod tidy
echo " - Resetando o banco de dados..."
rm -rf db.sqlite
touch db.sqlite
sqlite3 db.sqlite 'CREATE TABLE products (id string PRIMARY KEY,name string,price float,status string);'
echo " - Criando os produtos de teste..."
sqlite3 db.sqlite 'INSERT INTO products (id,name,price,status) VALUES (1234,"product1",100.0,"disabled");'
sqlite3 db.sqlite 'INSERT INTO products (id,name,price,status) VALUES (5678,"product2",200.0,"enabled");'
echo " - Exibindo os produtos de teste..."
go run main.go cli -i=1234
go run main.go cli -i=5678
echo " - Criando um novo produto..."
go run main.go cli -a=create -n="Product CLI" -p=25.0
echo " - Ativando o product1..."
go run main.go cli -a=activate -i=1234
echo " - Desativando o product2..."
go run main.go cli -a=deactivate -i=5678
