package main

import (
	"fmt"
	"os"

	elastic "gopkg.in/olivere/elastic.v3"

	"github.com/urfave/cli"
)

var (
	app = cli.NewApp()

	hostname string
)

func init() {
	app.Name = "elastictool"
	app.Author = "Herbert Fischer"
	app.Usage = "Tools for ElasticSearch"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "hostname, H",
			Value:       "127.0.0.1:9200",
			Usage:       "ElasticSearch `HOST` to connect to",
			Destination: &hostname,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "status",
			Usage: "get cluster status",
			Action: func(c *cli.Context) error {
				println(hostname)
				es, err := elastic.NewClient(elastic.SetURL(hostname))
				if err != nil {
					return err
				}
				health, err := es.ClusterHealth().Do()
				if err != nil || health == nil {
					return err
				}
				println(health.Status)
				return nil
			},
		},
		{
			Name:  "rolling-restart",
			Usage: "execute a cluster rolling restart",
			Action: func(c *cli.Context) error {
				fmt.Println(c)
				return nil
			},
		},
	}
}

func main() {
	app.Run(os.Args)
}
