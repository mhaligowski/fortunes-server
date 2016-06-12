package fortunes

import (
    "net/http"
    "math/rand"
    "fmt"
)

type randomStringHandler struct {
    strings []string
}

func (h randomStringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fortuneNumber := rand.Intn(len(h.strings))
    fmt.Fprintf(w, h.strings[fortuneNumber])
}

func StartFortunesServer(fortunes []string) {
    h := randomStringHandler{ strings: fortunes }
    http.Handle("/", h)
    http.ListenAndServe(":8080", nil)
}

