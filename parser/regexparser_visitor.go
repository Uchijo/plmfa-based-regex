// Code generated from regexParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // regexParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by regexParser.
type regexParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by regexParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by regexParser#regExp.
	VisitRegExp(ctx *RegExpContext) interface{}

	// Visit a parse tree produced by regexParser#branch.
	VisitBranch(ctx *BranchContext) interface{}

	// Visit a parse tree produced by regexParser#piece.
	VisitPiece(ctx *PieceContext) interface{}

	// Visit a parse tree produced by regexParser#quantifier.
	VisitQuantifier(ctx *QuantifierContext) interface{}

	// Visit a parse tree produced by regexParser#quantity.
	VisitQuantity(ctx *QuantityContext) interface{}

	// Visit a parse tree produced by regexParser#quantRange.
	VisitQuantRange(ctx *QuantRangeContext) interface{}

	// Visit a parse tree produced by regexParser#quantMin.
	VisitQuantMin(ctx *QuantMinContext) interface{}

	// Visit a parse tree produced by regexParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by regexParser#charClass.
	VisitCharClass(ctx *CharClassContext) interface{}

	// Visit a parse tree produced by regexParser#charClassExpr.
	VisitCharClassExpr(ctx *CharClassExprContext) interface{}

	// Visit a parse tree produced by regexParser#charGroup.
	VisitCharGroup(ctx *CharGroupContext) interface{}

	// Visit a parse tree produced by regexParser#posCharGroup.
	VisitPosCharGroup(ctx *PosCharGroupContext) interface{}

	// Visit a parse tree produced by regexParser#charRange.
	VisitCharRange(ctx *CharRangeContext) interface{}

	// Visit a parse tree produced by regexParser#seRange.
	VisitSeRange(ctx *SeRangeContext) interface{}

	// Visit a parse tree produced by regexParser#charOrEsc.
	VisitCharOrEsc(ctx *CharOrEscContext) interface{}

	// Visit a parse tree produced by regexParser#charClassEsc.
	VisitCharClassEsc(ctx *CharClassEscContext) interface{}

	// Visit a parse tree produced by regexParser#catEsc.
	VisitCatEsc(ctx *CatEscContext) interface{}

	// Visit a parse tree produced by regexParser#complEsc.
	VisitComplEsc(ctx *ComplEscContext) interface{}

	// Visit a parse tree produced by regexParser#charProp.
	VisitCharProp(ctx *CharPropContext) interface{}
}
