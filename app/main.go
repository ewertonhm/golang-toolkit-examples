package main

import (
	"fmt"

	toolkit "github.com/ewertonhm/golang-toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomString(10)
	fmt.Println(s)
}
