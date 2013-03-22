package goskel

import(
	"fmt"
	_ "github.com/brianm/variant"
)

func Greet(name string) {
	fmt.Printf("hello %s\n", name)
}