package main

import (
	"github.com/mhaligowski/fortunes-server"
	"os"
)

func main() {
	f := os.Args[1]
	content := fortunes.ReadFortunesFromFile(f)
	fortunes.StartFortunesServer(content)
}
