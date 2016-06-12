package fortunereader

import (
    "bufio"
    "io"
    "os"
    "bytes"
)

var separator = []byte("\n%\n")

func scanFortunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if atEOF && len(data) == 0 {
        return 0, nil, nil
    }

    if i := bytes.Index(data, separator); i >= 0 {
        return i + 3, data[0:i], nil
    }

    if atEOF {
        return len(data), data, nil
    }

    return 0, nil, nil
}

func ReadFortunes(r io.Reader) []string {
    s := bufio.NewScanner(r)
    s.Split(scanFortunes)

    a := make([]string, 0)
    for s.Scan() {
        currentText := s.Text()
        a = append(a, currentText)
    }

    return a
}

func ReadFortunesFromFile(filename string) []string {
    file, err := os.Open(filename)

    if err != nil {
        panic(err)
    }

    return ReadFortunes(file)
}

