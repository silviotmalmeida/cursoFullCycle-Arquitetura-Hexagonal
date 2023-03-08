package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/adapters/db"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"
	"github.com/stretchr/testify/require"
)

// suíte de testes unitários da classe adapter productDB

//// inicializando o módulo de sql nativo do golang
var Db *sql.DB

//// função de inicialização do BD
func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

//// função de criação das tabelas
func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
			);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

//// função de criação da massa de dados
func createProduct(db *sql.DB) {
	insert := `insert into products values("abc","Product Test",0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

//// se for executada uma busca por product no bd, os atributos devem ser iguais aos do objeto de origem
func TestProductDb_Get(t *testing.T) {
	// inicializando o BD
	setUp()
	// o BD será finalizado ao fim do teste
	defer Db.Close()

	// criando o productDb
	productDb := db.NewProductDb(Db)
	// realizando a consulta pelo registro incluído na massa de dados
	product, err := productDb.Get("abc")
	// não devem haver erros
	require.Nil(t, err)
	// os atributos devem ser iguais aos inseridos anteriormente
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

//// se um registro for criando/atualizado no bd, seus atributos devem ser iguais aos do objeto de origem
func TestProductDb_Save(t *testing.T) {
	// inicializando o BD
	setUp()
	// o BD será finalizado ao fim do teste
	defer Db.Close()

	// criando o productDb
	productDb := db.NewProductDb(Db)
	// criando o product novo
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	// salvando o novo registro
	productResult, err := productDb.Save(product)
	// não devem haver erros
	require.Nil(t, err)
	// os atributos devem ser iguais aos inseridos anteriormente
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	// alterando o product
	product.Status = "enabled"

	// atualizando o registro
	productResult, err = productDb.Save(product)
	// não devem haver erros
	require.Nil(t, err)
	// os atributos devem ser iguais aos inseridos anteriormente
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
