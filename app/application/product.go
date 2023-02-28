package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// definindo as premissas a serem consideradas pelo goValidator
func init() {
	// todos os atributos devem ser preenchidos
	govalidator.SetFieldsRequiredByDefault(true)
}

// criando a interface de product
type ProductInterface interface {

	// definindo os métodos a serem implementados
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

// definindo as constantes utilizadas
const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

// definindo a classe product que implementa a interface de product
type Product struct {

	// definindo os atributos, tipos e restrições
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

// definindo o construtor
func NewProduct() *Product {
	// preenchendo os atributos mínimos para criação
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

// implementando o método da inteface IsValid()
func (p *Product) IsValid() (bool, error) {

	// se o status estiver vazio, considera DISABLED
	if p.Status == "" {
		p.Status = DISABLED
	}

	// se o status estiver com um valor inconsistente, retorna erro
	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	// se o preço for menor que zero, retorna erro
	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}

	// validação geral com o goValidator
	// as regras são definidas na função init() e nas tags dos atributos
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

// implementando o método da inteface Enable()
func (p *Product) Enable() error {
	// se o price > 0, deve ser possível ativar o product
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	// senão, retorna erro
	return errors.New("the price must be greater than zero to enable the product")
}

// implementando o método da inteface Disable()
func (p *Product) Disable() error {
	// se o price = 0, deve ser possível desativar o product
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	// senão, retorna erro
	return errors.New("the price must be zero in order to have the product disabled")
}

// implementando o método da inteface GetID()
func (p *Product) GetID() string {
	return p.ID
}

// implementando o método da inteface GetName()
func (p *Product) GetName() string {
	return p.Name
}

// implementando o método da inteface GetStatus()
func (p *Product) GetStatus() string {
	return p.Status
}

// implementando o método da inteface GetPrice()
func (p *Product) GetPrice() float64 {
	return p.Price
}
