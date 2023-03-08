package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"
)

// definindo a classe adapter productDb que utiliza o módulo de sql nativo do golang
type ProductDb struct {
	db *sql.DB
}

// definindo o construtor
func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

// implementando o método Get da interface productReader, utilizada na productPersistenceInterface
func (p *ProductDb) Get(id string) (application.ProductInterface, error) {

	// criando o product
	var product application.Product

	// preparando o sql
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	// em caso de erro, retorna-o
	if err != nil {
		return nil, err
	}

	// executando a query e relacionando o resultado aos atributos do product criado anteriormente
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	// em caso de erro, retorna-o
	if err != nil {
		return nil, err
	}

	// retorna o product populado
	return &product, nil
}

// implementando o método Save da interface productWriter, utilizada na productPersistenceInterface
func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {

	// inicializando a variável para verificação de existência de registros
	var rows int

	// executando a query para verificação de existência de registro e populando a variável de quantitativo de registros
	p.db.QueryRow("Select id from products where id=?", product.GetID()).Scan(&rows)

	// se não existirem registros
	if rows == 0 {
		// realiza a criação
		_, err := p.create(product)
		// em caso de erro, retorna-o
		if err != nil {
			return nil, err
		}
		// se existir registro
	} else {
		// realiza a atualização
		_, err := p.update(product)
		// em caso de erro, retorna-o
		if err != nil {
			return nil, err
		}
	}
	// retorna o product
	return product, nil
}

// implementando o método de criação
func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {

	// preparando o sql
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values(?,?,?,?)`)
	// em caso de erro, retorna-o
	if err != nil {
		return nil, err
	}

	// executando a query
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	// em caso de erro, retorna-o
	if err != nil {
		return nil, err
	}

	// encerrando a transação
	err = stmt.Close()
	// em caso de erro, retorna-o
	if err != nil {
		return nil, err
	}

	// retorna o product
	return product, nil
}

// implementando o método de atualização
func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {

	// executando a query
	_, err := p.db.Exec("update products set name = ?, price=?, status=? where id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	// em caso de erro, retorna-o
	if err != nil {
		return nil, err
	}

	// retorna o product
	return product, nil
}
