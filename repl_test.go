package main

import "testing"

type testCase struct {
	input, expected string
}

func TestCleanText(t *testing.T) {

	var testCases = []testCase{
		{" hallo", "hallo"},
		{"ballo ", "ballo"},
		{"FOO", "foo"},
		{"     BAR  ", "bar"},
	}

	for _, test := range testCases {
		if output := CleanText(test.input); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", test.expected, test.input)
		}
	}

}
