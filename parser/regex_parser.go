// Code generated from regexParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // regexParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type regexParser struct {
	*antlr.BaseParser
}

var RegexParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func regexparserParserInit() {
	staticData := &RegexParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'|'", "'+'", "'?'", "'*'", "'.'", "", "'{'", "",
		"", "", "", "", "", "", "", "','", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "']'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "LPAREN", "RPAREN", "PIPE", "PLUS", "QUESTION", "STAR", "WildcardEsc",
		"Char", "StartQuantity", "SingleCharEsc", "MultiCharEsc", "CatEsc",
		"ComplEsc", "NegCharGroup", "PosCharGroup", "EndQuantity", "QuantExact",
		"COMMA", "EndCategory", "IsCategory", "Letters", "Marks", "Numbers",
		"Punctuation", "Separators", "Symbols", "Others", "IsBlock", "NestedSingleCharEsc",
		"NestedMultiCharEsc", "NestedCatEsc", "NestedComplEsc", "NestedNegCharGroup",
		"NestedPosCharGroup", "EndCharGroup", "DASH", "XmlChar",
	}
	staticData.RuleNames = []string{
		"root", "regExp", "branch", "piece", "quantifier", "quantity", "quantRange",
		"quantMin", "atom", "charClass", "charClassExpr", "charGroup", "posCharGroup",
		"charRange", "seRange", "charOrEsc", "charClassEsc", "catEsc", "complEsc",
		"charProp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 154, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1, 1, 2, 5, 2,
		53, 8, 2, 10, 2, 12, 2, 56, 9, 2, 1, 3, 1, 3, 3, 3, 60, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 69, 8, 4, 1, 5, 1, 5, 1, 5, 3, 5,
		74, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 89, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 94, 8, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 3, 11, 101, 8, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 112, 8, 11, 1, 11, 3, 11,
		115, 8, 11, 1, 12, 3, 12, 118, 8, 12, 1, 12, 1, 12, 4, 12, 122, 8, 12,
		11, 12, 12, 12, 123, 1, 13, 1, 13, 3, 13, 128, 8, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16,
		142, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 0, 0, 20, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 0, 5, 2, 0, 14, 15, 33, 34, 2, 0, 10, 10,
		37, 37, 2, 0, 12, 12, 31, 31, 2, 0, 13, 13, 32, 32, 2, 0, 20, 20, 28, 28,
		159, 0, 40, 1, 0, 0, 0, 2, 43, 1, 0, 0, 0, 4, 54, 1, 0, 0, 0, 6, 57, 1,
		0, 0, 0, 8, 68, 1, 0, 0, 0, 10, 73, 1, 0, 0, 0, 12, 75, 1, 0, 0, 0, 14,
		79, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18, 93, 1, 0, 0, 0, 20, 95, 1, 0, 0,
		0, 22, 114, 1, 0, 0, 0, 24, 117, 1, 0, 0, 0, 26, 127, 1, 0, 0, 0, 28, 129,
		1, 0, 0, 0, 30, 133, 1, 0, 0, 0, 32, 141, 1, 0, 0, 0, 34, 143, 1, 0, 0,
		0, 36, 147, 1, 0, 0, 0, 38, 151, 1, 0, 0, 0, 40, 41, 3, 2, 1, 0, 41, 42,
		5, 0, 0, 1, 42, 1, 1, 0, 0, 0, 43, 48, 3, 4, 2, 0, 44, 45, 5, 3, 0, 0,
		45, 47, 3, 4, 2, 0, 46, 44, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 3, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51,
		53, 3, 6, 3, 0, 52, 51, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0,
		0, 54, 55, 1, 0, 0, 0, 55, 5, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 59, 3,
		16, 8, 0, 58, 60, 3, 8, 4, 0, 59, 58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60,
		7, 1, 0, 0, 0, 61, 69, 5, 5, 0, 0, 62, 69, 5, 6, 0, 0, 63, 69, 5, 4, 0,
		0, 64, 65, 5, 9, 0, 0, 65, 66, 3, 10, 5, 0, 66, 67, 5, 16, 0, 0, 67, 69,
		1, 0, 0, 0, 68, 61, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0, 68, 63, 1, 0, 0, 0,
		68, 64, 1, 0, 0, 0, 69, 9, 1, 0, 0, 0, 70, 74, 3, 12, 6, 0, 71, 74, 3,
		14, 7, 0, 72, 74, 5, 17, 0, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		73, 72, 1, 0, 0, 0, 74, 11, 1, 0, 0, 0, 75, 76, 5, 17, 0, 0, 76, 77, 5,
		18, 0, 0, 77, 78, 5, 17, 0, 0, 78, 13, 1, 0, 0, 0, 79, 80, 5, 17, 0, 0,
		80, 81, 5, 18, 0, 0, 81, 15, 1, 0, 0, 0, 82, 89, 5, 8, 0, 0, 83, 89, 3,
		18, 9, 0, 84, 85, 5, 1, 0, 0, 85, 86, 3, 2, 1, 0, 86, 87, 5, 2, 0, 0, 87,
		89, 1, 0, 0, 0, 88, 82, 1, 0, 0, 0, 88, 83, 1, 0, 0, 0, 88, 84, 1, 0, 0,
		0, 89, 17, 1, 0, 0, 0, 90, 94, 3, 32, 16, 0, 91, 94, 3, 20, 10, 0, 92,
		94, 5, 7, 0, 0, 93, 90, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0,
		0, 94, 19, 1, 0, 0, 0, 95, 96, 7, 0, 0, 0, 96, 97, 3, 22, 11, 0, 97, 98,
		5, 35, 0, 0, 98, 21, 1, 0, 0, 0, 99, 101, 3, 24, 12, 0, 100, 99, 1, 0,
		0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 5, 36, 0, 0,
		103, 104, 5, 36, 0, 0, 104, 115, 3, 20, 10, 0, 105, 106, 3, 24, 12, 0,
		106, 107, 5, 36, 0, 0, 107, 108, 3, 20, 10, 0, 108, 115, 1, 0, 0, 0, 109,
		111, 3, 24, 12, 0, 110, 112, 5, 36, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112,
		1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 115, 5, 36, 0, 0, 114, 100, 1, 0,
		0, 0, 114, 105, 1, 0, 0, 0, 114, 109, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0,
		115, 23, 1, 0, 0, 0, 116, 118, 5, 36, 0, 0, 117, 116, 1, 0, 0, 0, 117,
		118, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 122, 3, 26, 13, 0, 120, 122,
		3, 32, 16, 0, 121, 119, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1,
		0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 25, 1, 0, 0,
		0, 125, 128, 3, 28, 14, 0, 126, 128, 5, 37, 0, 0, 127, 125, 1, 0, 0, 0,
		127, 126, 1, 0, 0, 0, 128, 27, 1, 0, 0, 0, 129, 130, 3, 30, 15, 0, 130,
		131, 5, 36, 0, 0, 131, 132, 3, 30, 15, 0, 132, 29, 1, 0, 0, 0, 133, 134,
		7, 1, 0, 0, 134, 31, 1, 0, 0, 0, 135, 142, 5, 10, 0, 0, 136, 142, 5, 29,
		0, 0, 137, 142, 5, 11, 0, 0, 138, 142, 5, 30, 0, 0, 139, 142, 3, 34, 17,
		0, 140, 142, 3, 36, 18, 0, 141, 135, 1, 0, 0, 0, 141, 136, 1, 0, 0, 0,
		141, 137, 1, 0, 0, 0, 141, 138, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141,
		140, 1, 0, 0, 0, 142, 33, 1, 0, 0, 0, 143, 144, 7, 2, 0, 0, 144, 145, 3,
		38, 19, 0, 145, 146, 5, 19, 0, 0, 146, 35, 1, 0, 0, 0, 147, 148, 7, 3,
		0, 0, 148, 149, 3, 38, 19, 0, 149, 150, 5, 19, 0, 0, 150, 37, 1, 0, 0,
		0, 151, 152, 7, 4, 0, 0, 152, 39, 1, 0, 0, 0, 15, 48, 54, 59, 68, 73, 88,
		93, 100, 111, 114, 117, 121, 123, 127, 141,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// regexParserInit initializes any static state used to implement regexParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewregexParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RegexParserInit() {
	staticData := &RegexParserParserStaticData
	staticData.once.Do(regexparserParserInit)
}

// NewregexParser produces a new parser instance for the optional input antlr.TokenStream.
func NewregexParser(input antlr.TokenStream) *regexParser {
	RegexParserInit()
	this := new(regexParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RegexParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "regexParser.g4"

	return this
}

// regexParser tokens.
const (
	regexParserEOF                 = antlr.TokenEOF
	regexParserLPAREN              = 1
	regexParserRPAREN              = 2
	regexParserPIPE                = 3
	regexParserPLUS                = 4
	regexParserQUESTION            = 5
	regexParserSTAR                = 6
	regexParserWildcardEsc         = 7
	regexParserChar                = 8
	regexParserStartQuantity       = 9
	regexParserSingleCharEsc       = 10
	regexParserMultiCharEsc        = 11
	regexParserCatEsc              = 12
	regexParserComplEsc            = 13
	regexParserNegCharGroup        = 14
	regexParserPosCharGroup        = 15
	regexParserEndQuantity         = 16
	regexParserQuantExact          = 17
	regexParserCOMMA               = 18
	regexParserEndCategory         = 19
	regexParserIsCategory          = 20
	regexParserLetters             = 21
	regexParserMarks               = 22
	regexParserNumbers             = 23
	regexParserPunctuation         = 24
	regexParserSeparators          = 25
	regexParserSymbols             = 26
	regexParserOthers              = 27
	regexParserIsBlock             = 28
	regexParserNestedSingleCharEsc = 29
	regexParserNestedMultiCharEsc  = 30
	regexParserNestedCatEsc        = 31
	regexParserNestedComplEsc      = 32
	regexParserNestedNegCharGroup  = 33
	regexParserNestedPosCharGroup  = 34
	regexParserEndCharGroup        = 35
	regexParserDASH                = 36
	regexParserXmlChar             = 37
)

// regexParser rules.
const (
	regexParserRULE_root          = 0
	regexParserRULE_regExp        = 1
	regexParserRULE_branch        = 2
	regexParserRULE_piece         = 3
	regexParserRULE_quantifier    = 4
	regexParserRULE_quantity      = 5
	regexParserRULE_quantRange    = 6
	regexParserRULE_quantMin      = 7
	regexParserRULE_atom          = 8
	regexParserRULE_charClass     = 9
	regexParserRULE_charClassExpr = 10
	regexParserRULE_charGroup     = 11
	regexParserRULE_posCharGroup  = 12
	regexParserRULE_charRange     = 13
	regexParserRULE_seRange       = 14
	regexParserRULE_charOrEsc     = 15
	regexParserRULE_charClassEsc  = 16
	regexParserRULE_catEsc        = 17
	regexParserRULE_complEsc      = 18
	regexParserRULE_charProp      = 19
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RegExp() IRegExpContext
	EOF() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) RegExp() IRegExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegExpContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(regexParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, regexParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.RegExp()
	}
	{
		p.SetState(41)
		p.Match(regexParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegExpContext is an interface to support dynamic dispatch.
type IRegExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBranch() []IBranchContext
	Branch(i int) IBranchContext
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode

	// IsRegExpContext differentiates from other interfaces.
	IsRegExpContext()
}

type RegExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegExpContext() *RegExpContext {
	var p = new(RegExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_regExp
	return p
}

func InitEmptyRegExpContext(p *RegExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_regExp
}

func (*RegExpContext) IsRegExpContext() {}

func NewRegExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegExpContext {
	var p = new(RegExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_regExp

	return p
}

func (s *RegExpContext) GetParser() antlr.Parser { return s.parser }

func (s *RegExpContext) AllBranch() []IBranchContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBranchContext); ok {
			len++
		}
	}

	tst := make([]IBranchContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBranchContext); ok {
			tst[i] = t.(IBranchContext)
			i++
		}
	}

	return tst
}

func (s *RegExpContext) Branch(i int) IBranchContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBranchContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *RegExpContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(regexParserPIPE)
}

func (s *RegExpContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(regexParserPIPE, i)
}

func (s *RegExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterRegExp(s)
	}
}

func (s *RegExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitRegExp(s)
	}
}

func (s *RegExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitRegExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) RegExp() (localctx IRegExpContext) {
	localctx = NewRegExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, regexParserRULE_regExp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Branch()
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == regexParserPIPE {
		{
			p.SetState(44)
			p.Match(regexParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Branch()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPiece() []IPieceContext
	Piece(i int) IPieceContext

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_branch
	return p
}

func InitEmptyBranchContext(p *BranchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_branch
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) AllPiece() []IPieceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPieceContext); ok {
			len++
		}
	}

	tst := make([]IPieceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPieceContext); ok {
			tst[i] = t.(IPieceContext)
			i++
		}
	}

	return tst
}

func (s *BranchContext) Piece(i int) IPieceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPieceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPieceContext)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterBranch(s)
	}
}

func (s *BranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitBranch(s)
	}
}

func (s *BranchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitBranch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) Branch() (localctx IBranchContext) {
	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, regexParserRULE_branch)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33822932354) != 0 {
		{
			p.SetState(51)
			p.Piece()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPieceContext is an interface to support dynamic dispatch.
type IPieceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Atom() IAtomContext
	Quantifier() IQuantifierContext

	// IsPieceContext differentiates from other interfaces.
	IsPieceContext()
}

type PieceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPieceContext() *PieceContext {
	var p = new(PieceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_piece
	return p
}

func InitEmptyPieceContext(p *PieceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_piece
}

func (*PieceContext) IsPieceContext() {}

func NewPieceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PieceContext {
	var p = new(PieceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_piece

	return p
}

func (s *PieceContext) GetParser() antlr.Parser { return s.parser }

func (s *PieceContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *PieceContext) Quantifier() IQuantifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuantifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuantifierContext)
}

func (s *PieceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PieceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PieceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterPiece(s)
	}
}

func (s *PieceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitPiece(s)
	}
}

func (s *PieceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitPiece(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) Piece() (localctx IPieceContext) {
	localctx = NewPieceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, regexParserRULE_piece)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Atom()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&624) != 0 {
		{
			p.SetState(58)
			p.Quantifier()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuantifierContext is an interface to support dynamic dispatch.
type IQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUESTION() antlr.TerminalNode
	STAR() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	StartQuantity() antlr.TerminalNode
	Quantity() IQuantityContext
	EndQuantity() antlr.TerminalNode

	// IsQuantifierContext differentiates from other interfaces.
	IsQuantifierContext()
}

type QuantifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantifierContext() *QuantifierContext {
	var p = new(QuantifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantifier
	return p
}

func InitEmptyQuantifierContext(p *QuantifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantifier
}

func (*QuantifierContext) IsQuantifierContext() {}

func NewQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantifierContext {
	var p = new(QuantifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantifier

	return p
}

func (s *QuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantifierContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(regexParserQUESTION, 0)
}

func (s *QuantifierContext) STAR() antlr.TerminalNode {
	return s.GetToken(regexParserSTAR, 0)
}

func (s *QuantifierContext) PLUS() antlr.TerminalNode {
	return s.GetToken(regexParserPLUS, 0)
}

func (s *QuantifierContext) StartQuantity() antlr.TerminalNode {
	return s.GetToken(regexParserStartQuantity, 0)
}

func (s *QuantifierContext) Quantity() IQuantityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuantityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuantityContext)
}

func (s *QuantifierContext) EndQuantity() antlr.TerminalNode {
	return s.GetToken(regexParserEndQuantity, 0)
}

func (s *QuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantifier(s)
	}
}

func (s *QuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantifier(s)
	}
}

func (s *QuantifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitQuantifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) Quantifier() (localctx IQuantifierContext) {
	localctx = NewQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, regexParserRULE_quantifier)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case regexParserQUESTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(regexParserQUESTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(regexParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserPLUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Match(regexParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserStartQuantity:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Match(regexParserStartQuantity)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Quantity()
		}
		{
			p.SetState(66)
			p.Match(regexParserEndQuantity)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuantityContext is an interface to support dynamic dispatch.
type IQuantityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QuantRange() IQuantRangeContext
	QuantMin() IQuantMinContext
	QuantExact() antlr.TerminalNode

	// IsQuantityContext differentiates from other interfaces.
	IsQuantityContext()
}

type QuantityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantityContext() *QuantityContext {
	var p = new(QuantityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantity
	return p
}

func InitEmptyQuantityContext(p *QuantityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantity
}

func (*QuantityContext) IsQuantityContext() {}

func NewQuantityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantityContext {
	var p = new(QuantityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantity

	return p
}

func (s *QuantityContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantityContext) QuantRange() IQuantRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuantRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuantRangeContext)
}

func (s *QuantityContext) QuantMin() IQuantMinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuantMinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuantMinContext)
}

func (s *QuantityContext) QuantExact() antlr.TerminalNode {
	return s.GetToken(regexParserQuantExact, 0)
}

func (s *QuantityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantity(s)
	}
}

func (s *QuantityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantity(s)
	}
}

func (s *QuantityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitQuantity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) Quantity() (localctx IQuantityContext) {
	localctx = NewQuantityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, regexParserRULE_quantity)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.QuantRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.QuantMin()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Match(regexParserQuantExact)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuantRangeContext is an interface to support dynamic dispatch.
type IQuantRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQuantExact() []antlr.TerminalNode
	QuantExact(i int) antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsQuantRangeContext differentiates from other interfaces.
	IsQuantRangeContext()
}

type QuantRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantRangeContext() *QuantRangeContext {
	var p = new(QuantRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantRange
	return p
}

func InitEmptyQuantRangeContext(p *QuantRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantRange
}

func (*QuantRangeContext) IsQuantRangeContext() {}

func NewQuantRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantRangeContext {
	var p = new(QuantRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantRange

	return p
}

func (s *QuantRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantRangeContext) AllQuantExact() []antlr.TerminalNode {
	return s.GetTokens(regexParserQuantExact)
}

func (s *QuantRangeContext) QuantExact(i int) antlr.TerminalNode {
	return s.GetToken(regexParserQuantExact, i)
}

func (s *QuantRangeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(regexParserCOMMA, 0)
}

func (s *QuantRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantRange(s)
	}
}

func (s *QuantRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantRange(s)
	}
}

func (s *QuantRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitQuantRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) QuantRange() (localctx IQuantRangeContext) {
	localctx = NewQuantRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, regexParserRULE_quantRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(regexParserQuantExact)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(regexParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(regexParserQuantExact)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuantMinContext is an interface to support dynamic dispatch.
type IQuantMinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QuantExact() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsQuantMinContext differentiates from other interfaces.
	IsQuantMinContext()
}

type QuantMinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantMinContext() *QuantMinContext {
	var p = new(QuantMinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantMin
	return p
}

func InitEmptyQuantMinContext(p *QuantMinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_quantMin
}

func (*QuantMinContext) IsQuantMinContext() {}

func NewQuantMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantMinContext {
	var p = new(QuantMinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_quantMin

	return p
}

func (s *QuantMinContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantMinContext) QuantExact() antlr.TerminalNode {
	return s.GetToken(regexParserQuantExact, 0)
}

func (s *QuantMinContext) COMMA() antlr.TerminalNode {
	return s.GetToken(regexParserCOMMA, 0)
}

func (s *QuantMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantMinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterQuantMin(s)
	}
}

func (s *QuantMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitQuantMin(s)
	}
}

func (s *QuantMinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitQuantMin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) QuantMin() (localctx IQuantMinContext) {
	localctx = NewQuantMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, regexParserRULE_quantMin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(regexParserQuantExact)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(regexParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Char() antlr.TerminalNode
	CharClass() ICharClassContext
	LPAREN() antlr.TerminalNode
	RegExp() IRegExpContext
	RPAREN() antlr.TerminalNode

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Char() antlr.TerminalNode {
	return s.GetToken(regexParserChar, 0)
}

func (s *AtomContext) CharClass() ICharClassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharClassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharClassContext)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(regexParserLPAREN, 0)
}

func (s *AtomContext) RegExp() IRegExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegExpContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(regexParserRPAREN, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, regexParserRULE_atom)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case regexParserChar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(regexParserChar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserWildcardEsc, regexParserSingleCharEsc, regexParserMultiCharEsc, regexParserCatEsc, regexParserComplEsc, regexParserNegCharGroup, regexParserPosCharGroup, regexParserNestedSingleCharEsc, regexParserNestedMultiCharEsc, regexParserNestedCatEsc, regexParserNestedComplEsc, regexParserNestedNegCharGroup, regexParserNestedPosCharGroup:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.CharClass()
		}

	case regexParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(regexParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.RegExp()
		}
		{
			p.SetState(86)
			p.Match(regexParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharClassContext is an interface to support dynamic dispatch.
type ICharClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CharClassEsc() ICharClassEscContext
	CharClassExpr() ICharClassExprContext
	WildcardEsc() antlr.TerminalNode

	// IsCharClassContext differentiates from other interfaces.
	IsCharClassContext()
}

type CharClassContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharClassContext() *CharClassContext {
	var p = new(CharClassContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charClass
	return p
}

func InitEmptyCharClassContext(p *CharClassContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charClass
}

func (*CharClassContext) IsCharClassContext() {}

func NewCharClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharClassContext {
	var p = new(CharClassContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charClass

	return p
}

func (s *CharClassContext) GetParser() antlr.Parser { return s.parser }

func (s *CharClassContext) CharClassEsc() ICharClassEscContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharClassEscContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharClassEscContext)
}

func (s *CharClassContext) CharClassExpr() ICharClassExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharClassExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharClassExprContext)
}

func (s *CharClassContext) WildcardEsc() antlr.TerminalNode {
	return s.GetToken(regexParserWildcardEsc, 0)
}

func (s *CharClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharClass(s)
	}
}

func (s *CharClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharClass(s)
	}
}

func (s *CharClassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCharClass(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CharClass() (localctx ICharClassContext) {
	localctx = NewCharClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, regexParserRULE_charClass)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case regexParserSingleCharEsc, regexParserMultiCharEsc, regexParserCatEsc, regexParserComplEsc, regexParserNestedSingleCharEsc, regexParserNestedMultiCharEsc, regexParserNestedCatEsc, regexParserNestedComplEsc:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.CharClassEsc()
		}

	case regexParserNegCharGroup, regexParserPosCharGroup, regexParserNestedNegCharGroup, regexParserNestedPosCharGroup:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.CharClassExpr()
		}

	case regexParserWildcardEsc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(regexParserWildcardEsc)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharClassExprContext is an interface to support dynamic dispatch.
type ICharClassExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CharGroup() ICharGroupContext
	EndCharGroup() antlr.TerminalNode
	NegCharGroup() antlr.TerminalNode
	NestedNegCharGroup() antlr.TerminalNode
	PosCharGroup() antlr.TerminalNode
	NestedPosCharGroup() antlr.TerminalNode

	// IsCharClassExprContext differentiates from other interfaces.
	IsCharClassExprContext()
}

type CharClassExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharClassExprContext() *CharClassExprContext {
	var p = new(CharClassExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charClassExpr
	return p
}

func InitEmptyCharClassExprContext(p *CharClassExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charClassExpr
}

func (*CharClassExprContext) IsCharClassExprContext() {}

func NewCharClassExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharClassExprContext {
	var p = new(CharClassExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charClassExpr

	return p
}

func (s *CharClassExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CharClassExprContext) CharGroup() ICharGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharGroupContext)
}

func (s *CharClassExprContext) EndCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserEndCharGroup, 0)
}

func (s *CharClassExprContext) NegCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserNegCharGroup, 0)
}

func (s *CharClassExprContext) NestedNegCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserNestedNegCharGroup, 0)
}

func (s *CharClassExprContext) PosCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserPosCharGroup, 0)
}

func (s *CharClassExprContext) NestedPosCharGroup() antlr.TerminalNode {
	return s.GetToken(regexParserNestedPosCharGroup, 0)
}

func (s *CharClassExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharClassExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharClassExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharClassExpr(s)
	}
}

func (s *CharClassExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharClassExpr(s)
	}
}

func (s *CharClassExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCharClassExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CharClassExpr() (localctx ICharClassExprContext) {
	localctx = NewCharClassExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, regexParserRULE_charClassExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25769852928) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(96)
		p.CharGroup()
	}
	{
		p.SetState(97)
		p.Match(regexParserEndCharGroup)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharGroupContext is an interface to support dynamic dispatch.
type ICharGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDASH() []antlr.TerminalNode
	DASH(i int) antlr.TerminalNode
	CharClassExpr() ICharClassExprContext
	PosCharGroup() IPosCharGroupContext

	// IsCharGroupContext differentiates from other interfaces.
	IsCharGroupContext()
}

type CharGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharGroupContext() *CharGroupContext {
	var p = new(CharGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charGroup
	return p
}

func InitEmptyCharGroupContext(p *CharGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charGroup
}

func (*CharGroupContext) IsCharGroupContext() {}

func NewCharGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharGroupContext {
	var p = new(CharGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charGroup

	return p
}

func (s *CharGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *CharGroupContext) AllDASH() []antlr.TerminalNode {
	return s.GetTokens(regexParserDASH)
}

func (s *CharGroupContext) DASH(i int) antlr.TerminalNode {
	return s.GetToken(regexParserDASH, i)
}

func (s *CharGroupContext) CharClassExpr() ICharClassExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharClassExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharClassExprContext)
}

func (s *CharGroupContext) PosCharGroup() IPosCharGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPosCharGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPosCharGroupContext)
}

func (s *CharGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharGroup(s)
	}
}

func (s *CharGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharGroup(s)
	}
}

func (s *CharGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCharGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CharGroup() (localctx ICharGroupContext) {
	localctx = NewCharGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, regexParserRULE_charGroup)
	var _la int

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(100)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(99)
				p.PosCharGroup()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(102)
			p.Match(regexParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(regexParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.CharClassExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.PosCharGroup()
		}
		{
			p.SetState(106)
			p.Match(regexParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.CharClassExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.PosCharGroup()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == regexParserDASH {
			{
				p.SetState(110)
				p.Match(regexParserDASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Match(regexParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPosCharGroupContext is an interface to support dynamic dispatch.
type IPosCharGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DASH() antlr.TerminalNode
	AllCharRange() []ICharRangeContext
	CharRange(i int) ICharRangeContext
	AllCharClassEsc() []ICharClassEscContext
	CharClassEsc(i int) ICharClassEscContext

	// IsPosCharGroupContext differentiates from other interfaces.
	IsPosCharGroupContext()
}

type PosCharGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPosCharGroupContext() *PosCharGroupContext {
	var p = new(PosCharGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_posCharGroup
	return p
}

func InitEmptyPosCharGroupContext(p *PosCharGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_posCharGroup
}

func (*PosCharGroupContext) IsPosCharGroupContext() {}

func NewPosCharGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PosCharGroupContext {
	var p = new(PosCharGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_posCharGroup

	return p
}

func (s *PosCharGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *PosCharGroupContext) DASH() antlr.TerminalNode {
	return s.GetToken(regexParserDASH, 0)
}

func (s *PosCharGroupContext) AllCharRange() []ICharRangeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharRangeContext); ok {
			len++
		}
	}

	tst := make([]ICharRangeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharRangeContext); ok {
			tst[i] = t.(ICharRangeContext)
			i++
		}
	}

	return tst
}

func (s *PosCharGroupContext) CharRange(i int) ICharRangeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharRangeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharRangeContext)
}

func (s *PosCharGroupContext) AllCharClassEsc() []ICharClassEscContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharClassEscContext); ok {
			len++
		}
	}

	tst := make([]ICharClassEscContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharClassEscContext); ok {
			tst[i] = t.(ICharClassEscContext)
			i++
		}
	}

	return tst
}

func (s *PosCharGroupContext) CharClassEsc(i int) ICharClassEscContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharClassEscContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharClassEscContext)
}

func (s *PosCharGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PosCharGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PosCharGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterPosCharGroup(s)
	}
}

func (s *PosCharGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitPosCharGroup(s)
	}
}

func (s *PosCharGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitPosCharGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) PosCharGroup() (localctx IPosCharGroupContext) {
	localctx = NewPosCharGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, regexParserRULE_posCharGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == regexParserDASH {
		{
			p.SetState(116)
			p.Match(regexParserDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&145492032512) != 0) {
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(119)
				p.CharRange()
			}

		case 2:
			{
				p.SetState(120)
				p.CharClassEsc()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharRangeContext is an interface to support dynamic dispatch.
type ICharRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SeRange() ISeRangeContext
	XmlChar() antlr.TerminalNode

	// IsCharRangeContext differentiates from other interfaces.
	IsCharRangeContext()
}

type CharRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharRangeContext() *CharRangeContext {
	var p = new(CharRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charRange
	return p
}

func InitEmptyCharRangeContext(p *CharRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charRange
}

func (*CharRangeContext) IsCharRangeContext() {}

func NewCharRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharRangeContext {
	var p = new(CharRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charRange

	return p
}

func (s *CharRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *CharRangeContext) SeRange() ISeRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeRangeContext)
}

func (s *CharRangeContext) XmlChar() antlr.TerminalNode {
	return s.GetToken(regexParserXmlChar, 0)
}

func (s *CharRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharRange(s)
	}
}

func (s *CharRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharRange(s)
	}
}

func (s *CharRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCharRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CharRange() (localctx ICharRangeContext) {
	localctx = NewCharRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, regexParserRULE_charRange)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.SeRange()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
			p.Match(regexParserXmlChar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISeRangeContext is an interface to support dynamic dispatch.
type ISeRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCharOrEsc() []ICharOrEscContext
	CharOrEsc(i int) ICharOrEscContext
	DASH() antlr.TerminalNode

	// IsSeRangeContext differentiates from other interfaces.
	IsSeRangeContext()
}

type SeRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeRangeContext() *SeRangeContext {
	var p = new(SeRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_seRange
	return p
}

func InitEmptySeRangeContext(p *SeRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_seRange
}

func (*SeRangeContext) IsSeRangeContext() {}

func NewSeRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeRangeContext {
	var p = new(SeRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_seRange

	return p
}

func (s *SeRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *SeRangeContext) AllCharOrEsc() []ICharOrEscContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICharOrEscContext); ok {
			len++
		}
	}

	tst := make([]ICharOrEscContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICharOrEscContext); ok {
			tst[i] = t.(ICharOrEscContext)
			i++
		}
	}

	return tst
}

func (s *SeRangeContext) CharOrEsc(i int) ICharOrEscContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharOrEscContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharOrEscContext)
}

func (s *SeRangeContext) DASH() antlr.TerminalNode {
	return s.GetToken(regexParserDASH, 0)
}

func (s *SeRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterSeRange(s)
	}
}

func (s *SeRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitSeRange(s)
	}
}

func (s *SeRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitSeRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) SeRange() (localctx ISeRangeContext) {
	localctx = NewSeRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, regexParserRULE_seRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.CharOrEsc()
	}
	{
		p.SetState(130)
		p.Match(regexParserDASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.CharOrEsc()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharOrEscContext is an interface to support dynamic dispatch.
type ICharOrEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	XmlChar() antlr.TerminalNode
	SingleCharEsc() antlr.TerminalNode

	// IsCharOrEscContext differentiates from other interfaces.
	IsCharOrEscContext()
}

type CharOrEscContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharOrEscContext() *CharOrEscContext {
	var p = new(CharOrEscContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charOrEsc
	return p
}

func InitEmptyCharOrEscContext(p *CharOrEscContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charOrEsc
}

func (*CharOrEscContext) IsCharOrEscContext() {}

func NewCharOrEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharOrEscContext {
	var p = new(CharOrEscContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charOrEsc

	return p
}

func (s *CharOrEscContext) GetParser() antlr.Parser { return s.parser }

func (s *CharOrEscContext) XmlChar() antlr.TerminalNode {
	return s.GetToken(regexParserXmlChar, 0)
}

func (s *CharOrEscContext) SingleCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserSingleCharEsc, 0)
}

func (s *CharOrEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharOrEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharOrEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharOrEsc(s)
	}
}

func (s *CharOrEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharOrEsc(s)
	}
}

func (s *CharOrEscContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCharOrEsc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CharOrEsc() (localctx ICharOrEscContext) {
	localctx = NewCharOrEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, regexParserRULE_charOrEsc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserSingleCharEsc || _la == regexParserXmlChar) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharClassEscContext is an interface to support dynamic dispatch.
type ICharClassEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleCharEsc() antlr.TerminalNode
	NestedSingleCharEsc() antlr.TerminalNode
	MultiCharEsc() antlr.TerminalNode
	NestedMultiCharEsc() antlr.TerminalNode
	CatEsc() ICatEscContext
	ComplEsc() IComplEscContext

	// IsCharClassEscContext differentiates from other interfaces.
	IsCharClassEscContext()
}

type CharClassEscContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharClassEscContext() *CharClassEscContext {
	var p = new(CharClassEscContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charClassEsc
	return p
}

func InitEmptyCharClassEscContext(p *CharClassEscContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charClassEsc
}

func (*CharClassEscContext) IsCharClassEscContext() {}

func NewCharClassEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharClassEscContext {
	var p = new(CharClassEscContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charClassEsc

	return p
}

func (s *CharClassEscContext) GetParser() antlr.Parser { return s.parser }

func (s *CharClassEscContext) SingleCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserSingleCharEsc, 0)
}

func (s *CharClassEscContext) NestedSingleCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedSingleCharEsc, 0)
}

func (s *CharClassEscContext) MultiCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserMultiCharEsc, 0)
}

func (s *CharClassEscContext) NestedMultiCharEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedMultiCharEsc, 0)
}

func (s *CharClassEscContext) CatEsc() ICatEscContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatEscContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatEscContext)
}

func (s *CharClassEscContext) ComplEsc() IComplEscContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComplEscContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComplEscContext)
}

func (s *CharClassEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharClassEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharClassEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharClassEsc(s)
	}
}

func (s *CharClassEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharClassEsc(s)
	}
}

func (s *CharClassEscContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCharClassEsc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CharClassEsc() (localctx ICharClassEscContext) {
	localctx = NewCharClassEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, regexParserRULE_charClassEsc)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case regexParserSingleCharEsc:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(regexParserSingleCharEsc)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserNestedSingleCharEsc:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Match(regexParserNestedSingleCharEsc)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserMultiCharEsc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Match(regexParserMultiCharEsc)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserNestedMultiCharEsc:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Match(regexParserNestedMultiCharEsc)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case regexParserCatEsc, regexParserNestedCatEsc:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.CatEsc()
		}

	case regexParserComplEsc, regexParserNestedComplEsc:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.ComplEsc()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatEscContext is an interface to support dynamic dispatch.
type ICatEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CharProp() ICharPropContext
	EndCategory() antlr.TerminalNode
	CatEsc() antlr.TerminalNode
	NestedCatEsc() antlr.TerminalNode

	// IsCatEscContext differentiates from other interfaces.
	IsCatEscContext()
}

type CatEscContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatEscContext() *CatEscContext {
	var p = new(CatEscContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_catEsc
	return p
}

func InitEmptyCatEscContext(p *CatEscContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_catEsc
}

func (*CatEscContext) IsCatEscContext() {}

func NewCatEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatEscContext {
	var p = new(CatEscContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_catEsc

	return p
}

func (s *CatEscContext) GetParser() antlr.Parser { return s.parser }

func (s *CatEscContext) CharProp() ICharPropContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharPropContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharPropContext)
}

func (s *CatEscContext) EndCategory() antlr.TerminalNode {
	return s.GetToken(regexParserEndCategory, 0)
}

func (s *CatEscContext) CatEsc() antlr.TerminalNode {
	return s.GetToken(regexParserCatEsc, 0)
}

func (s *CatEscContext) NestedCatEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedCatEsc, 0)
}

func (s *CatEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCatEsc(s)
	}
}

func (s *CatEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCatEsc(s)
	}
}

func (s *CatEscContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCatEsc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CatEsc() (localctx ICatEscContext) {
	localctx = NewCatEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, regexParserRULE_catEsc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserCatEsc || _la == regexParserNestedCatEsc) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(144)
		p.CharProp()
	}
	{
		p.SetState(145)
		p.Match(regexParserEndCategory)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComplEscContext is an interface to support dynamic dispatch.
type IComplEscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CharProp() ICharPropContext
	EndCategory() antlr.TerminalNode
	ComplEsc() antlr.TerminalNode
	NestedComplEsc() antlr.TerminalNode

	// IsComplEscContext differentiates from other interfaces.
	IsComplEscContext()
}

type ComplEscContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComplEscContext() *ComplEscContext {
	var p = new(ComplEscContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_complEsc
	return p
}

func InitEmptyComplEscContext(p *ComplEscContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_complEsc
}

func (*ComplEscContext) IsComplEscContext() {}

func NewComplEscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComplEscContext {
	var p = new(ComplEscContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_complEsc

	return p
}

func (s *ComplEscContext) GetParser() antlr.Parser { return s.parser }

func (s *ComplEscContext) CharProp() ICharPropContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharPropContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharPropContext)
}

func (s *ComplEscContext) EndCategory() antlr.TerminalNode {
	return s.GetToken(regexParserEndCategory, 0)
}

func (s *ComplEscContext) ComplEsc() antlr.TerminalNode {
	return s.GetToken(regexParserComplEsc, 0)
}

func (s *ComplEscContext) NestedComplEsc() antlr.TerminalNode {
	return s.GetToken(regexParserNestedComplEsc, 0)
}

func (s *ComplEscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComplEscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComplEscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterComplEsc(s)
	}
}

func (s *ComplEscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitComplEsc(s)
	}
}

func (s *ComplEscContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitComplEsc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) ComplEsc() (localctx IComplEscContext) {
	localctx = NewComplEscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, regexParserRULE_complEsc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserComplEsc || _la == regexParserNestedComplEsc) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(148)
		p.CharProp()
	}
	{
		p.SetState(149)
		p.Match(regexParserEndCategory)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharPropContext is an interface to support dynamic dispatch.
type ICharPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IsCategory() antlr.TerminalNode
	IsBlock() antlr.TerminalNode

	// IsCharPropContext differentiates from other interfaces.
	IsCharPropContext()
}

type CharPropContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharPropContext() *CharPropContext {
	var p = new(CharPropContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charProp
	return p
}

func InitEmptyCharPropContext(p *CharPropContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = regexParserRULE_charProp
}

func (*CharPropContext) IsCharPropContext() {}

func NewCharPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharPropContext {
	var p = new(CharPropContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = regexParserRULE_charProp

	return p
}

func (s *CharPropContext) GetParser() antlr.Parser { return s.parser }

func (s *CharPropContext) IsCategory() antlr.TerminalNode {
	return s.GetToken(regexParserIsCategory, 0)
}

func (s *CharPropContext) IsBlock() antlr.TerminalNode {
	return s.GetToken(regexParserIsBlock, 0)
}

func (s *CharPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharPropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.EnterCharProp(s)
	}
}

func (s *CharPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(regexParserListener); ok {
		listenerT.ExitCharProp(s)
	}
}

func (s *CharPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case regexParserVisitor:
		return t.VisitCharProp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *regexParser) CharProp() (localctx ICharPropContext) {
	localctx = NewCharPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, regexParserRULE_charProp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(_la == regexParserIsCategory || _la == regexParserIsBlock) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
