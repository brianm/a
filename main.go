package main

import (
	"fmt"
	"github.com/brianm/a/asana"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v1"
	"os"
)

var key = os.Getenv("ASANA_KEY")

func main() {
	app := cli.NewApp()
	app.Name = "asn"
	app.Usage = "asn <command>"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "me",
			Usage:  "Who am I?",
			Action: me,
		},
		{
			Name:   "tasks",
			Usage:  "list tasks in first workspace",
			Action: tasks,
		},
		{
			Name:      "workspace",
			ShortName: "ws",
			Action:    workspace,
		},
		{
			Name:         "finish",
			Usage:        "Finish a task",
			Action:       finish,
			BashComplete: finishCompletion,
		},
	}

	app.Run(os.Args)
}

func client() asana.Client {
	c, err := asana.NewClient(key)
	if err != nil {
		panic(err)
	}
	return c
}

func workspace(*cli.Context) {
	c := client()
	bs, err := yaml.Marshal(&c.Me.Workspaces)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(bs))
}

func tasks(*cli.Context) {
	c := client()

	for _, w := range c.Me.Workspaces {
		tasks, err := c.Tasks(w)
		if err != nil {
			panic(err)
		}

		bs, err := yaml.Marshal(&tasks)
		if err != nil {
			panic(err)
		}
		fmt.Print(string(bs))
	}
}

func finish(*cli.Context) {
	println("finished!")
}

func finishCompletion(ctx *cli.Context) {
	if len(ctx.Args()) > 0 {
		return
	}

	c, err := asana.NewClient(key)
	if err != nil {
		panic(err)
	}

	tasks, err := c.Tasks(c.Me.Workspaces[0])
	if err != nil {
		panic(err)
	}

	for _, t := range tasks {
		fmt.Printf("%d\n", t.Id)
	}
}

func me(_ *cli.Context) {
	c, err := asana.NewClient(key)
	if err != nil {
		panic(err)
	}

	me := c.Me
	bs, err := yaml.Marshal(&me)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(bs))
}

func init() {
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} command

ENVIRONMENT:
   ASANA_KEY Environment variable which must contain 
             Asana API Key

VERSION:
   {{.Version}}

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}
`
}
