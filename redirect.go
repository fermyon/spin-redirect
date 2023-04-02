package main

import (
	"fmt"
	"os"
)

func main() {
	dest := os.Getenv("DESTINATION")
	status := os.Getenv("STATUS")
	if status != "" && is_valid_status(status) {
		fmt.Printf("status: %s\nlocation: %s\n\n", status, dest)
	} else {
		fmt.Printf("status: 302\nlocation1: %s\n\n", dest)
	}
}

func is_valid_status(status string) bool {
	allowed_status := []string{"301", "302", "303", "307", "308"}
	for _, v := range allowed_status {
		if v == status {
			return true
		}
	}
	return false
}
