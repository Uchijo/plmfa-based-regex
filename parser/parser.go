package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/uchijo/plmfa-based-regex/model"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

func GenAst(regexStr string) model.RegExp {
	is := antlr.NewInputStream(regexStr)
	lexer := gen.NewPCRELexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := gen.NewPCREParser(stream)
	p.BuildParseTrees = true
	tree := p.Pcre()
	regex := tree.Accept(&RegexBuilder{}).(model.RegExp)

	return regex
}
