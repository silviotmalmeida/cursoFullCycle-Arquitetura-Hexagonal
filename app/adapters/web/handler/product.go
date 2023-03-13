package handler

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/adapters/dto"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"
)

// função responsável por definir as rotas
func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {

	// rota do resource enable
	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")

	// rota do resource disable
	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("GET", "OPTIONS")

	// rota do resource get
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	// rota do resource create
	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
}

// função que retorna os dados de um product a partir do id
func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// definindo o retorno como json
		w.Header().Set("Content-Type", "application/json")
		// obtendo todas as variáveis da requisição
		vars := mux.Vars(r)
		// obtendo o id
		id := vars["id"]
		// utilizando o service para obter os dados do product
		product, err := service.Get(id)
		// em caso de erro, retorna 404
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// convertendo o retorno para json
		err = json.NewEncoder(w).Encode(product)
		// em caso de erro, retorna 500
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

// função cria um novo product
func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

// função que ativa um product a partir do id
func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

// função que desativa um product a partir do id
func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
