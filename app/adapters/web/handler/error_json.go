package handler

import "encoding/json"

// função para converter string em json e retornar possíveis erros
func jsonError(msg string) []byte {
	// criando o objeto com a mensagem
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	// tenta converter para json
	r, err := json.Marshal(error)
	// em caso de erro, retorna-o como array de bytes
	if err != nil {
		return []byte(err.Error())
	}
	// retorna a mensagem como array de bytes
	return r
}
