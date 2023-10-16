package main

import (
	"fmt"
	toolkit "github.com/50u7h/go-module"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(100)
	fmt.Println("Random string:", s)
}
