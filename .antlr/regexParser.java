// Generated from /home/yuta/repos/plmfa-based-regex/regexParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class regexParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		LPAREN=1, RPAREN=2, PIPE=3, PLUS=4, QUESTION=5, STAR=6, WildcardEsc=7, 
		Char=8, StartQuantity=9, SingleCharEsc=10, MultiCharEsc=11, CatEsc=12, 
		ComplEsc=13, NegCharGroup=14, PosCharGroup=15, EndQuantity=16, QuantExact=17, 
		COMMA=18, EndCategory=19, IsCategory=20, Letters=21, Marks=22, Numbers=23, 
		Punctuation=24, Separators=25, Symbols=26, Others=27, IsBlock=28, NestedSingleCharEsc=29, 
		NestedMultiCharEsc=30, NestedCatEsc=31, NestedComplEsc=32, NestedNegCharGroup=33, 
		NestedPosCharGroup=34, EndCharGroup=35, DASH=36, XmlChar=37;
	public static final int
		RULE_root = 0, RULE_regExp = 1, RULE_branch = 2, RULE_piece = 3, RULE_quantifier = 4, 
		RULE_quantity = 5, RULE_quantRange = 6, RULE_quantMin = 7, RULE_atom = 8, 
		RULE_charClass = 9, RULE_charClassExpr = 10, RULE_charGroup = 11, RULE_posCharGroup = 12, 
		RULE_charRange = 13, RULE_seRange = 14, RULE_charOrEsc = 15, RULE_charClassEsc = 16, 
		RULE_catEsc = 17, RULE_complEsc = 18, RULE_charProp = 19;
	private static String[] makeRuleNames() {
		return new String[] {
			"root", "regExp", "branch", "piece", "quantifier", "quantity", "quantRange", 
			"quantMin", "atom", "charClass", "charClassExpr", "charGroup", "posCharGroup", 
			"charRange", "seRange", "charOrEsc", "charClassEsc", "catEsc", "complEsc", 
			"charProp"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'|'", "'+'", "'?'", "'*'", "'.'", null, "'{'", null, 
			null, null, null, null, null, null, null, "','", null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"']'", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "LPAREN", "RPAREN", "PIPE", "PLUS", "QUESTION", "STAR", "WildcardEsc", 
			"Char", "StartQuantity", "SingleCharEsc", "MultiCharEsc", "CatEsc", "ComplEsc", 
			"NegCharGroup", "PosCharGroup", "EndQuantity", "QuantExact", "COMMA", 
			"EndCategory", "IsCategory", "Letters", "Marks", "Numbers", "Punctuation", 
			"Separators", "Symbols", "Others", "IsBlock", "NestedSingleCharEsc", 
			"NestedMultiCharEsc", "NestedCatEsc", "NestedComplEsc", "NestedNegCharGroup", 
			"NestedPosCharGroup", "EndCharGroup", "DASH", "XmlChar"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "regexParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public regexParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class RootContext extends ParserRuleContext {
		public RegExpContext regExp() {
			return getRuleContext(RegExpContext.class,0);
		}
		public TerminalNode EOF() { return getToken(regexParser.EOF, 0); }
		public RootContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_root; }
	}

	public final RootContext root() throws RecognitionException {
		RootContext _localctx = new RootContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_root);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			regExp();
			setState(41);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RegExpContext extends ParserRuleContext {
		public List<BranchContext> branch() {
			return getRuleContexts(BranchContext.class);
		}
		public BranchContext branch(int i) {
			return getRuleContext(BranchContext.class,i);
		}
		public List<TerminalNode> PIPE() { return getTokens(regexParser.PIPE); }
		public TerminalNode PIPE(int i) {
			return getToken(regexParser.PIPE, i);
		}
		public RegExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_regExp; }
	}

	public final RegExpContext regExp() throws RecognitionException {
		RegExpContext _localctx = new RegExpContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_regExp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			branch();
			setState(48);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PIPE) {
				{
				{
				setState(44);
				match(PIPE);
				setState(45);
				branch();
				}
				}
				setState(50);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BranchContext extends ParserRuleContext {
		public List<PieceContext> piece() {
			return getRuleContexts(PieceContext.class);
		}
		public PieceContext piece(int i) {
			return getRuleContext(PieceContext.class,i);
		}
		public BranchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_branch; }
	}

	public final BranchContext branch() throws RecognitionException {
		BranchContext _localctx = new BranchContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_branch);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LPAREN) | (1L << WildcardEsc) | (1L << Char) | (1L << SingleCharEsc) | (1L << MultiCharEsc) | (1L << CatEsc) | (1L << ComplEsc) | (1L << NegCharGroup) | (1L << PosCharGroup) | (1L << NestedSingleCharEsc) | (1L << NestedMultiCharEsc) | (1L << NestedCatEsc) | (1L << NestedComplEsc) | (1L << NestedNegCharGroup) | (1L << NestedPosCharGroup))) != 0)) {
				{
				{
				setState(51);
				piece();
				}
				}
				setState(56);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PieceContext extends ParserRuleContext {
		public AtomContext atom() {
			return getRuleContext(AtomContext.class,0);
		}
		public QuantifierContext quantifier() {
			return getRuleContext(QuantifierContext.class,0);
		}
		public PieceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_piece; }
	}

	public final PieceContext piece() throws RecognitionException {
		PieceContext _localctx = new PieceContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_piece);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			atom();
			setState(59);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PLUS) | (1L << QUESTION) | (1L << STAR) | (1L << StartQuantity))) != 0)) {
				{
				setState(58);
				quantifier();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QuantifierContext extends ParserRuleContext {
		public TerminalNode QUESTION() { return getToken(regexParser.QUESTION, 0); }
		public TerminalNode STAR() { return getToken(regexParser.STAR, 0); }
		public TerminalNode PLUS() { return getToken(regexParser.PLUS, 0); }
		public TerminalNode StartQuantity() { return getToken(regexParser.StartQuantity, 0); }
		public QuantityContext quantity() {
			return getRuleContext(QuantityContext.class,0);
		}
		public TerminalNode EndQuantity() { return getToken(regexParser.EndQuantity, 0); }
		public QuantifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_quantifier; }
	}

	public final QuantifierContext quantifier() throws RecognitionException {
		QuantifierContext _localctx = new QuantifierContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_quantifier);
		try {
			setState(68);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case QUESTION:
				enterOuterAlt(_localctx, 1);
				{
				setState(61);
				match(QUESTION);
				}
				break;
			case STAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(62);
				match(STAR);
				}
				break;
			case PLUS:
				enterOuterAlt(_localctx, 3);
				{
				setState(63);
				match(PLUS);
				}
				break;
			case StartQuantity:
				enterOuterAlt(_localctx, 4);
				{
				setState(64);
				match(StartQuantity);
				setState(65);
				quantity();
				setState(66);
				match(EndQuantity);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QuantityContext extends ParserRuleContext {
		public QuantRangeContext quantRange() {
			return getRuleContext(QuantRangeContext.class,0);
		}
		public QuantMinContext quantMin() {
			return getRuleContext(QuantMinContext.class,0);
		}
		public TerminalNode QuantExact() { return getToken(regexParser.QuantExact, 0); }
		public QuantityContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_quantity; }
	}

	public final QuantityContext quantity() throws RecognitionException {
		QuantityContext _localctx = new QuantityContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_quantity);
		try {
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(70);
				quantRange();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(71);
				quantMin();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(72);
				match(QuantExact);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QuantRangeContext extends ParserRuleContext {
		public List<TerminalNode> QuantExact() { return getTokens(regexParser.QuantExact); }
		public TerminalNode QuantExact(int i) {
			return getToken(regexParser.QuantExact, i);
		}
		public TerminalNode COMMA() { return getToken(regexParser.COMMA, 0); }
		public QuantRangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_quantRange; }
	}

	public final QuantRangeContext quantRange() throws RecognitionException {
		QuantRangeContext _localctx = new QuantRangeContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_quantRange);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			match(QuantExact);
			setState(76);
			match(COMMA);
			setState(77);
			match(QuantExact);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QuantMinContext extends ParserRuleContext {
		public TerminalNode QuantExact() { return getToken(regexParser.QuantExact, 0); }
		public TerminalNode COMMA() { return getToken(regexParser.COMMA, 0); }
		public QuantMinContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_quantMin; }
	}

	public final QuantMinContext quantMin() throws RecognitionException {
		QuantMinContext _localctx = new QuantMinContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_quantMin);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			match(QuantExact);
			setState(80);
			match(COMMA);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AtomContext extends ParserRuleContext {
		public TerminalNode Char() { return getToken(regexParser.Char, 0); }
		public CharClassContext charClass() {
			return getRuleContext(CharClassContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(regexParser.LPAREN, 0); }
		public RegExpContext regExp() {
			return getRuleContext(RegExpContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(regexParser.RPAREN, 0); }
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_atom);
		try {
			setState(88);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Char:
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				match(Char);
				}
				break;
			case WildcardEsc:
			case SingleCharEsc:
			case MultiCharEsc:
			case CatEsc:
			case ComplEsc:
			case NegCharGroup:
			case PosCharGroup:
			case NestedSingleCharEsc:
			case NestedMultiCharEsc:
			case NestedCatEsc:
			case NestedComplEsc:
			case NestedNegCharGroup:
			case NestedPosCharGroup:
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				charClass();
				}
				break;
			case LPAREN:
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(84);
				match(LPAREN);
				setState(85);
				regExp();
				setState(86);
				match(RPAREN);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CharClassContext extends ParserRuleContext {
		public CharClassEscContext charClassEsc() {
			return getRuleContext(CharClassEscContext.class,0);
		}
		public CharClassExprContext charClassExpr() {
			return getRuleContext(CharClassExprContext.class,0);
		}
		public TerminalNode WildcardEsc() { return getToken(regexParser.WildcardEsc, 0); }
		public CharClassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_charClass; }
	}

	public final CharClassContext charClass() throws RecognitionException {
		CharClassContext _localctx = new CharClassContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_charClass);
		try {
			setState(93);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SingleCharEsc:
			case MultiCharEsc:
			case CatEsc:
			case ComplEsc:
			case NestedSingleCharEsc:
			case NestedMultiCharEsc:
			case NestedCatEsc:
			case NestedComplEsc:
				enterOuterAlt(_localctx, 1);
				{
				setState(90);
				charClassEsc();
				}
				break;
			case NegCharGroup:
			case PosCharGroup:
			case NestedNegCharGroup:
			case NestedPosCharGroup:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				charClassExpr();
				}
				break;
			case WildcardEsc:
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				match(WildcardEsc);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CharClassExprContext extends ParserRuleContext {
		public CharGroupContext charGroup() {
			return getRuleContext(CharGroupContext.class,0);
		}
		public TerminalNode EndCharGroup() { return getToken(regexParser.EndCharGroup, 0); }
		public TerminalNode NegCharGroup() { return getToken(regexParser.NegCharGroup, 0); }
		public TerminalNode NestedNegCharGroup() { return getToken(regexParser.NestedNegCharGroup, 0); }
		public TerminalNode PosCharGroup() { return getToken(regexParser.PosCharGroup, 0); }
		public TerminalNode NestedPosCharGroup() { return getToken(regexParser.NestedPosCharGroup, 0); }
		public CharClassExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_charClassExpr; }
	}

	public final CharClassExprContext charClassExpr() throws RecognitionException {
		CharClassExprContext _localctx = new CharClassExprContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_charClassExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NegCharGroup) | (1L << PosCharGroup) | (1L << NestedNegCharGroup) | (1L << NestedPosCharGroup))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(96);
			charGroup();
			setState(97);
			match(EndCharGroup);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CharGroupContext extends ParserRuleContext {
		public List<TerminalNode> DASH() { return getTokens(regexParser.DASH); }
		public TerminalNode DASH(int i) {
			return getToken(regexParser.DASH, i);
		}
		public CharClassExprContext charClassExpr() {
			return getRuleContext(CharClassExprContext.class,0);
		}
		public PosCharGroupContext posCharGroup() {
			return getRuleContext(PosCharGroupContext.class,0);
		}
		public CharGroupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_charGroup; }
	}

	public final CharGroupContext charGroup() throws RecognitionException {
		CharGroupContext _localctx = new CharGroupContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_charGroup);
		int _la;
		try {
			setState(114);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(100);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
				case 1:
					{
					setState(99);
					posCharGroup();
					}
					break;
				}
				setState(102);
				match(DASH);
				setState(103);
				match(DASH);
				setState(104);
				charClassExpr();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(105);
				posCharGroup();
				setState(106);
				match(DASH);
				setState(107);
				charClassExpr();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(109);
				posCharGroup();
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DASH) {
					{
					setState(110);
					match(DASH);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(113);
				match(DASH);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PosCharGroupContext extends ParserRuleContext {
		public TerminalNode DASH() { return getToken(regexParser.DASH, 0); }
		public List<CharRangeContext> charRange() {
			return getRuleContexts(CharRangeContext.class);
		}
		public CharRangeContext charRange(int i) {
			return getRuleContext(CharRangeContext.class,i);
		}
		public List<CharClassEscContext> charClassEsc() {
			return getRuleContexts(CharClassEscContext.class);
		}
		public CharClassEscContext charClassEsc(int i) {
			return getRuleContext(CharClassEscContext.class,i);
		}
		public PosCharGroupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_posCharGroup; }
	}

	public final PosCharGroupContext posCharGroup() throws RecognitionException {
		PosCharGroupContext _localctx = new PosCharGroupContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_posCharGroup);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DASH) {
				{
				setState(116);
				match(DASH);
				}
			}

			setState(121); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				setState(121);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
				case 1:
					{
					setState(119);
					charRange();
					}
					break;
				case 2:
					{
					setState(120);
					charClassEsc();
					}
					break;
				}
				}
				setState(123); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << SingleCharEsc) | (1L << MultiCharEsc) | (1L << CatEsc) | (1L << ComplEsc) | (1L << NestedSingleCharEsc) | (1L << NestedMultiCharEsc) | (1L << NestedCatEsc) | (1L << NestedComplEsc) | (1L << XmlChar))) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CharRangeContext extends ParserRuleContext {
		public SeRangeContext seRange() {
			return getRuleContext(SeRangeContext.class,0);
		}
		public TerminalNode XmlChar() { return getToken(regexParser.XmlChar, 0); }
		public CharRangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_charRange; }
	}

	public final CharRangeContext charRange() throws RecognitionException {
		CharRangeContext _localctx = new CharRangeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_charRange);
		try {
			setState(127);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				seRange();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				match(XmlChar);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SeRangeContext extends ParserRuleContext {
		public List<CharOrEscContext> charOrEsc() {
			return getRuleContexts(CharOrEscContext.class);
		}
		public CharOrEscContext charOrEsc(int i) {
			return getRuleContext(CharOrEscContext.class,i);
		}
		public TerminalNode DASH() { return getToken(regexParser.DASH, 0); }
		public SeRangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_seRange; }
	}

	public final SeRangeContext seRange() throws RecognitionException {
		SeRangeContext _localctx = new SeRangeContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_seRange);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			charOrEsc();
			setState(130);
			match(DASH);
			setState(131);
			charOrEsc();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CharOrEscContext extends ParserRuleContext {
		public TerminalNode XmlChar() { return getToken(regexParser.XmlChar, 0); }
		public TerminalNode SingleCharEsc() { return getToken(regexParser.SingleCharEsc, 0); }
		public CharOrEscContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_charOrEsc; }
	}

	public final CharOrEscContext charOrEsc() throws RecognitionException {
		CharOrEscContext _localctx = new CharOrEscContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_charOrEsc);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			_la = _input.LA(1);
			if ( !(_la==SingleCharEsc || _la==XmlChar) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CharClassEscContext extends ParserRuleContext {
		public TerminalNode SingleCharEsc() { return getToken(regexParser.SingleCharEsc, 0); }
		public TerminalNode NestedSingleCharEsc() { return getToken(regexParser.NestedSingleCharEsc, 0); }
		public TerminalNode MultiCharEsc() { return getToken(regexParser.MultiCharEsc, 0); }
		public TerminalNode NestedMultiCharEsc() { return getToken(regexParser.NestedMultiCharEsc, 0); }
		public CatEscContext catEsc() {
			return getRuleContext(CatEscContext.class,0);
		}
		public ComplEscContext complEsc() {
			return getRuleContext(ComplEscContext.class,0);
		}
		public CharClassEscContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_charClassEsc; }
	}

	public final CharClassEscContext charClassEsc() throws RecognitionException {
		CharClassEscContext _localctx = new CharClassEscContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_charClassEsc);
		try {
			setState(141);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SingleCharEsc:
				enterOuterAlt(_localctx, 1);
				{
				setState(135);
				match(SingleCharEsc);
				}
				break;
			case NestedSingleCharEsc:
				enterOuterAlt(_localctx, 2);
				{
				setState(136);
				match(NestedSingleCharEsc);
				}
				break;
			case MultiCharEsc:
				enterOuterAlt(_localctx, 3);
				{
				setState(137);
				match(MultiCharEsc);
				}
				break;
			case NestedMultiCharEsc:
				enterOuterAlt(_localctx, 4);
				{
				setState(138);
				match(NestedMultiCharEsc);
				}
				break;
			case CatEsc:
			case NestedCatEsc:
				enterOuterAlt(_localctx, 5);
				{
				setState(139);
				catEsc();
				}
				break;
			case ComplEsc:
			case NestedComplEsc:
				enterOuterAlt(_localctx, 6);
				{
				setState(140);
				complEsc();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CatEscContext extends ParserRuleContext {
		public CharPropContext charProp() {
			return getRuleContext(CharPropContext.class,0);
		}
		public TerminalNode EndCategory() { return getToken(regexParser.EndCategory, 0); }
		public TerminalNode CatEsc() { return getToken(regexParser.CatEsc, 0); }
		public TerminalNode NestedCatEsc() { return getToken(regexParser.NestedCatEsc, 0); }
		public CatEscContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_catEsc; }
	}

	public final CatEscContext catEsc() throws RecognitionException {
		CatEscContext _localctx = new CatEscContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_catEsc);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			_la = _input.LA(1);
			if ( !(_la==CatEsc || _la==NestedCatEsc) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(144);
			charProp();
			setState(145);
			match(EndCategory);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ComplEscContext extends ParserRuleContext {
		public CharPropContext charProp() {
			return getRuleContext(CharPropContext.class,0);
		}
		public TerminalNode EndCategory() { return getToken(regexParser.EndCategory, 0); }
		public TerminalNode ComplEsc() { return getToken(regexParser.ComplEsc, 0); }
		public TerminalNode NestedComplEsc() { return getToken(regexParser.NestedComplEsc, 0); }
		public ComplEscContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_complEsc; }
	}

	public final ComplEscContext complEsc() throws RecognitionException {
		ComplEscContext _localctx = new ComplEscContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_complEsc);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			_la = _input.LA(1);
			if ( !(_la==ComplEsc || _la==NestedComplEsc) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(148);
			charProp();
			setState(149);
			match(EndCategory);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CharPropContext extends ParserRuleContext {
		public TerminalNode IsCategory() { return getToken(regexParser.IsCategory, 0); }
		public TerminalNode IsBlock() { return getToken(regexParser.IsBlock, 0); }
		public CharPropContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_charProp; }
	}

	public final CharPropContext charProp() throws RecognitionException {
		CharPropContext _localctx = new CharPropContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_charProp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			_la = _input.LA(1);
			if ( !(_la==IsCategory || _la==IsBlock) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\'\u009c\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\2\3\3\3\3\3\3\7\3\61\n\3\f\3"+
		"\16\3\64\13\3\3\4\7\4\67\n\4\f\4\16\4:\13\4\3\5\3\5\5\5>\n\5\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\5\6G\n\6\3\7\3\7\3\7\5\7L\n\7\3\b\3\b\3\b\3\b\3\t\3"+
		"\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\5\n[\n\n\3\13\3\13\3\13\5\13`\n\13\3\f"+
		"\3\f\3\f\3\f\3\r\5\rg\n\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\rr\n\r"+
		"\3\r\5\ru\n\r\3\16\5\16x\n\16\3\16\3\16\6\16|\n\16\r\16\16\16}\3\17\3"+
		"\17\5\17\u0082\n\17\3\20\3\20\3\20\3\20\3\21\3\21\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\5\22\u0090\n\22\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25"+
		"\3\25\3\25\2\2\26\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(\2\7\4\2"+
		"\20\21#$\4\2\f\f\'\'\4\2\16\16!!\4\2\17\17\"\"\4\2\26\26\36\36\2\u00a1"+
		"\2*\3\2\2\2\4-\3\2\2\2\68\3\2\2\2\b;\3\2\2\2\nF\3\2\2\2\fK\3\2\2\2\16"+
		"M\3\2\2\2\20Q\3\2\2\2\22Z\3\2\2\2\24_\3\2\2\2\26a\3\2\2\2\30t\3\2\2\2"+
		"\32w\3\2\2\2\34\u0081\3\2\2\2\36\u0083\3\2\2\2 \u0087\3\2\2\2\"\u008f"+
		"\3\2\2\2$\u0091\3\2\2\2&\u0095\3\2\2\2(\u0099\3\2\2\2*+\5\4\3\2+,\7\2"+
		"\2\3,\3\3\2\2\2-\62\5\6\4\2./\7\5\2\2/\61\5\6\4\2\60.\3\2\2\2\61\64\3"+
		"\2\2\2\62\60\3\2\2\2\62\63\3\2\2\2\63\5\3\2\2\2\64\62\3\2\2\2\65\67\5"+
		"\b\5\2\66\65\3\2\2\2\67:\3\2\2\28\66\3\2\2\289\3\2\2\29\7\3\2\2\2:8\3"+
		"\2\2\2;=\5\22\n\2<>\5\n\6\2=<\3\2\2\2=>\3\2\2\2>\t\3\2\2\2?G\7\7\2\2@"+
		"G\7\b\2\2AG\7\6\2\2BC\7\13\2\2CD\5\f\7\2DE\7\22\2\2EG\3\2\2\2F?\3\2\2"+
		"\2F@\3\2\2\2FA\3\2\2\2FB\3\2\2\2G\13\3\2\2\2HL\5\16\b\2IL\5\20\t\2JL\7"+
		"\23\2\2KH\3\2\2\2KI\3\2\2\2KJ\3\2\2\2L\r\3\2\2\2MN\7\23\2\2NO\7\24\2\2"+
		"OP\7\23\2\2P\17\3\2\2\2QR\7\23\2\2RS\7\24\2\2S\21\3\2\2\2T[\7\n\2\2U["+
		"\5\24\13\2VW\7\3\2\2WX\5\4\3\2XY\7\4\2\2Y[\3\2\2\2ZT\3\2\2\2ZU\3\2\2\2"+
		"ZV\3\2\2\2[\23\3\2\2\2\\`\5\"\22\2]`\5\26\f\2^`\7\t\2\2_\\\3\2\2\2_]\3"+
		"\2\2\2_^\3\2\2\2`\25\3\2\2\2ab\t\2\2\2bc\5\30\r\2cd\7%\2\2d\27\3\2\2\2"+
		"eg\5\32\16\2fe\3\2\2\2fg\3\2\2\2gh\3\2\2\2hi\7&\2\2ij\7&\2\2ju\5\26\f"+
		"\2kl\5\32\16\2lm\7&\2\2mn\5\26\f\2nu\3\2\2\2oq\5\32\16\2pr\7&\2\2qp\3"+
		"\2\2\2qr\3\2\2\2ru\3\2\2\2su\7&\2\2tf\3\2\2\2tk\3\2\2\2to\3\2\2\2ts\3"+
		"\2\2\2u\31\3\2\2\2vx\7&\2\2wv\3\2\2\2wx\3\2\2\2x{\3\2\2\2y|\5\34\17\2"+
		"z|\5\"\22\2{y\3\2\2\2{z\3\2\2\2|}\3\2\2\2}{\3\2\2\2}~\3\2\2\2~\33\3\2"+
		"\2\2\177\u0082\5\36\20\2\u0080\u0082\7\'\2\2\u0081\177\3\2\2\2\u0081\u0080"+
		"\3\2\2\2\u0082\35\3\2\2\2\u0083\u0084\5 \21\2\u0084\u0085\7&\2\2\u0085"+
		"\u0086\5 \21\2\u0086\37\3\2\2\2\u0087\u0088\t\3\2\2\u0088!\3\2\2\2\u0089"+
		"\u0090\7\f\2\2\u008a\u0090\7\37\2\2\u008b\u0090\7\r\2\2\u008c\u0090\7"+
		" \2\2\u008d\u0090\5$\23\2\u008e\u0090\5&\24\2\u008f\u0089\3\2\2\2\u008f"+
		"\u008a\3\2\2\2\u008f\u008b\3\2\2\2\u008f\u008c\3\2\2\2\u008f\u008d\3\2"+
		"\2\2\u008f\u008e\3\2\2\2\u0090#\3\2\2\2\u0091\u0092\t\4\2\2\u0092\u0093"+
		"\5(\25\2\u0093\u0094\7\25\2\2\u0094%\3\2\2\2\u0095\u0096\t\5\2\2\u0096"+
		"\u0097\5(\25\2\u0097\u0098\7\25\2\2\u0098\'\3\2\2\2\u0099\u009a\t\6\2"+
		"\2\u009a)\3\2\2\2\21\628=FKZ_fqtw{}\u0081\u008f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}