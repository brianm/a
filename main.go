package main

import (
	"fmt"
	"github.com/brianm/asn/asana"
	"os"
	"github.com/codegangsta/cli"
)

var key = os.Getenv("ASANA_KEY")

func main() {

	app := cli.NewApp()
	app.Name = "asn"
	app.Usage = "asn <command>"
	//	app.Action = func(c *cli.Context) {
	//		println("boom! I say!")
	//	}
	app.Commands = []cli.Command{
		{
			Name:      "me",
			ShortName: "m",
			Usage:     "Who am I?",
			Action: me,
		},
		{
			Name:      "tasks",
			ShortName: "t",
			Usage:     "list tasks in first workspace",
			Action: tasks,
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

func me(_ *cli.Context) {
	c, err := asana.NewClient(key)
	if err != nil {
		panic(err)
	}
	
	me := c.Me
	fmt.Printf("Me\n%+v\n\n", me)
}
