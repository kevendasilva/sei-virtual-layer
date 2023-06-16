# Use a imagem base do Go para construir a aplicação
FROM arm64v8/golang:1.20-alpine AS builder

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copie o código-fonte para o diretório de trabalho do contêiner
COPY . .

# Baixe as dependências do projeto
RUN  get -d -v ./...

# Compile o código Go
RUN go build -o main

# Defina o comando padrão para executar o binário
CMD ["./main"]
