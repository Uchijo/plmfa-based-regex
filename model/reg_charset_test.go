package model

import (
	"fmt"
	"testing"
)

func TestCharRange(t *testing.T) {
	tests := []struct {
		name    string
		input   CharRange
		matcher string
		matches bool
	}{
		{
			name: "[a-z], a",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: true,
			},
			matcher: "a",
			matches: true,
		},
		{
			name: "[a-z], c",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: true,
			},
			matcher: "c",
			matches: true,
		},
		{
			name: "[a-z], z",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: true,
			},
			matcher: "z",
			matches: true,
		},
		{
			name: "[a-z], A",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: true,
			},
			matcher: "A",
			matches: false,
		},
		{
			name: "[a-z], abc",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: true,
			},
			matcher: "abc",
			matches: false,
		},
		{
			name: "[^a-z], a",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: false,
			},
			matcher: "a",
			matches: false,
		},
		{
			name: "[^a-z], c",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: false,
			},
			matcher: "c",
			matches: false,
		},
		{
			name: "[^a-z], z",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: false,
			},
			matcher: "z",
			matches: false,
		},
		{
			name: "[^a-z], A",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: false,
			},
			matcher: "A",
			matches: true,
		},
		{
			name: "[^a-z], abc",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: false,
			},
			matcher: "abc",
			matches: false,
		},
		{
			name: "[^a-z], ABC",
			input: CharRange{
				Start:     'a',
				End:       'z',
				WhiteList: false,
			},
			matcher: "ABC",
			matches: false,
		},
		{
			name: "[0-9], 0",
			input: CharRange{
				Start:     '0',
				End:       '9',
				WhiteList: true,
			},
			matcher: "0",
			matches: true,
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf("CharRange: %s", td.name), func(t *testing.T) {
			if td.input.WholeMatches(td.matcher) != td.matches {
				t.Errorf("match %v against %v expected to be %v, but got %v", td.input, td.matcher, td.matches, td.input.WholeMatches(td.matcher))
			}
		})
	}
}

func TestCharList(t *testing.T) {
	tests := []struct {
		name    string
		input   CharList
		matcher string
		matches bool
	}{
		{
			name:    "[a], a",
			input:   CharList{
				Chars: []rune{'a'},
				WhiteList: true,
			},
			matcher: "a",
			matches: true,
		},
		{
			name:    "[a], b",
			input:   CharList{
				Chars: []rune{'a'},
				WhiteList: true,
			},
			matcher: "b",
			matches: false,
		},
		{
			name:    "[a], aa",
			input:   CharList{
				Chars: []rune{'a'},
				WhiteList: true,
			},
			matcher: "aa",
			matches: false,
		},
		{
			name:    "[a], bb",
			input:   CharList{
				Chars: []rune{'a'},
				WhiteList: true,
			},
			matcher: "bb",
			matches: false,
		},
		{
			name:    "[abc], a",
			input:   CharList{
				Chars: []rune{'a', 'b', 'c'},
				WhiteList: true,
			},
			matcher: "a",
			matches: true,
		},
		{
			name:    "[abc], c",
			input:   CharList{
				Chars: []rune{'a', 'b', 'c'},
				WhiteList: true,
			},
			matcher: "c",
			matches: true,
		},
		{
			name:    "[abc], ab",
			input:   CharList{
				Chars: []rune{'a', 'b', 'c'},
				WhiteList: true,
			},
			matcher: "ab",
			matches: false,
		},
		{
			name:    "[abc], AB",
			input:   CharList{
				Chars: []rune{'a', 'b', 'c'},
				WhiteList: true,
			},
			matcher: "AB",
			matches: false,
		},
		{
			name:    "[abc], d",
			input:   CharList{
				Chars: []rune{'a', 'b', 'c'},
				WhiteList: true,
			},
			matcher: "d",
			matches: false,
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf("CharList: %s", td.name), func(t *testing.T) {
			if td.input.WholeMatches(td.matcher) != td.matches {
				t.Errorf("match %v against %v expected to be %v, but got %v", td.input, td.matcher, td.matches, td.input.WholeMatches(td.matcher))
			}
		})
	}
}
