# Use a imagem base do Go para construir a aplicação
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copie o código-fonte para o diretório de trabalho do contêiner
COPY . .

# Baixe as dependências do projeto
RUN go get -d -v ./...

# Compile o código Go
RUN go build -o main

# Defina o comando padrão para executar o binário
CMD ["./main"]
