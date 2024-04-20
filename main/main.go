package main

import (
	"fmt"

	"github.com/spynacz/remitlyiamvalidator/iamvalid"
)

func main() {
	res, err := iamvalid.IsValid("example.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
