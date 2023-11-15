package eval

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/uchijo/plmfa-based-regex/model"
	"github.com/uchijo/plmfa-based-regex/parser"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

func TestSearch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		regex         string
		output        bool
		shouldFail    bool
		useEpsilonSem bool
	}{
		{
			name:   "test branch",
			input:  "a",
			regex:  "a|b|c",
			output: true,
		},
		{
			name:   "test lookahead and back-reference 1",
			input:  "aaaaabaaaaa",
			regex:  "(.*)b(?=aaaaa)\\1",
			output: true,
		},
		{
			name:   "test lookahead and back-reference 2",
			input:  "aacaa",
			regex:  "(?=(a*b*))aac\\1",
			output: true,
		},
		{
			name:          "unassigned ref with epsilon-semantics",
			input:         "aaa",
			regex:         "\\1aaa",
			output:        true,
			useEpsilonSem: true,
		},
		{
			name:       "unassigned ref without epsilon-semantics",
			input:      "aaa",
			regex:      "\\1aaa",
			output:     false,
			shouldFail: true,
		},
		{
			name:   "kleene with empty string",
			input:  "",
			regex:  "a*",
			output: true,
		},
		{
			name:   "kleene with single string",
			input:  "a",
			regex:  "a*",
			output: true,
		},
		{
			name:   "kleene with multiple string",
			input:  "aaaa",
			regex:  "a*",
			output: true,
		},
		{
			name:   "kleene with iregal string",
			input:  "b",
			regex:  "a*",
			output: false,
		},
		{
			name:   "{n} form quantifier success",
			input:  "aaa",
			regex:  "a{3}",
			output: true,
		},
		{
			name:   "{n} form quantifier too many",
			input:  "aaaaa",
			regex:  "a{3}",
			output: false,
		},
		{
			name:   "{n} form quantifier less",
			input:  "aa",
			regex:  "a{3}",
			output: false,
		},
		{
			name:   "{n,} form quantifier exact",
			input:  "aaa",
			regex:  "a{3,}",
			output: true,
		},
		{
			name:   "{n,} form quantifier many",
			input:  "aaaa",
			regex:  "a{3,}",
			output: true,
		},
		{
			name:   "{n,} form quantifier many 2",
			input:  "aaaaaaa",
			regex:  "a{3,}",
			output: true,
		},
		{
			name:   "{n,} form quantifier many 2",
			input:  "aaaaaaa",
			regex:  "a{3,}",
			output: true,
		},
		{
			name:   "{n,} form quantifier less",
			input:  "aa",
			regex:  "a{3,}",
			output: false,
		},
		{
			name:   "{n,m} form quantifier less",
			input:  "a",
			regex:  "a{2,4}",
			output: false,
		},
		{
			name:   "{n,m} form quantifier in range 1",
			input:  "aa",
			regex:  "a{2,4}",
			output: true,
		},
		{
			name:   "{n,m} form quantifier in range 2",
			input:  "aaa",
			regex:  "a{2,4}",
			output: true,
		},
		{
			name:   "{n,m} form quantifier in range 3",
			input:  "aaaa",
			regex:  "a{2,4}",
			output: true,
		},
		{
			name:   "{n,m} form quantifier too many",
			input:  "aaaaa",
			regex:  "a{2,4}",
			output: false,
		},
		{
			name:   "non-capture group success",
			input:  "abb",
			regex:  "(?:a)(b)\\1",
			output: true,
		},
		{
			name:   "non-capture group fail",
			input:  "aba",
			regex:  "(?:a)(b)\\1",
			output: false,
		},
		{
			name:   "ref capture 1",
			input:  "aaa",
			regex:  "(a)(\\1)\\2",
			output: true,
		},
		{
			name:   "ref capture 2",
			input:  "abaabaaba",
			regex:  "(aba)(\\1)\\2",
			output: true,
		},
		{
			name:   "check nested ref index",
			input:  "aabbbb",
			regex:  "(aa(bb))\\2",
			output: true,
		},
		{
			name:   "条件によってはキャプチャされてないものに当たったとき",
			input:  "aaaa",
			regex:  "(a|(aa))\\2",
			output: true,
		},
		{
			name:   "char group 1",
			input:  "aaaa",
			regex:  "[a]{4}",
			output: true,
		},
		{
			name:   "char group 2",
			input:  "aaaa",
			regex:  "[a]*",
			output: true,
		},
		{
			name:   "char group 3",
			input:  "aaaa",
			regex:  "[^b]*",
			output: true,
		},
		{
			name:   "char group 4",
			input:  "aaaa",
			regex:  "[^a]*",
			output: false,
		},
		{
			name:   "char group 5",
			input:  "abcdefg",
			regex:  "[a-z]*",
			output: true,
		},
		{
			name:   "char group 6",
			input:  "abcdefg",
			regex:  "[a-zABCD]*",
			output: true,
		},
		{
			name:   "char group 7",
			input:  "abcdefgABCD",
			regex:  "[a-zA-Z]*",
			output: true,
		},
		{
			name:   "char group 8",
			input:  "abcdefgABC",
			regex:  "[a-zABCD]*",
			output: true,
		},
		{
			name:   "char group 9",
			input:  "ABCDF",
			regex:  "[^a-z]*",
			output: true,
		},
		{
			name:   "char group 10",
			input:  "ABcDF",
			regex:  "[^a-z]*",
			output: false,
		},
		{
			name:   "char group 11",
			input:  "DEF",
			regex:  "[^a-zABC]*",
			output: true,
		},
		{
			name:   "char group 12",
			input:  "a",
			regex:  "[^a-zABC]*",
			output: false,
		},
		{
			name:   "char group 13",
			input:  "aABcdef",
			regex:  "[^a-zABC]*",
			output: false,
		},
		{
			name:   "\\d positive",
			input:  "1",
			regex:  "\\d",
			output: true,
		},
		{
			name:   "\\d negative",
			input:  "a",
			regex:  "\\d",
			output: false,
		},
		{
			name:   "\\D positive",
			input:  "a",
			regex:  "\\D",
			output: true,
		},
		{
			name:   "\\D negative",
			input:  "1",
			regex:  "\\D",
			output: false,
		},
		{
			name:   "\\w positive 1",
			input:  "1_2_3",
			regex:  "\\w*",
			output: true,
		},
		{
			name:   "\\w positive 2",
			input:  "1_2_3abcd",
			regex:  "\\w*",
			output: true,
		},
		{
			name:   "\\W positive 1",
			input:  ",.",
			regex:  "\\W*",
			output: true,
		},
	}
	for _, td := range tests {
		td := td
		t.Run(fmt.Sprintf("Search: %s", td.name), func(t *testing.T) {
			// 異常系検知の確認
			defer func() {
				err := recover()
				if err != nil && !td.shouldFail {
					t.Errorf("should not get into panic: %v", err)
				}
			}()

			is := antlr.NewInputStream(td.regex)
			lexer := gen.NewPCRELexer(is)
			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			p := gen.NewPCREParser(stream)
			p.BuildParseTrees = true
			tree := p.Pcre()
			regex := tree.Accept(&parser.RegexBuilder{}).(model.RegExp)

			states, start, _ := model.CreateCompleteStates(regex)
			input := InputBuffer{
				Input: td.input,
			}

			matched := Search(states, input, start, td.useEpsilonSem, false)
			if matched != td.output {
				t.Errorf("expected result is %v, but got %v", td.output, matched)
			}
		})
	}
}

func TestSearchFromAst(t *testing.T) {
	tests := []struct {
		name    string
		input   model.RegExp
		matcher string
		matches bool
	}{
		{
			name: "[a-z]*, abc",
			input: model.RegStar{
				Content: model.RegCharSet{
					Content: model.CharRange{
						Start: 'a',
						End:   'z',
						WhiteList: true,
					},
				},
			},
			matcher: "abc",
			matches: true,
		},
		{
			name: "[a-z]*, abcaalskfjalskdfjslkfdja",
			input: model.RegStar{
				Content: model.RegCharSet{
					Content: model.CharRange{
						Start: 'a',
						End:   'z',
						WhiteList: true,
					},
				},
			},
			matcher: "abcaalskfjalskdfjslkfdja",
			matches: true,
		},
		{
			name: "[a-z]*, AAAA",
			input: model.RegStar{
				Content: model.RegCharSet{
					Content: model.CharRange{
						Start: 'a',
						End:   'z',
						WhiteList: true,
					},
				},
			},
			matcher: "AAAA",
			matches: false,
		},
		{
			name: "[^a-z]*, AAAA",
			input: model.RegStar{
				Content: model.RegCharSet{
					Content: model.CharRange{
						Start: 'a',
						End:   'z',
						WhiteList: false,
					},
				},
			},
			matcher: "AAAA",
			matches: true,
		},
		{
			name: "[az]*, aazzaaaazzzzzz",
			input: model.RegStar{
				Content: model.RegCharSet{
					Content: model.CharList{
						Chars: []rune{'a', 'z'},
						WhiteList: true,
					},
				},
			},
			matcher: "aazzaaaazzzzzz",
			matches: true,
		},
		{
			name: "[az]*, bazzaaaazzzzzz",
			input: model.RegStar{
				Content: model.RegCharSet{
					Content: model.CharList{
						Chars: []rune{'a', 'z'},
						WhiteList: true,
					},
				},
			},
			matcher: "bazzaaaazzzzzz",
			matches: false,
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf("SearchFromAst: %s", td.name), func(t *testing.T) {
			states, start, err := model.CreateCompleteStates(td.input)
			if err != nil {
				t.Error("unexpected error occurred.")
			}
			matched := Search(states, InputBuffer{Input: td.matcher}, start, false, true)
			if matched != td.matches {
				t.Errorf("expected %v, got %v\n", td.matches, matched)
			}
		})
	}
}
