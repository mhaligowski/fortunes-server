package fortunes

import (
    "testing"
    "reflect"
    "strings"
)

func TestFortunes(t *testing.T) {
    var tests = []struct {
        input string
        expected []string
    } {
        { "", []string{} },
        { "fortune", []string{"fortune"} },
        { "fortune\n", []string{"fortune\n"} },
        { "fortune\n%\n", []string{"fortune"} },
        { "fortune1\n%\nfortune2", []string{"fortune1", "fortune2"}},
        { "fortune1\n%\nfortune2\n%\n", []string{"fortune1", "fortune2"}},
    }

    for _, tt := range tests {
        s := strings.NewReader(tt.input)
        actual := ReadFortunes(s)

        if !reflect.DeepEqual(actual, tt.expected) {
            t.Errorf("parseFortune(%v): expected %d, got %d", tt.input, tt.expected, actual)
        }
    }
}
