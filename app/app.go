// app.go
package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Variáveis globais para permitir sobrescrever as funções nos testes
var lookupIP = net.LookupIP
var lookupNS = net.LookupNS

// Build vai retornar a aplicação de linha de comando pronta para ser executada
func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servidores",
			Usage:  "Buscar o nome do servidores na internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchServers(c *cli.Context) {
	host := c.String("host")
	servidores, erro := lookupNS(host) // Usando a variável global lookupNS
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := lookupIP(host) // Usando a variável global lookupIP
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
