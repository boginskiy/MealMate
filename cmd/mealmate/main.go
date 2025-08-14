package main

import (
	"fmt"
	"mealmate/internal/server"
	"os"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
}
