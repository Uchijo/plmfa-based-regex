// Code generated from PCRE.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // PCRE

import "github.com/antlr4-go/antlr/v4"

// BasePCREListener is a complete listener for a parse tree produced by PCREParser.
type BasePCREListener struct{}

var _ PCREListener = &BasePCREListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePCREListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePCREListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePCREListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePCREListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPcre is called when production pcre is entered.
func (s *BasePCREListener) EnterPcre(ctx *PcreContext) {}

// ExitPcre is called when production pcre is exited.
func (s *BasePCREListener) ExitPcre(ctx *PcreContext) {}

// EnterAlternation is called when production alternation is entered.
func (s *BasePCREListener) EnterAlternation(ctx *AlternationContext) {}

// ExitAlternation is called when production alternation is exited.
func (s *BasePCREListener) ExitAlternation(ctx *AlternationContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePCREListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePCREListener) ExitExpr(ctx *ExprContext) {}

// EnterElement is called when production element is entered.
func (s *BasePCREListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasePCREListener) ExitElement(ctx *ElementContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasePCREListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasePCREListener) ExitAtom(ctx *AtomContext) {}

// EnterCapture is called when production capture is entered.
func (s *BasePCREListener) EnterCapture(ctx *CaptureContext) {}

// ExitCapture is called when production capture is exited.
func (s *BasePCREListener) ExitCapture(ctx *CaptureContext) {}

// EnterAtomic_group is called when production atomic_group is entered.
func (s *BasePCREListener) EnterAtomic_group(ctx *Atomic_groupContext) {}

// ExitAtomic_group is called when production atomic_group is exited.
func (s *BasePCREListener) ExitAtomic_group(ctx *Atomic_groupContext) {}

// EnterLookaround is called when production lookaround is entered.
func (s *BasePCREListener) EnterLookaround(ctx *LookaroundContext) {}

// ExitLookaround is called when production lookaround is exited.
func (s *BasePCREListener) ExitLookaround(ctx *LookaroundContext) {}

// EnterBackreference is called when production backreference is entered.
func (s *BasePCREListener) EnterBackreference(ctx *BackreferenceContext) {}

// ExitBackreference is called when production backreference is exited.
func (s *BasePCREListener) ExitBackreference(ctx *BackreferenceContext) {}

// EnterSubroutine_reference is called when production subroutine_reference is entered.
func (s *BasePCREListener) EnterSubroutine_reference(ctx *Subroutine_referenceContext) {}

// ExitSubroutine_reference is called when production subroutine_reference is exited.
func (s *BasePCREListener) ExitSubroutine_reference(ctx *Subroutine_referenceContext) {}

// EnterConditional_pattern is called when production conditional_pattern is entered.
func (s *BasePCREListener) EnterConditional_pattern(ctx *Conditional_patternContext) {}

// ExitConditional_pattern is called when production conditional_pattern is exited.
func (s *BasePCREListener) ExitConditional_pattern(ctx *Conditional_patternContext) {}

// EnterComment is called when production comment is entered.
func (s *BasePCREListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasePCREListener) ExitComment(ctx *CommentContext) {}

// EnterQuantifier is called when production quantifier is entered.
func (s *BasePCREListener) EnterQuantifier(ctx *QuantifierContext) {}

// ExitQuantifier is called when production quantifier is exited.
func (s *BasePCREListener) ExitQuantifier(ctx *QuantifierContext) {}

// EnterOption_setting is called when production option_setting is entered.
func (s *BasePCREListener) EnterOption_setting(ctx *Option_settingContext) {}

// ExitOption_setting is called when production option_setting is exited.
func (s *BasePCREListener) ExitOption_setting(ctx *Option_settingContext) {}

// EnterOption_setting_flag is called when production option_setting_flag is entered.
func (s *BasePCREListener) EnterOption_setting_flag(ctx *Option_setting_flagContext) {}

// ExitOption_setting_flag is called when production option_setting_flag is exited.
func (s *BasePCREListener) ExitOption_setting_flag(ctx *Option_setting_flagContext) {}

// EnterBacktracking_control is called when production backtracking_control is entered.
func (s *BasePCREListener) EnterBacktracking_control(ctx *Backtracking_controlContext) {}

// ExitBacktracking_control is called when production backtracking_control is exited.
func (s *BasePCREListener) ExitBacktracking_control(ctx *Backtracking_controlContext) {}

// EnterCallout is called when production callout is entered.
func (s *BasePCREListener) EnterCallout(ctx *CalloutContext) {}

// ExitCallout is called when production callout is exited.
func (s *BasePCREListener) ExitCallout(ctx *CalloutContext) {}

// EnterNewline_conventions is called when production newline_conventions is entered.
func (s *BasePCREListener) EnterNewline_conventions(ctx *Newline_conventionsContext) {}

// ExitNewline_conventions is called when production newline_conventions is exited.
func (s *BasePCREListener) ExitNewline_conventions(ctx *Newline_conventionsContext) {}

// EnterCharacter is called when production character is entered.
func (s *BasePCREListener) EnterCharacter(ctx *CharacterContext) {}

// ExitCharacter is called when production character is exited.
func (s *BasePCREListener) ExitCharacter(ctx *CharacterContext) {}

// EnterCharacter_type is called when production character_type is entered.
func (s *BasePCREListener) EnterCharacter_type(ctx *Character_typeContext) {}

// ExitCharacter_type is called when production character_type is exited.
func (s *BasePCREListener) ExitCharacter_type(ctx *Character_typeContext) {}

// EnterCharacter_class is called when production character_class is entered.
func (s *BasePCREListener) EnterCharacter_class(ctx *Character_classContext) {}

// ExitCharacter_class is called when production character_class is exited.
func (s *BasePCREListener) ExitCharacter_class(ctx *Character_classContext) {}

// EnterCharacter_class_atom is called when production character_class_atom is entered.
func (s *BasePCREListener) EnterCharacter_class_atom(ctx *Character_class_atomContext) {}

// ExitCharacter_class_atom is called when production character_class_atom is exited.
func (s *BasePCREListener) ExitCharacter_class_atom(ctx *Character_class_atomContext) {}

// EnterCharacter_class_range is called when production character_class_range is entered.
func (s *BasePCREListener) EnterCharacter_class_range(ctx *Character_class_rangeContext) {}

// ExitCharacter_class_range is called when production character_class_range is exited.
func (s *BasePCREListener) ExitCharacter_class_range(ctx *Character_class_rangeContext) {}

// EnterCharacter_class_range_atom is called when production character_class_range_atom is entered.
func (s *BasePCREListener) EnterCharacter_class_range_atom(ctx *Character_class_range_atomContext) {}

// ExitCharacter_class_range_atom is called when production character_class_range_atom is exited.
func (s *BasePCREListener) ExitCharacter_class_range_atom(ctx *Character_class_range_atomContext) {}

// EnterPosix_character_class is called when production posix_character_class is entered.
func (s *BasePCREListener) EnterPosix_character_class(ctx *Posix_character_classContext) {}

// ExitPosix_character_class is called when production posix_character_class is exited.
func (s *BasePCREListener) ExitPosix_character_class(ctx *Posix_character_classContext) {}

// EnterAnchor is called when production anchor is entered.
func (s *BasePCREListener) EnterAnchor(ctx *AnchorContext) {}

// ExitAnchor is called when production anchor is exited.
func (s *BasePCREListener) ExitAnchor(ctx *AnchorContext) {}

// EnterMatch_point_reset is called when production match_point_reset is entered.
func (s *BasePCREListener) EnterMatch_point_reset(ctx *Match_point_resetContext) {}

// ExitMatch_point_reset is called when production match_point_reset is exited.
func (s *BasePCREListener) ExitMatch_point_reset(ctx *Match_point_resetContext) {}

// EnterQuoting is called when production quoting is entered.
func (s *BasePCREListener) EnterQuoting(ctx *QuotingContext) {}

// ExitQuoting is called when production quoting is exited.
func (s *BasePCREListener) ExitQuoting(ctx *QuotingContext) {}

// EnterDigits is called when production digits is entered.
func (s *BasePCREListener) EnterDigits(ctx *DigitsContext) {}

// ExitDigits is called when production digits is exited.
func (s *BasePCREListener) ExitDigits(ctx *DigitsContext) {}

// EnterDigit is called when production digit is entered.
func (s *BasePCREListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BasePCREListener) ExitDigit(ctx *DigitContext) {}

// EnterHex is called when production hex is entered.
func (s *BasePCREListener) EnterHex(ctx *HexContext) {}

// ExitHex is called when production hex is exited.
func (s *BasePCREListener) ExitHex(ctx *HexContext) {}

// EnterLetters is called when production letters is entered.
func (s *BasePCREListener) EnterLetters(ctx *LettersContext) {}

// ExitLetters is called when production letters is exited.
func (s *BasePCREListener) ExitLetters(ctx *LettersContext) {}

// EnterLetter is called when production letter is entered.
func (s *BasePCREListener) EnterLetter(ctx *LetterContext) {}

// ExitLetter is called when production letter is exited.
func (s *BasePCREListener) ExitLetter(ctx *LetterContext) {}

// EnterName is called when production name is entered.
func (s *BasePCREListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasePCREListener) ExitName(ctx *NameContext) {}

// EnterOther is called when production other is entered.
func (s *BasePCREListener) EnterOther(ctx *OtherContext) {}

// ExitOther is called when production other is exited.
func (s *BasePCREListener) ExitOther(ctx *OtherContext) {}

// EnterUtf is called when production utf is entered.
func (s *BasePCREListener) EnterUtf(ctx *UtfContext) {}

// ExitUtf is called when production utf is exited.
func (s *BasePCREListener) ExitUtf(ctx *UtfContext) {}

// EnterUcp is called when production ucp is entered.
func (s *BasePCREListener) EnterUcp(ctx *UcpContext) {}

// ExitUcp is called when production ucp is exited.
func (s *BasePCREListener) ExitUcp(ctx *UcpContext) {}

// EnterNo_auto_possess is called when production no_auto_possess is entered.
func (s *BasePCREListener) EnterNo_auto_possess(ctx *No_auto_possessContext) {}

// ExitNo_auto_possess is called when production no_auto_possess is exited.
func (s *BasePCREListener) ExitNo_auto_possess(ctx *No_auto_possessContext) {}

// EnterNo_start_opt is called when production no_start_opt is entered.
func (s *BasePCREListener) EnterNo_start_opt(ctx *No_start_optContext) {}

// ExitNo_start_opt is called when production no_start_opt is exited.
func (s *BasePCREListener) ExitNo_start_opt(ctx *No_start_optContext) {}

// EnterCr is called when production cr is entered.
func (s *BasePCREListener) EnterCr(ctx *CrContext) {}

// ExitCr is called when production cr is exited.
func (s *BasePCREListener) ExitCr(ctx *CrContext) {}

// EnterLf is called when production lf is entered.
func (s *BasePCREListener) EnterLf(ctx *LfContext) {}

// ExitLf is called when production lf is exited.
func (s *BasePCREListener) ExitLf(ctx *LfContext) {}

// EnterCrlf is called when production crlf is entered.
func (s *BasePCREListener) EnterCrlf(ctx *CrlfContext) {}

// ExitCrlf is called when production crlf is exited.
func (s *BasePCREListener) ExitCrlf(ctx *CrlfContext) {}

// EnterAnycrlf is called when production anycrlf is entered.
func (s *BasePCREListener) EnterAnycrlf(ctx *AnycrlfContext) {}

// ExitAnycrlf is called when production anycrlf is exited.
func (s *BasePCREListener) ExitAnycrlf(ctx *AnycrlfContext) {}

// EnterAny is called when production any is entered.
func (s *BasePCREListener) EnterAny(ctx *AnyContext) {}

// ExitAny is called when production any is exited.
func (s *BasePCREListener) ExitAny(ctx *AnyContext) {}

// EnterLimit_match is called when production limit_match is entered.
func (s *BasePCREListener) EnterLimit_match(ctx *Limit_matchContext) {}

// ExitLimit_match is called when production limit_match is exited.
func (s *BasePCREListener) ExitLimit_match(ctx *Limit_matchContext) {}

// EnterLimit_recursion is called when production limit_recursion is entered.
func (s *BasePCREListener) EnterLimit_recursion(ctx *Limit_recursionContext) {}

// ExitLimit_recursion is called when production limit_recursion is exited.
func (s *BasePCREListener) ExitLimit_recursion(ctx *Limit_recursionContext) {}

// EnterBsr_anycrlf is called when production bsr_anycrlf is entered.
func (s *BasePCREListener) EnterBsr_anycrlf(ctx *Bsr_anycrlfContext) {}

// ExitBsr_anycrlf is called when production bsr_anycrlf is exited.
func (s *BasePCREListener) ExitBsr_anycrlf(ctx *Bsr_anycrlfContext) {}

// EnterBsr_unicode is called when production bsr_unicode is entered.
func (s *BasePCREListener) EnterBsr_unicode(ctx *Bsr_unicodeContext) {}

// ExitBsr_unicode is called when production bsr_unicode is exited.
func (s *BasePCREListener) ExitBsr_unicode(ctx *Bsr_unicodeContext) {}

// EnterAccept_ is called when production accept_ is entered.
func (s *BasePCREListener) EnterAccept_(ctx *Accept_Context) {}

// ExitAccept_ is called when production accept_ is exited.
func (s *BasePCREListener) ExitAccept_(ctx *Accept_Context) {}

// EnterFail is called when production fail is entered.
func (s *BasePCREListener) EnterFail(ctx *FailContext) {}

// ExitFail is called when production fail is exited.
func (s *BasePCREListener) ExitFail(ctx *FailContext) {}

// EnterMark is called when production mark is entered.
func (s *BasePCREListener) EnterMark(ctx *MarkContext) {}

// ExitMark is called when production mark is exited.
func (s *BasePCREListener) ExitMark(ctx *MarkContext) {}

// EnterCommit is called when production commit is entered.
func (s *BasePCREListener) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *BasePCREListener) ExitCommit(ctx *CommitContext) {}

// EnterPrune is called when production prune is entered.
func (s *BasePCREListener) EnterPrune(ctx *PruneContext) {}

// ExitPrune is called when production prune is exited.
func (s *BasePCREListener) ExitPrune(ctx *PruneContext) {}

// EnterSkip is called when production skip is entered.
func (s *BasePCREListener) EnterSkip(ctx *SkipContext) {}

// ExitSkip is called when production skip is exited.
func (s *BasePCREListener) ExitSkip(ctx *SkipContext) {}

// EnterThen is called when production then is entered.
func (s *BasePCREListener) EnterThen(ctx *ThenContext) {}

// ExitThen is called when production then is exited.
func (s *BasePCREListener) ExitThen(ctx *ThenContext) {}
