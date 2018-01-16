package main

// go test -coverprofile c.out
// go tool cover -html=c.out -o c.html

import (
	"fmt"

	"github.com/devimc/api"
)

func exercise() bool {
	if api.ValidUser() {
		if err := api.WriteLog("exercise"); err == nil {
			return true
		}

		fmt.Println("unable to write log")
		return false
	}

	fmt.Println("invalid user")
	return false
}
