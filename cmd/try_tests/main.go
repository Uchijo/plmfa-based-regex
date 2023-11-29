package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/uchijo/plmfa-based-regex/eval"
	"github.com/uchijo/plmfa-based-regex/model"
	"github.com/uchijo/plmfa-based-regex/parser"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

type Test struct {
	RawPattern       string
	StrippedPattern  string
	PositiveExamples []string
	NegativeExamples []string
	CanHandle        bool
}

type Report struct {
	RawPttern       string
	StrippedPattern string
	Input           string
	Expected        bool
	Got             bool
}

func main() {
	filename := os.Args[1]
	if filename == "" {
		panic("give me filename!")
	}

	rawJson, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	tests := []Test{}
	json.Unmarshal(rawJson, &tests)
	fmt.Println("tests length: ", len(tests))

	// 変な結果を入れておく
	reports := []Report{}

	for _, test := range tests {
		is := antlr.NewInputStream(test.StrippedPattern)
		lexer := gen.NewPCRELexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		p := gen.NewPCREParser(stream)
		p.BuildParseTrees = true
		tree := p.Pcre()
		regex := tree.Accept(&parser.RegexBuilder{}).(model.RegExp)

		states, start, _ := model.CreateCompleteStates(regex)

		// 正例について回していく
		for _, matcher := range test.PositiveExamples {
			input := eval.InputBuffer{
				Input: matcher,
			}

			matched := eval.Search(
				states,
				input,
				start,
				// epsilon semantics
				false,
				// show log
				false,
			)

			if !matched {
				reports = append(reports, Report{
					RawPttern:       test.RawPattern,
					StrippedPattern: test.StrippedPattern,
					Expected:        true,
					Got:             false,
					Input:           matcher,
				})
			}
		}

		// 負例について回していく
		for _, matcher := range test.NegativeExamples {
			input := eval.InputBuffer{
				Input: matcher,
			}

			matched := eval.Search(
				states,
				input,
				start,
				// epsilon semantics
				false,
				// show log
				false,
			)

			if matched {
				reports = append(reports, Report{
					RawPttern:       test.RawPattern,
					StrippedPattern: test.StrippedPattern,
					Expected:        false,
					Got:             true,
					Input:           matcher,
				})
			}
		}
	}

	marshaled, _ := json.Marshal(reports)
	fmt.Println(string(marshaled))
}
