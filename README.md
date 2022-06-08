# Golang

## Test

    go test ./... executa todos os testes do projeto
    t.Parallel() se colocado dentro das funções de testes elas rodam em paralelo
    go test --cover
    go test --coverprofile <nome_arquivo> comando para gerar um relatório sobre o que falta cobrir nas funções, 
    go tool cover --func=nome_arquivo comando para ler o relatório de cover
    go tool cover --html=nome_arquivo comando para mostrar detalhadamente em html
