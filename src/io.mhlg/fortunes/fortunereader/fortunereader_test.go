package fortunereader

import (
    "testing"
    "reflect"
)

func TestFortunes(t *testing.T) {
    var tests = []struct {
        input string
        expected []string
    } {
        { "fortune", []string{"fortune"} },
        { "fortune\n", []string{"fortune\n"} },
        { "fortune\n%\n", []string{"fortune\n"} },
        { "fortune1\n%\nfortune2", []string{"fortune1\n", "fortune2"}},
        { "fortune1\n%\nfortune2\n%\n", []string{"fortune1\n", "fortune2\n"}},
    }

    for _, tt := range tests {
        actual := parseFortunes(tt.input)

        if !reflect.DeepEqual(actual, tt.expected) {
            t.Errorf("parseFortune(%v): expected %d, got %d", tt.input, tt.expected, actual)
        }
    }
}
