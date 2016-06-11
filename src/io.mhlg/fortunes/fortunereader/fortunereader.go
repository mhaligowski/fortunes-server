package fortunereader

import (
    "strings"
    "io/ioutil"
)

func parseFortunes(input string) []string {
    splitValues := strings.Split(input, "%\n")
    if splitValues[len(splitValues)-1] == "" {
        splitValues = splitValues[:len(splitValues)-1]
    }

    result := make([]string, len(splitValues))
    for i, v := range splitValues {
        result[i] = strings.TrimPrefix(v, "\n")
    }

    return splitValues
}

func ReadFortunes(filename string) []string {
    file, _ := ioutil.ReadFile(filename)

    return parseFortunes(string(file))
}


