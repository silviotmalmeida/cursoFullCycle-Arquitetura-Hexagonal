package cli

import (
	"fmt"

	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"
)

// função para iniciar o adapter de command line interface
func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	// inicializando a saída do cli vazia
	var result = ""

	// iterando sobre as opções de action
	switch action {
	// criação
	case "create":
		// cria o product pelo service
		product, err := service.Create(productName, price)
		// em caso de erro, retorna-o
		if err != nil {
			return result, err
		}
		// definindo a saída do cli
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	// ativação
	case "enable":
		// obtém o product pelo service através do id
		product, err := service.Get(productId)
		// em caso de erro, retorna-o
		if err != nil {
			return result, err
		}
		// ativa o product pelo service
		res, err := service.Enable(product)
		// em caso de erro, retorna-o
		if err != nil {
			return result, err
		}
		// definindo a saída do cli
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	// desativação
	case "disable":
		// obtém o product pelo service através do id
		product, err := service.Get(productId)
		// em caso de erro, retorna-o
		if err != nil {
			return result, err
		}
		// desativa o product pelo service
		res, err := service.Disable(product)
		// em caso de erro, retorna-o
		if err != nil {
			return result, err
		}
		// definindo a saída do cli
		result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
	// padrão
	default:
		// obtém o product pelo service através do id
		res, err := service.Get(productId)
		// em caso de erro, retorna-o
		if err != nil {
			return result, err
		}
		// definindo a saída do cli
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
