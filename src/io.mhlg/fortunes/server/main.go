package main

import (
    "os"
    "fmt"
    "net/http"
    "io.mhlg/fortunes/fortunereader"
    "math/rand"
)

type handler func(http.ResponseWriter, *http.Request)

func getHandler(filename string) handler {
    fortunes := fortunereader.ReadFortunes(filename)

    return func(w http.ResponseWriter, r *http.Request) {
        fortuneNumber := rand.Intn(len(fortunes))
        fmt.Fprintf(w, fortunes[fortuneNumber])
    }
}

func main() {
    fortunesFile := os.Args[1] 
    http.HandleFunc("/", getHandler(fortunesFile))
    http.ListenAndServe(":8080", nil)
}

