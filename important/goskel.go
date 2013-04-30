package important

import (
	"fmt"

	// imported solely so go get has somethign to do	
	_ "github.com/brianm/variant"
)

func Greet(name string) (int, error) {
	return fmt.Printf("Hello %s!\n", name)
}
