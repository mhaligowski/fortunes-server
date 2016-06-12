package main

import (
    "github.com/mhaligowski/fortunes"
    "os"
)

func main() {
    f := os.Args[1]
    content := fortunes.ReadFortunesFromFile(f)
    fortunes.StartFortunesServer(content)
}
