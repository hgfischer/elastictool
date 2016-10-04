package main

import (
	"fmt"
	"os"

	elastic "gopkg.in/olivere/elastic.v3"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var (
	app = cli.NewApp()

	hostname string
	sniff    bool
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
		cli.BoolFlag{
			Name:        "sniff",
			Usage:       "If true, sniffs all nodes on the ElasticSearch node",
			Destination: &sniff,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "status",
			Usage: "get cluster status",
			Action: func(c *cli.Context) error {
				es, err := elastic.NewClient(
					elastic.SetURL(hostname),
					elastic.SetSniff(sniff),
				)
				if err != nil {
					return err
				}
				health, err := es.ClusterHealth().Do()
				if err != nil || health == nil {
					return err
				}
				switch health.Status {
				case "red":
					color.Red(health.Status)
				case "yellow":
					color.Yellow(health.Status)
				case "green":
					color.Green(health.Status)
				}
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
