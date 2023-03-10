-- para criar um módulo go utiliza-se: go mod init nome-do-repositório-ou-url;

-- para atualizar as dependências do projeto utiliza-se: go mod tidy;

-- para gerar os mocks necessários para testar o service utiliza-se: mockgen -destination={nome-da-pasta-de-destino/nome-do-arquivo-de-mocks} -source={nome-do-arquivo-com-as-interfaces} {nome-do-pacote}
Ex: mockgen -destination=application/mocks/application.go -source=application/product.go application

-- para inicializar o cli do cobra utiliza-se, na pasta onde está o go.mod: cobra-cli init 

