package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
)

const (
	version           = "0.0.1"
	keymasterBasePath = "/opsee.co/keymaster"
)

func getClient(etcdhost []string) *etcd.Client {
	return etcd.NewClient(etcdhost)
}

func get(c *cli.Context) {
	client := getClient([]string{c.GlobalString("url")})
	defer client.Close()

	user := c.Args().First()
	if user == "" {
		fmt.Println("Must provide a username.")
		return
	}

	keyPath := fmt.Sprintf("%s/%s", keymasterBasePath, user)

	resp, err := client.Get(keyPath, false, true)
	if err != nil {
		fmt.Println("Unable to get key for %s", user)
		fmt.Println(err.Error())
	}

	keyNodes := resp.Node.Nodes
	fmt.Println(keyNodes)
}

func put(c *cli.Context) {
	client := getClient([]string{c.GlobalString("url")})
	defer client.Close()
}

func main() {
	app := cli.NewApp()
	app.Name = "keymaster"
	app.Usage = "Retrieve a public key for a user from Etcd2"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Usage: "Etcd2 URL",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "get an SSH public key from etcd2",
			Action: get,
		},
		{
			Name:   "put",
			Usage:  "put an SSH public key into etcd2",
			Action: put,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "identity, i",
					Usage: "Path to public key file",
				},
			},
		},
	}

	app.Run(os.Args)
}
