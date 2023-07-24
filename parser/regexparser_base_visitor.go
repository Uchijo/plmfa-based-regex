// Code generated from regexParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // regexParser

import "github.com/antlr4-go/antlr/v4"

type BaseregexParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseregexParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitRegExp(ctx *RegExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitBranch(ctx *BranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitPiece(ctx *PieceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitQuantifier(ctx *QuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitQuantity(ctx *QuantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitQuantRange(ctx *QuantRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitQuantMin(ctx *QuantMinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCharClass(ctx *CharClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCharClassExpr(ctx *CharClassExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCharGroup(ctx *CharGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitPosCharGroup(ctx *PosCharGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCharRange(ctx *CharRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitSeRange(ctx *SeRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCharOrEsc(ctx *CharOrEscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCharClassEsc(ctx *CharClassEscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCatEsc(ctx *CatEscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitComplEsc(ctx *ComplEscContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseregexParserVisitor) VisitCharProp(ctx *CharPropContext) interface{} {
	return v.VisitChildren(ctx)
}
