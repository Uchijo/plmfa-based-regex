package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/uchijo/plmfa-based-regex/eval"
	"github.com/uchijo/plmfa-based-regex/model"
	"github.com/uchijo/plmfa-based-regex/parser"
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
	showLog   bool
}

var args Args

func init() {
	var epsilonSemantics = flag.Bool("e", false, "use epsilon semantics")
	var noRecover = flag.Bool("noRecover", false, "doesnt recover from panic when true")
	var showLog = flag.Bool("showLog", false, "show log on matching")
	flag.Parse()
	rawArgs := flag.Args()
	args = Args{
		input:     rawArgs[0],
		regex:     rawArgs[1],
		eSem:      *epsilonSemantics,
		noRecover: *noRecover,
		showLog:   *showLog,
	}
}

func main() {
	// エラーで死ぬか、エラーがあったことをjsonで通知するかの分岐
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

	regex := parser.GenAst(args.regex)

	states, start, _ := model.CreateCompleteStates(regex)
	input := eval.InputBuffer{
		Input: args.input,
	}

	matched := eval.Search(states, input, start, args.eSem, args.showLog)
	result := Output{
		Match: matched,
		Error: false,
		Input: args.input,
		Regex: args.regex,
	}
	byteText, _ := json.Marshal(result)
	fmt.Println(string(byteText))
}
