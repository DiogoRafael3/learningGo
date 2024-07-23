package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return a cmd application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"                       //name of the app
	app.Usage = "IP and Server name searches from the internet" //usage of the app

	flags := []cli.Flag{ //these are the flags that will be used in the command line, such as --host in this case
		cli.StringFlag{
			Name:  "host",
			Value: "default.com", //this is the default value if the flag is not provided
		},
	}

	app.Commands = []cli.Command{ //this creates a slice of commands, because we need more than a static number of commands
		{
			Name:   "ip",                               //name of the command
			Usage:  "Gets IP adress from the internet", //usage nof the command
			Flags:  flags,                              //this is usually done like above, however since the flgs will be the same between the two commands, we will instead declare them as so
			Action: getIps,                             //func can be declared here as well
		},
		{
			Name:   "server",                              //name of the command
			Usage:  "Gets server names from the internet", //usage nof the command
			Flags:  flags,
			Action: getServers, //func can be declared here as well
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host") //this gets the command flag "host"

	//get IPs from net
	ips, error := net.LookupIP(host) //ips is a slice because one adress could have multiple different ips, an error might also be output
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips { //for each
		fmt.Println(ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")

	//gets servers from net
	servers, error := net.LookupNS(host) //NS stands for Named Server

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
