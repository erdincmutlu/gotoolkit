package main

import "github.com/erdincmutlu/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExists("./test-dir")
}
