package main

import toolkit "github.com/50u7h/go-module"

func main() {
	var tools toolkit.Tools

	err := tools.CreateDirIfNotExist("./test-dir")
	if err != nil {
		return
	}
}
