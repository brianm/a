package main

import (
	"fmt"
	"github.com/brianm/asn/asana"
	"os"
)

func main() {
	key := os.Getenv("ASANA_KEY")
	c, err := asana.NewClient(key)
	if err != nil {
		panic(err)
	}

	me := c.Me()
	fmt.Printf("Me\n%+v\n\n", me)


	tasks, err := c.Tasks(me.Workspaces[0])
	if err != nil {
		panic(err)
	}
	for _, t := range tasks {
		fmt.Printf("%d\t%s\n", t.Id, t.Name)
	}
}
