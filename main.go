package main

import (
	"fmt"
	"github.com/brianm/asn/asana"
	"github.com/codegangsta/cli"
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
			Name: "finish",
			Usage: "Finish a task",
			Action: finish,
			BashComplete: finishCompletion,
		},
	}

	app.Run(os.Args)
}

func tasks(_ *cli.Context) {
	c, err := asana.NewClient(key)
	if err != nil {
		panic(err)
	}

	tasks, err := c.Tasks(c.Me.Workspaces[0])
	if err != nil {
		panic(err)
	}
	for _, t := range tasks {
		fmt.Printf("%d\t%s\n", t.Id, t.Name)
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
	fmt.Printf("Me\n%+v\n\n", me)
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
}
`
}
