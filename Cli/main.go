package main

import (
	"cli/app"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {

	aplication := app.Generate()

	// if err := aplication.Run(os.Args); err != nil {
	// 	log.Fatal(err)
	// }

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	aplication.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search Ips on Ethernet",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "server",
			Usage:  "Search Server Names on Ethernet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	err := aplication.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func searchServers(context *cli.Context) {
	host := context.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for i, serv := range servers {
		fmt.Println(i+1, "-", serv.Host)
	}
}

func searchIPs(context *cli.Context) {
	host := context.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for i, ip := range ips {
		fmt.Println(i+1, "-", ip)
	}
}

// go run main.go --host amazon.com.br
