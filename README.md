# Golang

## Test

    go test ./... executa todos os testes do projeto
    t.Parallel() se colocado dentro das funções de testes elas rodam em paralelo
    go test --cover
    go test --coverprofile <nome_arquivo> comando para gerar um relatório sobre o que falta cobrir nas funções, para ler esse relatório usa-se o comando go tool cover --func=nome_arquivo, mas para mostrar detalhada mente usa-se o comando go tool cover --html=nome_arquivo