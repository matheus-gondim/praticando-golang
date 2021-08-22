package main

import (
	"aplicacao-cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida...")

	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
