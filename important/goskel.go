package important

import (
	"fmt"
	_ "github.com/brianm/variant" // imported solely so go get has somethign to do
)

func Greet(name string) (int, error) {
	return fmt.Printf("Hello %s!\n", name)
}
