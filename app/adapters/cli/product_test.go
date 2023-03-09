package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/adapters/cli"
	mock_application "github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

// suíte de testes unitários da command line interface

func TestRun(t *testing.T) {

	// criando o controlador do gomock
	ctrl := gomock.NewController(t)
	// o controlador será finalizado ao fim do teste
	defer ctrl.Finish()

	// inicializando os atributos do product
	productId := "abc"
	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"

	// criando o mock de product, para retornar os atributos definidos anteriormente
	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	// criando o mock de productService, para retornar os dados definidos no mock de produvt
	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// testando o create
	//// montando a saida esperada
	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus)
	//// executando a cli
	result, err := cli.Run(service, "create", "", productName, productPrice)
	//// não deve retornar erro
	require.Nil(t, err)
	//// deve retornar a saída esperada
	require.Equal(t, resultExpected, result)

	// testando o enable
	//// montando a saida esperada
	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)
	//// executando a cli
	result, err = cli.Run(service, "enable", productId, "", 0)
	//// não deve retornar erro
	require.Nil(t, err)
	//// deve retornar a saída esperada
	require.Equal(t, resultExpected, result)

	// testando o disable
	//// montando a saida esperada
	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)
	//// executando a cli
	result, err = cli.Run(service, "disable", productId, "", 0)
	//// não deve retornar erro
	require.Nil(t, err)
	//// deve retornar a saída esperada
	require.Equal(t, resultExpected, result)

	// testando a opção padrão
	//// montando a saida esperada
	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)
	//// executando a cli
	result, err = cli.Run(service, "", productId, "", 0)
	//// não deve retornar erro
	require.Nil(t, err)
	//// deve retornar a saída esperada
	require.Equal(t, resultExpected, result)

}
