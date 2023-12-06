# Usa a imagem base oficial do Golang 1.21
FROM golang:latest

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código-fonte e go.mod para o container
COPY . .

# Faz o build da aplicação
RUN go build -o easytrade-backend ./cmd

# Expõe a porta que a aplicação irá escutar
EXPOSE 8000

# Comando para executar a aplicação
CMD [ "./easytrade-backend" ]