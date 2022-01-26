# BUild
FROM golang:1.16.3-alpine

# Definino o direitório de trabalho, a partir de agora todos os comandos serão executados a partir desse path
WORKDIR $GOPATH/src/ntopus2



# Copiando os fontes do projeto para a pasta cmd do GOPATH
COPY . .

# Baixando todas as dependencias do go, compilando e instalando o projeto
RUN go install


# Abrindo a porta 5000 padrão  para conseguir usar o index do front para teste
EXPOSE 5000


CMD ["ntopus2"]

