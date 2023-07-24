package parser

import "fmt"

type RegexBuilder struct {
	*BaseregexParserVisitor
}

func (rv *RegexBuilder) VisitRoot(ctx *RootContext) interface{} {
	fmt.Printf("%v\n", ctx)
	return ctx.
}

func (rv *RegexBuilder) VisitQuantifier(ctx *QuantifierContext) interface{} {
	ctx.GetTokens()
}

// func (rv *RegexBuilder) VisitChar
