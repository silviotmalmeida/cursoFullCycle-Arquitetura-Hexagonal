#!/bin/bash

echo " - Resetando o banco de dados..."
rm -rf db.sqlite
touch db.sqlite
sqlite3 db.sqlite 'CREATE TABLE products (id string PRIMARY KEY,name string,price float,status string);'
echo " - Criando os produtos de teste..."
sqlite3 db.sqlite 'INSERT INTO products (id,name,price,status) VALUES (1234,"product1",100.0,"disabled");'
sqlite3 db.sqlite 'INSERT INTO products (id,name,price,status) VALUES (5678,"product2",0.0,"enabled");'
echo " - Exibindo os produtos de teste..."
curl http://localhost:9000/product/1234
curl http://localhost:9000/product/5678
echo " - Ativando o product1..."
curl http://localhost:9000/product/1234/enable
echo " - Desativando o product2..."
curl http://localhost:9000/product/5678/disable
