// Code generated from PCRE.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // PCRE

import "github.com/antlr4-go/antlr/v4"

// PCREListener is a complete listener for a parse tree produced by PCREParser.
type PCREListener interface {
	antlr.ParseTreeListener

	// EnterPcre is called when entering the pcre production.
	EnterPcre(c *PcreContext)

	// EnterAlternation is called when entering the alternation production.
	EnterAlternation(c *AlternationContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterCapture is called when entering the capture production.
	EnterCapture(c *CaptureContext)

	// EnterAtomic_group is called when entering the atomic_group production.
	EnterAtomic_group(c *Atomic_groupContext)

	// EnterLookaround is called when entering the lookaround production.
	EnterLookaround(c *LookaroundContext)

	// EnterBackreference is called when entering the backreference production.
	EnterBackreference(c *BackreferenceContext)

	// EnterSubroutine_reference is called when entering the subroutine_reference production.
	EnterSubroutine_reference(c *Subroutine_referenceContext)

	// EnterConditional_pattern is called when entering the conditional_pattern production.
	EnterConditional_pattern(c *Conditional_patternContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterQuantifier is called when entering the quantifier production.
	EnterQuantifier(c *QuantifierContext)

	// EnterOption_setting is called when entering the option_setting production.
	EnterOption_setting(c *Option_settingContext)

	// EnterOption_setting_flag is called when entering the option_setting_flag production.
	EnterOption_setting_flag(c *Option_setting_flagContext)

	// EnterBacktracking_control is called when entering the backtracking_control production.
	EnterBacktracking_control(c *Backtracking_controlContext)

	// EnterCallout is called when entering the callout production.
	EnterCallout(c *CalloutContext)

	// EnterNewline_conventions is called when entering the newline_conventions production.
	EnterNewline_conventions(c *Newline_conventionsContext)

	// EnterCharacter is called when entering the character production.
	EnterCharacter(c *CharacterContext)

	// EnterCharacter_type is called when entering the character_type production.
	EnterCharacter_type(c *Character_typeContext)

	// EnterCharacter_class is called when entering the character_class production.
	EnterCharacter_class(c *Character_classContext)

	// EnterCharacter_class_atom is called when entering the character_class_atom production.
	EnterCharacter_class_atom(c *Character_class_atomContext)

	// EnterCharacter_class_range is called when entering the character_class_range production.
	EnterCharacter_class_range(c *Character_class_rangeContext)

	// EnterCharacter_class_range_atom is called when entering the character_class_range_atom production.
	EnterCharacter_class_range_atom(c *Character_class_range_atomContext)

	// EnterPosix_character_class is called when entering the posix_character_class production.
	EnterPosix_character_class(c *Posix_character_classContext)

	// EnterAnchor is called when entering the anchor production.
	EnterAnchor(c *AnchorContext)

	// EnterMatch_point_reset is called when entering the match_point_reset production.
	EnterMatch_point_reset(c *Match_point_resetContext)

	// EnterQuoting is called when entering the quoting production.
	EnterQuoting(c *QuotingContext)

	// EnterDigits is called when entering the digits production.
	EnterDigits(c *DigitsContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// EnterHex is called when entering the hex production.
	EnterHex(c *HexContext)

	// EnterLetters is called when entering the letters production.
	EnterLetters(c *LettersContext)

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterOther is called when entering the other production.
	EnterOther(c *OtherContext)

	// EnterUtf is called when entering the utf production.
	EnterUtf(c *UtfContext)

	// EnterUcp is called when entering the ucp production.
	EnterUcp(c *UcpContext)

	// EnterNo_auto_possess is called when entering the no_auto_possess production.
	EnterNo_auto_possess(c *No_auto_possessContext)

	// EnterNo_start_opt is called when entering the no_start_opt production.
	EnterNo_start_opt(c *No_start_optContext)

	// EnterCr is called when entering the cr production.
	EnterCr(c *CrContext)

	// EnterLf is called when entering the lf production.
	EnterLf(c *LfContext)

	// EnterCrlf is called when entering the crlf production.
	EnterCrlf(c *CrlfContext)

	// EnterAnycrlf is called when entering the anycrlf production.
	EnterAnycrlf(c *AnycrlfContext)

	// EnterAny is called when entering the any production.
	EnterAny(c *AnyContext)

	// EnterLimit_match is called when entering the limit_match production.
	EnterLimit_match(c *Limit_matchContext)

	// EnterLimit_recursion is called when entering the limit_recursion production.
	EnterLimit_recursion(c *Limit_recursionContext)

	// EnterBsr_anycrlf is called when entering the bsr_anycrlf production.
	EnterBsr_anycrlf(c *Bsr_anycrlfContext)

	// EnterBsr_unicode is called when entering the bsr_unicode production.
	EnterBsr_unicode(c *Bsr_unicodeContext)

	// EnterAccept_ is called when entering the accept_ production.
	EnterAccept_(c *Accept_Context)

	// EnterFail is called when entering the fail production.
	EnterFail(c *FailContext)

	// EnterMark is called when entering the mark production.
	EnterMark(c *MarkContext)

	// EnterCommit is called when entering the commit production.
	EnterCommit(c *CommitContext)

	// EnterPrune is called when entering the prune production.
	EnterPrune(c *PruneContext)

	// EnterSkip is called when entering the skip production.
	EnterSkip(c *SkipContext)

	// EnterThen is called when entering the then production.
	EnterThen(c *ThenContext)

	// ExitPcre is called when exiting the pcre production.
	ExitPcre(c *PcreContext)

	// ExitAlternation is called when exiting the alternation production.
	ExitAlternation(c *AlternationContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitCapture is called when exiting the capture production.
	ExitCapture(c *CaptureContext)

	// ExitAtomic_group is called when exiting the atomic_group production.
	ExitAtomic_group(c *Atomic_groupContext)

	// ExitLookaround is called when exiting the lookaround production.
	ExitLookaround(c *LookaroundContext)

	// ExitBackreference is called when exiting the backreference production.
	ExitBackreference(c *BackreferenceContext)

	// ExitSubroutine_reference is called when exiting the subroutine_reference production.
	ExitSubroutine_reference(c *Subroutine_referenceContext)

	// ExitConditional_pattern is called when exiting the conditional_pattern production.
	ExitConditional_pattern(c *Conditional_patternContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitQuantifier is called when exiting the quantifier production.
	ExitQuantifier(c *QuantifierContext)

	// ExitOption_setting is called when exiting the option_setting production.
	ExitOption_setting(c *Option_settingContext)

	// ExitOption_setting_flag is called when exiting the option_setting_flag production.
	ExitOption_setting_flag(c *Option_setting_flagContext)

	// ExitBacktracking_control is called when exiting the backtracking_control production.
	ExitBacktracking_control(c *Backtracking_controlContext)

	// ExitCallout is called when exiting the callout production.
	ExitCallout(c *CalloutContext)

	// ExitNewline_conventions is called when exiting the newline_conventions production.
	ExitNewline_conventions(c *Newline_conventionsContext)

	// ExitCharacter is called when exiting the character production.
	ExitCharacter(c *CharacterContext)

	// ExitCharacter_type is called when exiting the character_type production.
	ExitCharacter_type(c *Character_typeContext)

	// ExitCharacter_class is called when exiting the character_class production.
	ExitCharacter_class(c *Character_classContext)

	// ExitCharacter_class_atom is called when exiting the character_class_atom production.
	ExitCharacter_class_atom(c *Character_class_atomContext)

	// ExitCharacter_class_range is called when exiting the character_class_range production.
	ExitCharacter_class_range(c *Character_class_rangeContext)

	// ExitCharacter_class_range_atom is called when exiting the character_class_range_atom production.
	ExitCharacter_class_range_atom(c *Character_class_range_atomContext)

	// ExitPosix_character_class is called when exiting the posix_character_class production.
	ExitPosix_character_class(c *Posix_character_classContext)

	// ExitAnchor is called when exiting the anchor production.
	ExitAnchor(c *AnchorContext)

	// ExitMatch_point_reset is called when exiting the match_point_reset production.
	ExitMatch_point_reset(c *Match_point_resetContext)

	// ExitQuoting is called when exiting the quoting production.
	ExitQuoting(c *QuotingContext)

	// ExitDigits is called when exiting the digits production.
	ExitDigits(c *DigitsContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)

	// ExitHex is called when exiting the hex production.
	ExitHex(c *HexContext)

	// ExitLetters is called when exiting the letters production.
	ExitLetters(c *LettersContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitOther is called when exiting the other production.
	ExitOther(c *OtherContext)

	// ExitUtf is called when exiting the utf production.
	ExitUtf(c *UtfContext)

	// ExitUcp is called when exiting the ucp production.
	ExitUcp(c *UcpContext)

	// ExitNo_auto_possess is called when exiting the no_auto_possess production.
	ExitNo_auto_possess(c *No_auto_possessContext)

	// ExitNo_start_opt is called when exiting the no_start_opt production.
	ExitNo_start_opt(c *No_start_optContext)

	// ExitCr is called when exiting the cr production.
	ExitCr(c *CrContext)

	// ExitLf is called when exiting the lf production.
	ExitLf(c *LfContext)

	// ExitCrlf is called when exiting the crlf production.
	ExitCrlf(c *CrlfContext)

	// ExitAnycrlf is called when exiting the anycrlf production.
	ExitAnycrlf(c *AnycrlfContext)

	// ExitAny is called when exiting the any production.
	ExitAny(c *AnyContext)

	// ExitLimit_match is called when exiting the limit_match production.
	ExitLimit_match(c *Limit_matchContext)

	// ExitLimit_recursion is called when exiting the limit_recursion production.
	ExitLimit_recursion(c *Limit_recursionContext)

	// ExitBsr_anycrlf is called when exiting the bsr_anycrlf production.
	ExitBsr_anycrlf(c *Bsr_anycrlfContext)

	// ExitBsr_unicode is called when exiting the bsr_unicode production.
	ExitBsr_unicode(c *Bsr_unicodeContext)

	// ExitAccept_ is called when exiting the accept_ production.
	ExitAccept_(c *Accept_Context)

	// ExitFail is called when exiting the fail production.
	ExitFail(c *FailContext)

	// ExitMark is called when exiting the mark production.
	ExitMark(c *MarkContext)

	// ExitCommit is called when exiting the commit production.
	ExitCommit(c *CommitContext)

	// ExitPrune is called when exiting the prune production.
	ExitPrune(c *PruneContext)

	// ExitSkip is called when exiting the skip production.
	ExitSkip(c *SkipContext)

	// ExitThen is called when exiting the then production.
	ExitThen(c *ThenContext)
}
