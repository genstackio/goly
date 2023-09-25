package main

import (
	"fmt"
	"github.com/genstackio/goly"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Syntax: goly <vendorToken> <env> <action>")
		os.Exit(1)
	}
	vendorToken := os.Args[1]
	env := os.Args[2]
	action := os.Args[3]
	switch action {
	case "create-request":
		c := goly.Client{}
		c.Init(vendorToken, env)
		var data url.Values
		r, err := c.CreateRequest(data)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			os.Exit(4)
		}
		fmt.Println(r)
	default:
		fmt.Println("Error: unrecognized action '" + action + "'")
		os.Exit(2)
	}
}
