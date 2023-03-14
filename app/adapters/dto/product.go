package dto

import "github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"

// definindo a classe de DTO para transferência de dados entre a requisição e o product
type Product struct {
	// realizando a correlação entre os atributos do json da requisição e do product
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

// definindo o construtor
func NewProduct() *Product {
	return &Product{}
}

// função para relacionar/validar os dados para um product
func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	// relacionando os dados
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	// validando os dados
	_, err := product.IsValid()
	// em caso de erro, retorna-o
	if err != nil {
		return &application.Product{}, err
	}
	//retorna o product
	return product, nil
}
