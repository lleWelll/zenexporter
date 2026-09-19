package main

import (
	"fmt"
	"os"

	"zenexport/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
