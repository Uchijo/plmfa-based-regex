package parser

import gen "github.com/uchijo/plmfa-based-regex/parser/gen"

type RegexBuilder struct {}

var _ gen.BasePCREVisitor = (*RegexBuilder)(nil)

func (rb *RegexBuilder) VisitPcre(ctx gen.PcreContext) interface {} {
	return ctx.Alternation().Accept(rb)
}
