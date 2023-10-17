package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/uchijo/plmfa-based-regex/eval"
	"github.com/uchijo/plmfa-based-regex/model"
	"github.com/uchijo/plmfa-based-regex/parser"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

type Output struct {
	Match bool   `json:"match"`
	Error bool   `json:"error"`
	Regex string `json:"regex"`
	Input string `json:"input"`
}

type Args struct {
	input     string
	regex     string
	eSem      bool
	noRecover bool
}

func main() {
	var epsilonSemantics = flag.Bool("e", false, "use epsilon semantics")
	var noRecover = flag.Bool("noRecover", false, "doesnt recover from panic when true")
	flag.Parse()
	rawArgs := flag.Args()
	args := Args{
		input:     rawArgs[0],
		regex:     rawArgs[1],
		eSem:      *epsilonSemantics,
		noRecover: *noRecover,
	}

	defer func() {
		if args.noRecover {
			return
		}
		if r := recover(); r != nil {
			result := Output{
				Error: true,
			}
			byteText, _ := json.Marshal(result)
			fmt.Println(string(byteText))
		}
	}()

	is := antlr.NewInputStream(args.regex)
	lexer := gen.NewPCRELexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := gen.NewPCREParser(stream)
	p.BuildParseTrees = true
	tree := p.Pcre()
	regex := tree.Accept(&parser.RegexBuilder{}).(model.RegExp)

	states, start, _ := model.CreateCompleteStates(regex)
	input := eval.InputBuffer{
		Input: args.input,
	}

	matched := eval.Search(states, input, start, args.eSem)
	result := Output{
		Match: matched,
		Error: false,
		Input: args.input,
		Regex: args.regex,
	}
	byteText, _ := json.Marshal(result)
	fmt.Println(string(byteText))
}
