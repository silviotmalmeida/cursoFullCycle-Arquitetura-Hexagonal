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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

//// se um product for ativado/desativado no bd, seus atributos devem ser iguais aos do objeto de origem
func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
