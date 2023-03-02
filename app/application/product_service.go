package application

// definindo a classe productService que implementa a interface de productService
// para manipulação de products sem acessar a classe diretamente
type ProductService struct {
	// definindo os atributos, tipos e restrições
	Persistence ProductPersistenceInterface
}

// definindo o construtor
func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	// preenchendo os atributos mínimos para criação
	return &ProductService{Persistence: persistence}
}

// método para consulta
func (s *ProductService) Get(id string) (ProductInterface, error) {
	// consultando no BD
	product, err := s.Persistence.Get(id)
	// caso existam erros, retorna-os
	if err != nil {
		return nil, err
	}
	// senão retorna o product
	return product, nil
}

// método para criação
func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	// criando o product
	product := NewProduct()
	product.Name = name
	product.Price = price
	// validando o product
	_, err := product.IsValid()
	// caso existam erros, retorna um product vazio e o erro
	if err != nil {
		return &Product{}, err
	}
	// salvando o registro no BD
	result, err := s.Persistence.Save(product)
	// caso existam erros, retorna um product vazio e o erro
	if err != nil {
		return &Product{}, err
	}
	// senão retorna o resultado da query
	return result, nil
}

// método para ativação do product
func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	// ativando o product
	err := product.Enable()
	// caso existam erros, retorna um product vazio e o erro
	if err != nil {
		return &Product{}, err
	}
	// salvando o registro no BD
	result, err := s.Persistence.Save(product)
	// caso existam erros, retorna um product vazio e o erro
	if err != nil {
		return &Product{}, err
	}
	// senão retorna o resultado da query
	return result, nil
}

// método para desativação do product
func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	// desativando o product
	err := product.Disable()
	// caso existam erros, retorna um product vazio e o erro
	if err != nil {
		return &Product{}, err
	}
	// salvando o registro no BD
	result, err := s.Persistence.Save(product)
	// caso existam erros, retorna um product vazio e o erro
	if err != nil {
		return &Product{}, err
	}
	// senão retorna o resultado da query
	return result, nil
}
