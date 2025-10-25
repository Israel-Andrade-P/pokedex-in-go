package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello     world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "PIKACHU  Lucario CharmanDER",
			expected: []string{"pikachu", "lucario", "charmander"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test failed\nExpected Length: %d\nActual Length: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test failed\nExpected Word: %s\nActual Word: %s", expectedWord, word)
			}
		}
	}
}

func TestCleanParameter(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"cumbuco", "BEACH"},
			expected: "cumbuco-beach",
		},
		{
			input:    []string{"SAItama"},
			expected: "saitama",
		},
		{
			input:    []string{"Canoa", "Quebrada", "Beach"},
			expected: "canoa-quebrada-beach",
		},
	}

	for _, c := range cases {
		actual := cleanParameter(c.input)
		if actual != c.expected {
			t.Errorf("Test failed\nExpected string: %s\nActual string: %s", c.expected, actual)
		}
	}
}
