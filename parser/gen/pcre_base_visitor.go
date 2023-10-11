// Code generated from PCRE.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // PCRE

import "github.com/antlr4-go/antlr/v4"

type BasePCREVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePCREVisitor) VisitPcre(ctx *PcreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAlternation(ctx *AlternationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCapture(ctx *CaptureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAtomic_group(ctx *Atomic_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLookaround(ctx *LookaroundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitBackreference(ctx *BackreferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitSubroutine_reference(ctx *Subroutine_referenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitConditional_pattern(ctx *Conditional_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitQuantifier(ctx *QuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOption_setting(ctx *Option_settingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOption_setting_flag(ctx *Option_setting_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitBacktracking_control(ctx *Backtracking_controlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCallout(ctx *CalloutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNewline_conventions(ctx *Newline_conventionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCharacter(ctx *CharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCharacter_type(ctx *Character_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCharacter_class(ctx *Character_classContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCharacter_class_atom(ctx *Character_class_atomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCharacter_class_range(ctx *Character_class_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCharacter_class_range_atom(ctx *Character_class_range_atomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitPosix_character_class(ctx *Posix_character_classContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAnchor(ctx *AnchorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitMatch_point_reset(ctx *Match_point_resetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitQuoting(ctx *QuotingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitDigits(ctx *DigitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitDigit(ctx *DigitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitHex(ctx *HexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLetters(ctx *LettersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLetter(ctx *LetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitOther(ctx *OtherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitUtf(ctx *UtfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitUcp(ctx *UcpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNo_auto_possess(ctx *No_auto_possessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitNo_start_opt(ctx *No_start_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCr(ctx *CrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLf(ctx *LfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCrlf(ctx *CrlfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAnycrlf(ctx *AnycrlfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAny(ctx *AnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLimit_match(ctx *Limit_matchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitLimit_recursion(ctx *Limit_recursionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitBsr_anycrlf(ctx *Bsr_anycrlfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitBsr_unicode(ctx *Bsr_unicodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitAccept_(ctx *Accept_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitFail(ctx *FailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitMark(ctx *MarkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitCommit(ctx *CommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitPrune(ctx *PruneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitSkip(ctx *SkipContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePCREVisitor) VisitThen(ctx *ThenContext) interface{} {
	return v.VisitChildren(ctx)
}
