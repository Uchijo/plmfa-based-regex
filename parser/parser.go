package parser

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/uchijo/plmfa-based-regex/model"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

type RegexBuilder struct {
	*gen.BasePCREVisitor
}

func (rb *RegexBuilder) VisitPcre(ctx *gen.PcreContext) interface{} {
	captureIndex = 1
	memIndex = 0
	result := ctx.Alternation().Accept(rb)
	return result
}

func (rb *RegexBuilder) VisitAccept_(ctx *gen.Accept_Context) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitAlternation(ctx *gen.AlternationContext) interface{} {
	exprs := ctx.AllExpr()
	exprsLen := len(exprs)

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

var captureIndex = 1

func (rb *RegexBuilder) VisitAtom(ctx *gen.AtomContext) interface{} {
	if capture := ctx.Capture(); capture != nil {
		// 非キャプチャの場合
		if ctx.GetText()[:3] == "(?:" {
			return capture.Accept(rb).(model.RegExp)
		}
		retval := model.RegCapture{MemoryIndex: captureIndex}
		captureIndex++
		retval.Content = capture.Accept(rb).(model.RegExp)
		return retval
	}
	if lookaround := ctx.Lookaround(); lookaround != nil {
		return lookaround.Accept(rb).(model.RegExp)
	}
	if ch := ctx.Character(); ch != nil {
		return ch.Accept(rb).(model.RegExp)
	}
	if cht := ctx.Character_type(); cht != nil {
		return cht.Accept(rb).(model.RegExp)
	}
	if lt := ctx.Letter(); lt != nil {
		return lt.Accept(rb).(model.RegExp)
	}
	if dg := ctx.Digit(); dg != nil {
		return dg.Accept(rb).(model.RegExp)
	}
	if chc := ctx.Character_class(); chc != nil {
		return chc.Accept(rb).(model.RegExp)
	}
	if ac := ctx.Anchor(); ac != nil {
		return model.RegSkip{}
	}
	if qt := ctx.Quoting(); qt != nil {
		return qt.Accept(rb).(model.RegExp)
	}
	panic("parse error in VisitAtom. cannot parse " + ctx.GetText())
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
	return ctx.Alternation().Accept(rb).(model.RegExp)
}

func (rb *RegexBuilder) VisitCharacter(ctx *gen.CharacterContext) interface{} {
	txt := ctx.GetText()
	if txt[:1] == "\\" {
		digitList := []string{}
		for _, v := range ctx.AllDigit() {
			digitList = append(digitList, v.GetText())
		}
		digits, err := strconv.Atoi(strings.Join(digitList, ""))
		if err != nil {
			panic("parse error.")
		}

		return model.RegCapRef{
			MemIndex: digits,
		}
	}
	panic("cannot parse " + txt)
}

// fixme: []] []a] []-a] [^]] [^]a] [^]-a] などのパターンに特別に対応する必要があるが、できていない
func (rb *RegexBuilder) VisitCharacter_class(ctx *gen.Character_classContext) interface{} {
	negate := ctx.GetNegate() != nil
	atoms := ctx.AllCharacter_class_atom()
	regexes := []model.CharContainer{}
	for _, v := range atoms {
		if parsed := v.Character_class_range(); parsed != nil {
			result := parsed.Accept(rb).(model.CharRange)
			regexes = append(regexes, result)
			continue
		}
		if parsed := v.Posix_character_class(); parsed != nil {
			panic("Posix_character_class is not supported.")
		}
		if parsed := v.Character(); parsed != nil {
			panic("Character in CharacterClass is not supported.")
		}
		if parsed := v.Character_type(); parsed != nil {
			panic("Character_type in CharacterClass is not supported.")
		}
		// 1文字ずつ列挙した場合を拾う
		t := v.GetText()
		regexes = append(regexes, model.CharList{Chars: []rune(t), WhiteList: true})
	}
	container := model.CompositeContainer{Contents: regexes, WhiteList: !negate}
	return model.RegCharSet{Content: container}
}

func (rb *RegexBuilder) VisitCharacter_class_atom(ctx *gen.Character_class_atomContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitCharacter_class_range(ctx *gen.Character_class_rangeContext) interface{} {
	atoms := ctx.AllCharacter_class_range_atom()
	first := atoms[0].Accept(rb).(rune)
	second := atoms[1].Accept(rb).(rune)
	return model.CharRange{
		Start: first,
		End:   second,
		// 実際はCompositeContainerでwhitelist判定するので実質ここは常にtrue
		WhiteList: true,
	}
}

func (rb *RegexBuilder) VisitCharacter_class_range_atom(ctx *gen.Character_class_range_atomContext) interface{} {
	r, _ := utf8.DecodeRuneInString(ctx.GetText())
	return r
}

func (rb *RegexBuilder) VisitCharacter_type(ctx *gen.Character_typeContext) interface{} {
	txt := ctx.GetText()
	if txt == "." {
		return model.RegArb{}
	}
	if txt == "\\d" {
		return model.RegCharSet{
			Content: model.CharRange{
				Start:     '0',
				End:       '9',
				WhiteList: true,
			},
		}
	}
	if txt == "\\D" {
		return model.RegCharSet{
			Content: model.CharRange{
				Start:     '0',
				End:       '9',
				WhiteList: false,
			},
		}
	}
	if txt == "\\w" {
		return model.RegCharSet{
			Content: model.CompositeContainer{
				Contents: []model.CharContainer{
					model.CharRange{
						Start:     '0',
						End:       '9',
						WhiteList: true,
					},
					model.CharRange{
						Start:     'a',
						End:       'z',
						WhiteList: true,
					},
					model.CharRange{
						Start:     'A',
						End:       'Z',
						WhiteList: true,
					},
					model.CharList{
						Chars: []rune{
							'_',
						},
						WhiteList: true,
					},
				},
				WhiteList: true,
			},
		}
	}
	if txt == "\\W" {
		return model.RegCharSet{
			Content: model.CompositeContainer{
				Contents: []model.CharContainer{
					model.CharRange{
						Start:     '0',
						End:       '9',
						WhiteList: true,
					},
					model.CharRange{
						Start:     'a',
						End:       'z',
						WhiteList: true,
					},
					model.CharRange{
						Start:     'A',
						End:       'Z',
						WhiteList: true,
					},
					model.CharList{
						Chars: []rune{
							'_',
						},
						WhiteList: true,
					},
				},
				WhiteList: false,
			},
		}
	}
	if txt == "\\R" {
		return model.RegCharSet{
			Content: model.CompositeContainer{
				WhiteList: true,
				Contents: []model.CharContainer{
					model.CharList{
						Chars: []rune{
							'\n',
							'\r',
						},
					},
					model.CRLF{},
				},
			},
		}
	}
	if txt == "\\N" {
		return model.RegCharSet{
			Content: model.CompositeContainer{
				WhiteList: false,
				Contents: []model.CharContainer{
					model.CharList{
						Chars: []rune{
							'\n',
							'\r',
						},
					},
					model.CRLF{},
				},
			},
		}
	}
	if txt == "\\s" {
		return model.RegCharSet{
			Content: model.CharList{
				WhiteList: true,
				Chars: []rune{
					' ',
					'\t',
					'\f',
					'\r',
					'\n',
					'\v',
				},
			},
		}
	}
	if txt == "\\S" {
		return model.RegCharSet{
			Content: model.CharList{
				WhiteList: false,
				Chars: []rune{
					' ',
					'\t',
					'\f',
					'\r',
					'\n',
					'\v',
				},
			},
		}
	}
	if txt == "\\h" {
		return model.RegCharSet{
			Content: model.CharList{
				WhiteList: true,
				Chars: []rune{
					' ',
					'\t',
				},
			},
		}
	}
	if txt == "\\H" {
		return model.RegCharSet{
			Content: model.CharList{
				WhiteList: false,
				Chars: []rune{
					' ',
					'\t',
				},
			},
		}
	}
	if txt == "\\v" {
		return model.RegCharSet{
			Content: model.CharList{
				WhiteList: true,
				Chars: []rune{
					'\v',
				},
			},
		}
	}
	if txt == "\\V" {
		return model.RegCharSet{
			Content: model.CharList{
				WhiteList: false,
				Chars: []rune{
					'\v',
				},
			},
		}
	}
	panic("cannot parse " + txt)
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
	return model.RegString{Content: ctx.GetText()}
}

func (rb *RegexBuilder) VisitDigits(ctx *gen.DigitsContext) interface{} {
	return nil
}

func (rb *RegexBuilder) VisitElement(ctx *gen.ElementContext) interface{} {
	atom := ctx.Atom().Accept(rb).(model.RegExp)
	quant := ctx.Quantifier()
	if quant != nil {
		container, ok := quant.Accept(rb).(model.RegQuantifier)
		if !ok {
			panic("parse error in quantifier")
		}
		container.Content = atom
		return container
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
	return model.RegString{Content: ctx.GetText()}
}

func (rb *RegexBuilder) VisitLetters(ctx *gen.LettersContext) interface{} {
	return model.RegString{Content: ctx.GetText()}
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

var memIndex = 0

func (rb *RegexBuilder) VisitLookaround(ctx *gen.LookaroundContext) interface{} {
	str := ctx.GetText()
	if str[:3] != "(?=" {
		panic("lookaround except lookahead is not supported in this regex engine.")
	}
	content := ctx.Alternation().Accept(rb).(model.RegExp)
	retval := model.RegPosLa{Content: content, MemIndex: memIndex}
	memIndex += 1
	return retval
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
	q := ctx.GetText()

	// ?, *, + を処理
	switch q {
	case "*":
		// RegStarでも良いけどこっちのほうが呼び出し元でキャストしやすい
		return model.RegQuantifier{
			Min: 0,
			Max: 0,
		}
	case "+":
		return model.RegQuantifier{
			Min: 1,
			Max: 0,
		}
	case "?":
		return model.RegQuantifier{
			Min: 0,
			Max: 1,
		}
	}

	// {n[,m]}を処理
	from, err := strconv.Atoi(ctx.GetFrom().GetText())
	if err != nil {
		panic("parse error in quantifier of form {n[,m]}")
	}
	// , を含まない場合
	if !strings.Contains(q, ",") {
		return model.RegQuantifier{
			Min: from,
			Max: from,
		}
	}
	// , を含む場合、mが含まれるかでさらに分岐
	rawTo := ctx.GetTo()
	if rawTo != nil { // m含む
		to, err := strconv.Atoi(rawTo.GetText())
		if err != nil {
			panic("parse error in quantifier of form {n[,m]}")
		}
		return model.RegQuantifier{
			Min: from,
			Max: to,
		}
	} else { // m含まない
		return model.RegQuantifier{
			Min: from,
			Max: 0, // 0にセットする=上限なしと同義
		}
	}
}

func (rb *RegexBuilder) VisitQuoting(ctx *gen.QuotingContext) interface{} {
	txt := ctx.GetText()
	if txt == "\\." ||
		txt == "\\^" ||
		txt == "\\$" ||
		txt == "\\*" ||
		txt == "\\+" ||
		txt == "\\?" ||
		txt == "\\(" ||
		txt == "\\)" ||
		txt == "\\[" ||
		txt == "\\{" ||
		txt == "\\\\" ||
		txt == "\\|" {
		return model.RegString{Content: string(txt[1])}
	}
	panic("unexpected input in quoting: " + txt)
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
