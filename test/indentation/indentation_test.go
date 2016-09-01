package main

import (
	"fmt"
	"testing"
)

var cases = map[string]string{
	`if a:
    a = z
b = c
`: "[{condition a [{assignment a z}]} {assignment b c}]",
}

func TestIndentation(t *testing.T) {
	for tc, exp := range cases {
		got, err := Parse("", []byte(tc), Debug(true), Memoize(true))
		if err != nil {
			got = err.Error()
		}
		got = fmt.Sprintf("%v", got)
		if got != exp {
			t.Errorf("%q: want %v, got %v", tc, exp, got)
		}
	}
}
