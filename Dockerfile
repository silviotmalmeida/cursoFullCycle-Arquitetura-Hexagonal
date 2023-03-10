# definindo a imagem base
FROM golang:1.16

# definindo a pasta de trabalho a ser criada e focada no acesso
WORKDIR /go/src

#variáveis de ambiente
ENV PATH="/go/bin:${PATH}"

# comandos necessários
RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest

RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

# comando para manter o container funcionando
CMD ["tail", "-f", "/dev/null"]