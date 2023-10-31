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
