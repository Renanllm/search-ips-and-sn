package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CMD app"
	app.Description = "Search IPs and Server names on the internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br", // default value
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs on the internet",
			Flags:  flags,
			Action: searchIPs,
		},
		{
			Name:   "server",
			Usage:  "Search server names on the internet",
			Flags:  flags,
			Action: searchServerName,
		},
	}
	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServerName(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
