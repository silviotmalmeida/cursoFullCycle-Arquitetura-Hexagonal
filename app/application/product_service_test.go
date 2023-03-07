package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"
	mock_application "github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

// suíte de testes unitários da classe Product

//// se for executada uma busca por product no bd, os atributos devem ser iguais aos do objeto de origem
func TestProductService_Get(t *testing.T) {

	// criando o controlador do gomock
	ctrl := gomock.NewController(t)
	// o controlador será finalizado ao fim do teste
	defer ctrl.Finish()

	// criando o mock de product
	product := mock_application.NewMockProductInterface(ctrl)
	// criando o mock do mecanismo de persistência
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	// sempre que método Get do mecanismo de interface for chamado, retorna um product
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	// criando o productService a ser testado
	service := application.ProductService{Persistence: persistence}
	// realizando a consulta pelo service
	result, err := service.Get("abc")
	// não pode retornar erro
	require.Nil(t, err)
	// deve retornar um product igual ao mock
	require.Equal(t, product, result)
}

//// se um registro for armazenado no bd, seus atributos devem ser iguais aos do objeto de origem
func TestProductService_Create(t *testing.T) {

	// criando o controlador do gomock
	ctrl := gomock.NewController(t)
	// o controlador será finalizado ao fim do teste
	defer ctrl.Finish()

	// criando o mock de product
	product := mock_application.NewMockProductInterface(ctrl)
	// criando o mock do mecanismo de persistência
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	// sempre que método Save do mecanismo de interface for chamado, retorna um product
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	// criando o productService a ser testado
	service := application.ProductService{Persistence: persistence}
	// realizando a criação pelo service
	result, err := service.Create("Product 1", 10)
	// não pode retornar erro
	require.Nil(t, err)
	// deve retornar um product igual ao mock
	require.Equal(t, product, result)
}

//// se um product for ativado/desativado no bd, seus atributos devem ser iguais aos do objeto de origem
func TestProductService_EnableDisable(t *testing.T) {

	// criando o controlador do gomock
	ctrl := gomock.NewController(t)
	// o controlador será finalizado ao fim do teste
	defer ctrl.Finish()

	// criando o mock de product
	product := mock_application.NewMockProductInterface(ctrl)
	// retorna nil ao ativar
	product.EXPECT().Enable().Return(nil)
	// retorna nil ao desativar
	product.EXPECT().Disable().Return(nil)

	// criando o mock do mecanismo de persistência
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	// sempre que método Save do mecanismo de interface for chamado, retorna um product
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	// criando o productService a ser testado
	service := application.ProductService{Persistence: persistence}
	// realizando a ativação pelo service
	result, err := service.Enable(product)
	// não pode retornar erro
	require.Nil(t, err)
	// deve retornar um product igual ao mock
	require.Equal(t, product, result)

	// realizando a desativação pelo service
	result, err = service.Disable(product)
	// não pode retornar erro
	require.Nil(t, err)
	// deve retornar um product igual ao mock
	require.Equal(t, product, result)
}
