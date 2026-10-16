package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len(cleanInput(%d)) != len(%d)", len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) == %q, want %q", c.input, actual, c.expected)
				t.Fail()
			}
		}
	}
}
