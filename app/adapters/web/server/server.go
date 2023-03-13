package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/adapters/web/handler"
	"github.com/silviotmalmeida/cursoFullCycle-Arquitetura-Hexagonal/application"
)

// definindo a classe adapter webserver http
type Webserver struct {
	Service application.ProductServiceInterface
}

// definindo o construtor
func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

// função que define as configurações do webserver http
func (w Webserver) Serve() {

	// criando o roteador
	r := mux.NewRouter()

	// criando o logger
	n := negroni.New(
		negroni.NewLogger(),
	)

	// criando o handler para tratar a requisição
	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	// configurações do webserver
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	// subindo o servidor
	err := server.ListenAndServe()
	// em caso de erro, retorna-o
	if err != nil {
		log.Fatal(err)
	}
}
