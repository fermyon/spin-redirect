package main

import (
	"fmt"
	"os"
)

func main() {
	dest := os.Getenv("DESTINATION")
	status := os.Getenv("STATUS")
	if status != "" {
		fmt.Printf("status: %s\nlocation: %s\n\n", status, dest)
	} else {
		fmt.Printf("status: 302\nlocation: %s\n\n", dest)
	}
}
