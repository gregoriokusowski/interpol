package main

import (
	"fmt"

	"os"

	"github.com/gregoriokusowski/interpol"
)

func main() {
	result := interpol.Check(".interpol.yml")
	if len(result.Errors) > 0 {
		for _, e := range result.Errors {
			fmt.Printf("%s - %s\n", e.Locale, e.Message)
		}
		os.Exit(1)
	}
}
