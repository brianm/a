package main

import (
	"fmt"
	"github.com/brianm/asn/asana"
	"os"
)

func main() {
	key := os.Getenv("ASANA_KEY")
	c := asana.NewClient(key)
	me, err := c.Me()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", me)
}
