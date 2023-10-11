package parser

import (
	"fmt"

	"github.com/uchijo/plmfa-based-regex/model"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

type RegexBuilder struct {
	*gen.BasePCREVisitor
}

func (rb *RegexBuilder) VisitPcre(ctx *gen.PcreContext) interface{} {
	fmt.Println("visiting pcre.")
	result := ctx.Alternation().Accept(rb)
	return result
}

func (rb *RegexBuilder) VisitAccept_(ctx *gen.Accept_Context) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitAlternation(ctx *gen.AlternationContext) interface{} {
	exprs := ctx.AllExpr()
	exprsLen := len(exprs)
	fmt.Printf("length: %v\n", exprsLen)

	// ブランチなし
	if exprsLen == 1 {
		return exprs[0].Accept(rb)
	}

	root := model.RegUnion{
		Left: exprs[0].Accept(rb).(model.RegExp),
	}
	for i, v := range exprs {
		// 1個目は入れてある
		if i == 0 {
			continue
		}

		root.Right = v.Accept(rb).(model.RegExp)

		// 最後の場合、そのまま返す
		if i+1 == exprsLen {
			break
		}

		root = model.RegUnion{
			Left: root,
		}
	}
	return root
}

func (rb *RegexBuilder) VisitAnchor(ctx *gen.AnchorContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitAny(ctx *gen.AnyContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitAnycrlf(ctx *gen.AnycrlfContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitAtom(ctx *gen.AtomContext) interface{} {
	return model.RegString{Content: "a"}
}

func (rb *RegexBuilder) VisitAtomic_group(ctx *gen.Atomic_groupContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitBackreferenceContext(ctx *gen.BackreferenceContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitBacktracking_control(ctx *gen.Backtracking_controlContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitBsr_anycrlf(ctx *gen.Bsr_anycrlfContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitBsr_unicode(ctx *gen.Bsr_unicodeContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCallout(ctx *gen.CalloutContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCapture(ctx *gen.CaptureContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCharacter(ctx *gen.CharacterContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCharacter_class(ctx *gen.Character_classContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCharacter_class_atom(ctx *gen.Character_class_atomContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCharacter_class_range(ctx *gen.Character_class_rangeContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCharacter_class_range_atom(ctx *gen.Character_class_range_atomContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCharacter_type(ctx *gen.Character_typeContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitComment(ctx *gen.CommentContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCommit(ctx *gen.CommitContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitConditional_pattern(ctx *gen.Conditional_patternContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCr(ctx *gen.CrContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCrlf(ctx *gen.CrlfContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitDigit(ctx *gen.DigitContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitDigits(ctx *gen.DigitsContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitElement(ctx *gen.ElementContext) interface{} {
	atom := ctx.Atom().Accept(rb).(model.RegExp)
	quant := ctx.Quantifier()
	if quant != nil {
		quant.Accept(rb)
		return model.RegStar{Content: atom}
	}
	return atom
}

func (rb *RegexBuilder) VisitExpr(ctx *gen.ExprContext) interface{} {
	all := ctx.AllElement()

	// 1個のみの場合そのまま返す
	if len(all) == 1 {
		return all[0].Accept(rb).(model.RegExp)
	}

	result := []model.RegExp{}
	for _, v := range all {
		result = append(result, v.Accept(rb).(model.RegExp))
	}
	return model.RegApp{Contents: result}
}

func (rb *RegexBuilder) VisitFail(ctx *gen.FailContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitHex(ctx *gen.HexContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitLetter(ctx *gen.LetterContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitLetters(ctx *gen.LettersContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitLf(ctx *gen.LfContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitLimit_match(ctx *gen.Limit_matchContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitLimit_recursion(ctx *gen.Limit_recursionContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitLookaround(ctx *gen.LookaroundContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitMark(ctx *gen.MarkContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitMatch_point_reset(ctx *gen.Match_point_resetContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitName(ctx *gen.NameContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitNewline_conventions(ctx *gen.Newline_conventionsContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitNo_auto_possess(ctx *gen.No_auto_possessContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitNo_start_opt(ctx *gen.No_start_optContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitOption_setting(ctx *gen.Option_settingContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitOption_setting_flag(ctx *gen.Option_setting_flagContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitOther(ctx *gen.OtherContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitPosix_character_class(ctx *gen.Posix_character_classContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitPrune(ctx *gen.PruneContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitQuantifier(ctx *gen.QuantifierContext) interface{} {
	if ctx.GetText() != "*" {
		panic("this regex engine doesn't support quantifier except *.")
	}
	return nil
}

func (rb *RegexBuilder) VisitQuoting(ctx *gen.QuotingContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitSkip(ctx *gen.SkipContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitSubroutine_reference(ctx *gen.Subroutine_referenceContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitThen(ctx *gen.ThenContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitUcp(ctx *gen.UcpContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitUtf(ctx *gen.UtfContext) interface{} {
	return nil
}
