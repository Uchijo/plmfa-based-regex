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
			name:   "test lookahead and back-reference",
			input:  "aaaaabaaaaa",
			regex:  "(.*)b(?=aaaaa)\\1",
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
					t.Errorf("should not get into panic")
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

			matched := Search(states, input, start, td.useEpsilonSem)
			if matched != td.output {
				t.Errorf("expected result is %v, but got %v", td.output, matched)
			}
		})
	}
}
