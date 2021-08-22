package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicação de lina de comando pronta para ser executada
func Gerar() (app *cli.App) {
	app = cli.NewApp()

	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca ip de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "server",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServer,
		},
	}

	return
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		log.Println(ip)
	}

}

func buscarServer(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		log.Println(server.Host)
	}

}
