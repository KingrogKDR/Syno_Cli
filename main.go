package main

import (
	"fmt"
	"os"
	"syno/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error in main: %v", err)
		os.Exit(1)
	}
}
