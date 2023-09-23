package parser

import (
	"github.com/Uchijo/plmfa-based-regex/model"
)

type RegexBuilder struct {
	*BaseregexParserVisitor
}

func (rv *RegexBuilder) VisitRoot(ctx *RootContext) interface{} {
	return ctx.RegExp().Accept(rv)
}

func (rv *RegexBuilder) VisitRegExp(ctx *RegExpContext) interface{} {
	branches := ctx.AllBranch()
	branchLength := len(branches)

	if branchLength <= 1 {
		return branches[0].Accept(rv).(model.RegExp)
	}

	var current *model.RegUnion
	base := model.RegUnion{}
	current = &base

	for i, v := range branches {
		parsed := v.Accept(rv).(model.RegExp)
		current.Left = parsed
		if hasNext(i, branchLength) {
			newUnion := model.RegUnion{}
			current.Right = newUnion
			current = &newUnion
		} else {
			current.Right = branches[branchLength-1].Accept(rv).(model.RegExp)
			break
		}
	}
	return base
}

func hasNext(currentIndex, arrayLength int) bool {
	return arrayLength-currentIndex > 2
}

func (rv *RegexBuilder) VisitBranch(ctx *BranchContext) interface{} {
	pieces := ctx.AllPiece()
	if len(pieces) == 1 {
		return pieces[0].Accept(rv).(model.RegExp)
	}

	contents := []model.RegExp{}
	for _, v := range pieces {
		contents = append(contents, v.Accept(rv).(model.RegExp))
	}
	return model.RegApp{Contents: contents}
}

func (rv *RegexBuilder) VisitPiece(ctx *PieceContext) interface{} {
	atom := ctx.Atom().Accept(rv).(model.RegExp)
	q := ctx.Quantifier()
	if q == nil {
		return atom
	}
	q.Accept(rv)
	return model.RegStar{Content: atom}
}

// Visit a parse tree produced by regexParser#quantifier.
func (rv *RegexBuilder) VisitQuantifier(ctx *QuantifierContext) interface{} {
	star := ctx.STAR()
	if star == nil {
		panic("quantifier except star is not supported!")
	}
	return nil
}

// Visit a parse tree produced by regexParser#quantity.
func (rv *RegexBuilder) VisitQuantity(ctx *QuantityContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#quantRange.
func (rv *RegexBuilder) VisitQuantRange(ctx *QuantRangeContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#quantMin.
func (rv *RegexBuilder) VisitQuantMin(ctx *QuantMinContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#atom.
func (rv *RegexBuilder) VisitAtom(ctx *AtomContext) interface{} {
	if regex := ctx.RegExp(); regex != nil {
		return regex.Accept(rv).(model.RegExp)
	}
	if char := ctx.Char(); char != nil {
		c := char.GetText()
		return model.RegString{Content: c}
	}
	panic("char class!")
}

// Visit a parse tree produced by regexParser#charClass.
func (rv *RegexBuilder) VisitCharClass(ctx *CharClassContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#charClassExpr.
func (rv *RegexBuilder) VisitCharClassExpr(ctx *CharClassExprContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#charGroup.
func (rv *RegexBuilder) VisitCharGroup(ctx *CharGroupContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#posCharGroup.
func (rv *RegexBuilder) VisitPosCharGroup(ctx *PosCharGroupContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#charRange.
func (rv *RegexBuilder) VisitCharRange(ctx *CharRangeContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#seRange.
func (rv *RegexBuilder) VisitSeRange(ctx *SeRangeContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#charOrEsc.
func (rv *RegexBuilder) VisitCharOrEsc(ctx *CharOrEscContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#charClassEsc.
func (rv *RegexBuilder) VisitCharClassEsc(ctx *CharClassEscContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#catEsc.
func (rv *RegexBuilder) VisitCatEsc(ctx *CatEscContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#complEsc.
func (rv *RegexBuilder) VisitComplEsc(ctx *ComplEscContext) interface{} {
	return nil
}

// Visit a parse tree produced by regexParser#charProp.
func (rv *RegexBuilder) VisitCharProp(ctx *CharPropContext) interface{} {
	return nil
}
