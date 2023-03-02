package application_test

import (
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"
	"github.com/stretchr/testify/require"
)

// suíte de testes unitários da classe Product

//// o product só é válido se possuir status consistente e price >= 0
func TestProduct_IsValid(t *testing.T) {
	// criando o product
	product := application.NewProduct()
	product.Name = "hello"
	product.Price = 10

	// validando o product
	_, err := product.IsValid()
	// deve-se retornar nil
	require.Nil(t, err)

	// validando o product com status inconsistente
	product.Status = "INVALID"
	_, err = product.IsValid()
	// deve-se retornar erro
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	// reajustando o status para válido
	product.Status = application.ENABLED
	_, err = product.IsValid()
	// deve-se retornar nil
	require.Nil(t, err)

	// validando o product com price < 0
	product.Price = -10
	_, err = product.IsValid()
	// deve-se retornar erro
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

//// o product só pode ser ativado se possuir price > 0
func TestProduct_Enable(t *testing.T) {
	// criando o product
	product := application.NewProduct()
	product.Name = "hello"
	product.Price = 10

	// ativando o product com price > 0
	err := product.Enable()
	// deve-se retornar nil
	require.Nil(t, err)

	// ativando o product com price = 0
	product.Price = 0
	err = product.Enable()
	// deve-se retornar erro
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())

	// ativando o product com price < 0
	product.Price = -10
	err = product.Enable()
	// deve-se retornar erro
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

//// o product só pode ser desativado se possuir price = 0
func TestProduct_Disable(t *testing.T) {
	// criando o product
	product := application.NewProduct()
	product.Name = "hello"
	product.Price = 0

	// desativando o product com price = 0
	err := product.Disable()
	// deve-se retornar nil
	require.Nil(t, err)

	// desativando o product com price > 0
	product.Price = 10
	err = product.Disable()
	// deve-se retornar erro
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())

	// desativando o product com price < 0
	product.Price = -10
	err = product.Disable()
	// deve-se retornar erro
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

//// o getter deve retornar corretamente o atributo
func TestProduct_GetID(t *testing.T) {
	// criando o product
	product := application.NewProduct()
	product.Name = "hello"
	product.Price = 10

	// obtendo o valor do atributo retornado
	id := product.GetID()

	// comparando-se com o valor passado
	require.Equal(t, product.ID, id)
}

//// o getter deve retornar corretamente o atributo
func TestProduct_GetName(t *testing.T) {
	// criando o product
	product := application.NewProduct()
	product.Name = "hello"
	product.Price = 10

	// obtendo o valor do atributo retornado
	name := product.GetName()

	// comparando-se com o valor passado
	require.Equal(t, product.Name, name)
}

//// o getter deve retornar corretamente o atributo
func TestProduct_GetStatus(t *testing.T) {
	// criando o product
	product := application.NewProduct()
	product.Name = "hello"
	product.Price = 10

	// obtendo o valor do atributo retornado
	status := product.GetStatus()

	// comparando-se com o valor passado
	require.Equal(t, product.Status, status)
}

//// o getter deve retornar corretamente o atributo
func TestProduct_GetPrice(t *testing.T) {
	// criando o product
	product := application.NewProduct()
	product.Name = "hello"
	product.Price = 10

	// obtendo o valor do atributo retornado
	price := product.GetPrice()

	// comparando-se com o valor passado
	require.Equal(t, product.Price, price)
}
