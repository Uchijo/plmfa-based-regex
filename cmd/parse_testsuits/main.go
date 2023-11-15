package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/uchijo/plmfa-based-regex/model"
	"github.com/uchijo/plmfa-based-regex/parser"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

func main() {
	b, err := os.ReadFile("tests/test1")
	if err != nil {
		panic("error" + err.Error())
	}
	rawContents := string(b)

	// テストケースごとで区切る
	unfilteredContents := strings.Split(rawContents, "\n\n")
	filteredTests := []string{}
	for _, v := range unfilteredContents {
		if !isTestCase(v) {
			continue
		}
		filteredTests = append(filteredTests, v)
	}

	fmt.Println("hogehoge!!!!")
	tests := []Test{}
	for _, v := range filteredTests {
		tests = append(tests, extractTest(v))
	}

	fmt.Printf("tests length: %v\n", len(tests))

	fmt.Println("[")
	for _, v := range tests {
		if !v.CanHandle {
			continue
		}
		json, _ := json.Marshal(v)
		fmt.Println(string(json) + ",")
	}
	fmt.Println("]")
}

func isTestCase(input string) bool {
	// 最初が/でない場合違う
	return input[:1] == "/"
}

func extractTest(input string) Test {
	lines := strings.Split(input, "\n")
	rawPattern := lines[0]
	canHandle := canHandlePattern(rawPattern)
	if rawPattern == "/^(\\d+)\\s+IN\\s+SOA\\s+(\\S+)\\s+(\\S+)\\s*\\(\\s*$/" {
		fmt.Println("contains \\d")
		fmt.Printf("can handle: %v\n", canHandle)
		fmt.Printf("extracted: %v\n", extractPattern(rawPattern))
		fmt.Printf("can parse: %v\n", canParsePattern(extractPattern(rawPattern), true))
	}
	stripped := ""
	if canHandle {
		stripped = extractPattern(rawPattern)
		canHandle = canHandle && canParsePattern(stripped, false)
	}
	positives := extractPositives(lines)
	negatives := extractNegatives(lines)
	return Test{
		RawPattern:       rawPattern,
		CanHandle:        canHandle,
		StrippedPattern:  stripped,
		PositiveExamples: positives,
		NegativeExamples: negatives,
	}
}

func extractPositives(input []string) []string {
	var raw []string
	for i, v := range input {
		// 1行目はパターンなので無視
		if i == 0 {
			continue
		}
		if len(v) >= 9 && v[:9] == "\\= Expect" {
			raw = input[1:i]
		}
	}
	if len(raw) == 0 {
		raw = input[1:]
	}

	positives := []string{}
	for _, v := range raw {
		positives = append(positives, strings.TrimSpace(v))
	}
	return positives
}

func extractNegatives(input []string) []string {
	var raw []string
	for i, v := range input {
		// 1行目はパターンなので無視
		if i == 0 {
			continue
		}
		if len(v) >= 9 && v[:9] == "\\= Expect" {
			raw = input[i+1:]
		}
	}

	negs := []string{}
	for _, v := range raw {
		negs = append(negs, strings.TrimSpace(v))
	}
	return negs
}

var caret = regexp.QuoteMeta("^")
var dollar = regexp.QuoteMeta("$")
var pattern = "/" + caret + "(.*)" + dollar + "/"
var compiledRe, _ = regexp.Compile(pattern)

func extractPattern(input string) string {
	matches := compiledRe.FindAllStringSubmatch(input, -1)
	for _, v := range matches {
		return v[1]
	}
	panic("something went wrong")
}

func canHandlePattern(input string) bool {
	// ^$で囲まれているか確認
	// match, err := regexp.Match("/^.*$/[dgimsuy]{0,6}", []byte(input))
	match, err := regexp.Match("^/"+caret+".*"+dollar+"/g?$", []byte(input))
	if err != nil {
		panic("match failed")
	}
	return match
}

// 途中でパニックが発生 -> パース不可 -> recoverでfalseが帰る
// 参考: https://h3poteto.hatenablog.com/entry/2015/12/13/010000
func canParsePattern(
	input string,
	report bool, // for debug
	) bool {
	if input == "" {
		return false
	}

	defer func() {
		if r := recover(); r != nil {
			// do nothing
			if report {
				fmt.Println(r)
			}
		}
	}()

	is := antlr.NewInputStream(input)
	lexer := gen.NewPCRELexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := gen.NewPCREParser(stream)
	p.BuildParseTrees = true
	tree := p.Pcre()
	_ = tree.Accept(&parser.RegexBuilder{}).(model.RegExp)

	return true
}

type Test struct {
	RawPattern       string
	StrippedPattern  string
	PositiveExamples []string
	NegativeExamples []string
	CanHandle        bool
}
