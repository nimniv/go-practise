package main

import (
	"fmt"
	// Add new module requirements and sums. `go mod tidy`
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
