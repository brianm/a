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
	fmt.Printf("Me\n%+v\n\n", me)

	me, err = c.User("me")
	if err != nil {
		panic(err)
	}
	fmt.Printf("'me'\n%+v\n\n", me)

	me, err = c.User(me.Id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number\n%+v\n", me)
}
