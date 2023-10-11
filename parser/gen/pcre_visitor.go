// Code generated from PCRE.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // PCRE

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PCREParser.
type PCREVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PCREParser#pcre.
	VisitPcre(ctx *PcreContext) interface{}

	// Visit a parse tree produced by PCREParser#alternation.
	VisitAlternation(ctx *AlternationContext) interface{}

	// Visit a parse tree produced by PCREParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by PCREParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by PCREParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by PCREParser#capture.
	VisitCapture(ctx *CaptureContext) interface{}

	// Visit a parse tree produced by PCREParser#atomic_group.
	VisitAtomic_group(ctx *Atomic_groupContext) interface{}

	// Visit a parse tree produced by PCREParser#lookaround.
	VisitLookaround(ctx *LookaroundContext) interface{}

	// Visit a parse tree produced by PCREParser#backreference.
	VisitBackreference(ctx *BackreferenceContext) interface{}

	// Visit a parse tree produced by PCREParser#subroutine_reference.
	VisitSubroutine_reference(ctx *Subroutine_referenceContext) interface{}

	// Visit a parse tree produced by PCREParser#conditional_pattern.
	VisitConditional_pattern(ctx *Conditional_patternContext) interface{}

	// Visit a parse tree produced by PCREParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by PCREParser#quantifier.
	VisitQuantifier(ctx *QuantifierContext) interface{}

	// Visit a parse tree produced by PCREParser#option_setting.
	VisitOption_setting(ctx *Option_settingContext) interface{}

	// Visit a parse tree produced by PCREParser#option_setting_flag.
	VisitOption_setting_flag(ctx *Option_setting_flagContext) interface{}

	// Visit a parse tree produced by PCREParser#backtracking_control.
	VisitBacktracking_control(ctx *Backtracking_controlContext) interface{}

	// Visit a parse tree produced by PCREParser#callout.
	VisitCallout(ctx *CalloutContext) interface{}

	// Visit a parse tree produced by PCREParser#newline_conventions.
	VisitNewline_conventions(ctx *Newline_conventionsContext) interface{}

	// Visit a parse tree produced by PCREParser#character.
	VisitCharacter(ctx *CharacterContext) interface{}

	// Visit a parse tree produced by PCREParser#character_type.
	VisitCharacter_type(ctx *Character_typeContext) interface{}

	// Visit a parse tree produced by PCREParser#character_class.
	VisitCharacter_class(ctx *Character_classContext) interface{}

	// Visit a parse tree produced by PCREParser#character_class_atom.
	VisitCharacter_class_atom(ctx *Character_class_atomContext) interface{}

	// Visit a parse tree produced by PCREParser#character_class_range.
	VisitCharacter_class_range(ctx *Character_class_rangeContext) interface{}

	// Visit a parse tree produced by PCREParser#character_class_range_atom.
	VisitCharacter_class_range_atom(ctx *Character_class_range_atomContext) interface{}

	// Visit a parse tree produced by PCREParser#posix_character_class.
	VisitPosix_character_class(ctx *Posix_character_classContext) interface{}

	// Visit a parse tree produced by PCREParser#anchor.
	VisitAnchor(ctx *AnchorContext) interface{}

	// Visit a parse tree produced by PCREParser#match_point_reset.
	VisitMatch_point_reset(ctx *Match_point_resetContext) interface{}

	// Visit a parse tree produced by PCREParser#quoting.
	VisitQuoting(ctx *QuotingContext) interface{}

	// Visit a parse tree produced by PCREParser#digits.
	VisitDigits(ctx *DigitsContext) interface{}

	// Visit a parse tree produced by PCREParser#digit.
	VisitDigit(ctx *DigitContext) interface{}

	// Visit a parse tree produced by PCREParser#hex.
	VisitHex(ctx *HexContext) interface{}

	// Visit a parse tree produced by PCREParser#letters.
	VisitLetters(ctx *LettersContext) interface{}

	// Visit a parse tree produced by PCREParser#letter.
	VisitLetter(ctx *LetterContext) interface{}

	// Visit a parse tree produced by PCREParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by PCREParser#other.
	VisitOther(ctx *OtherContext) interface{}

	// Visit a parse tree produced by PCREParser#utf.
	VisitUtf(ctx *UtfContext) interface{}

	// Visit a parse tree produced by PCREParser#ucp.
	VisitUcp(ctx *UcpContext) interface{}

	// Visit a parse tree produced by PCREParser#no_auto_possess.
	VisitNo_auto_possess(ctx *No_auto_possessContext) interface{}

	// Visit a parse tree produced by PCREParser#no_start_opt.
	VisitNo_start_opt(ctx *No_start_optContext) interface{}

	// Visit a parse tree produced by PCREParser#cr.
	VisitCr(ctx *CrContext) interface{}

	// Visit a parse tree produced by PCREParser#lf.
	VisitLf(ctx *LfContext) interface{}

	// Visit a parse tree produced by PCREParser#crlf.
	VisitCrlf(ctx *CrlfContext) interface{}

	// Visit a parse tree produced by PCREParser#anycrlf.
	VisitAnycrlf(ctx *AnycrlfContext) interface{}

	// Visit a parse tree produced by PCREParser#any.
	VisitAny(ctx *AnyContext) interface{}

	// Visit a parse tree produced by PCREParser#limit_match.
	VisitLimit_match(ctx *Limit_matchContext) interface{}

	// Visit a parse tree produced by PCREParser#limit_recursion.
	VisitLimit_recursion(ctx *Limit_recursionContext) interface{}

	// Visit a parse tree produced by PCREParser#bsr_anycrlf.
	VisitBsr_anycrlf(ctx *Bsr_anycrlfContext) interface{}

	// Visit a parse tree produced by PCREParser#bsr_unicode.
	VisitBsr_unicode(ctx *Bsr_unicodeContext) interface{}

	// Visit a parse tree produced by PCREParser#accept_.
	VisitAccept_(ctx *Accept_Context) interface{}

	// Visit a parse tree produced by PCREParser#fail.
	VisitFail(ctx *FailContext) interface{}

	// Visit a parse tree produced by PCREParser#mark.
	VisitMark(ctx *MarkContext) interface{}

	// Visit a parse tree produced by PCREParser#commit.
	VisitCommit(ctx *CommitContext) interface{}

	// Visit a parse tree produced by PCREParser#prune.
	VisitPrune(ctx *PruneContext) interface{}

	// Visit a parse tree produced by PCREParser#skip.
	VisitSkip(ctx *SkipContext) interface{}

	// Visit a parse tree produced by PCREParser#then.
	VisitThen(ctx *ThenContext) interface{}
}
