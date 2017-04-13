package core_test

import (
	"testing"

	"github.com/AlexsJones/warpcore/core"
)

func TestParser(t *testing.T) {

	var tests = []struct {
		input    []string
		expected string
	}{
		{nil, string("Empty file list")},
		{[]string{""}, string("Nothing to parse")},
	}

	for _, tt := range tests {
		_, err := core.ParseFiles(tt.input)
		t.Log(err.Error())
		if err.Error() != tt.expected {
			t.Fail()
		}
	}

}
