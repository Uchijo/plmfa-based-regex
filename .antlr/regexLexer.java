// Generated from /home/yuta/repos/plmfa-based-regex/regexLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class regexLexer extends Lexer {
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
		QUANTITY=1, CATEGORY=2, CHARGROUP=3;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE", "QUANTITY", "CATEGORY", "CHARGROUP"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"LPAREN", "RPAREN", "PIPE", "PLUS", "QUESTION", "STAR", "WildcardEsc", 
			"Char", "StartQuantity", "SingleCharEsc", "MultiCharEsc", "CatEsc", "ComplEsc", 
			"NegCharGroup", "PosCharGroup", "EndQuantity", "QuantExact", "COMMA", 
			"EndCategory", "IsCategory", "Letters", "Marks", "Numbers", "Punctuation", 
			"Separators", "Symbols", "Others", "IsBlock", "NestedSingleCharEsc", 
			"NestedMultiCharEsc", "NestedCatEsc", "NestedComplEsc", "NestedNegCharGroup", 
			"NestedPosCharGroup", "EndCharGroup", "DASH", "XmlChar", "CAT_ESC", "COMPL_ESC", 
			"MULTI_ESC", "SINGLE_ESC"
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


	public regexLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "regexLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\'\u00e8\b\1\b\1\b"+
		"\1\b\1\4\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t"+
		"\4\n\t\n\4\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21"+
		"\t\21\4\22\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30"+
		"\t\30\4\31\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37"+
		"\t\37\4 \t \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)"+
		"\4*\t*\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t"+
		"\3\n\3\n\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\22"+
		"\6\22\u0087\n\22\r\22\16\22\u0088\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\5\25\u0098\n\25\3\26\3\26\5\26\u009c\n\26"+
		"\3\27\3\27\5\27\u00a0\n\27\3\30\3\30\5\30\u00a4\n\30\3\31\3\31\5\31\u00a8"+
		"\n\31\3\32\3\32\5\32\u00ac\n\32\3\33\3\33\5\33\u00b0\n\33\3\34\3\34\5"+
		"\34\u00b4\n\34\3\35\3\35\3\35\3\35\6\35\u00ba\n\35\r\35\16\35\u00bb\3"+
		"\36\3\36\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3"+
		"#\3#\3$\3$\3$\3$\3%\3%\3&\3&\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3)\3*\3"+
		"*\3*\2\2+\6\3\b\4\n\5\f\6\16\7\20\b\22\t\24\n\26\13\30\f\32\r\34\16\36"+
		"\17 \20\"\21$\22&\23(\24*\25,\26.\27\60\30\62\31\64\32\66\338\34:\35<"+
		"\36>\37@ B!D\"F#H$J%L&N\'P\2R\2T\2V\2\6\2\3\4\5\17\7\2*-\60\60AA]_~~\3"+
		"\2\62;\5\2noqqvw\5\2eeggpp\5\2ffnnqq\6\2ehkkqquu\5\2nnrruu\6\2eemmooq"+
		"q\5\2eehhpq\6\2//\62;C\\c|\5\2//]]__\n\2EFKKUUYYefkkuuyy\n\2*-/\60AA]"+
		"`ppttvv}\177\2\u00ef\2\6\3\2\2\2\2\b\3\2\2\2\2\n\3\2\2\2\2\f\3\2\2\2\2"+
		"\16\3\2\2\2\2\20\3\2\2\2\2\22\3\2\2\2\2\24\3\2\2\2\2\26\3\2\2\2\2\30\3"+
		"\2\2\2\2\32\3\2\2\2\2\34\3\2\2\2\2\36\3\2\2\2\2 \3\2\2\2\2\"\3\2\2\2\3"+
		"$\3\2\2\2\3&\3\2\2\2\3(\3\2\2\2\4*\3\2\2\2\4,\3\2\2\2\4.\3\2\2\2\4\60"+
		"\3\2\2\2\4\62\3\2\2\2\4\64\3\2\2\2\4\66\3\2\2\2\48\3\2\2\2\4:\3\2\2\2"+
		"\4<\3\2\2\2\5>\3\2\2\2\5@\3\2\2\2\5B\3\2\2\2\5D\3\2\2\2\5F\3\2\2\2\5H"+
		"\3\2\2\2\5J\3\2\2\2\5L\3\2\2\2\5N\3\2\2\2\6X\3\2\2\2\bZ\3\2\2\2\n\\\3"+
		"\2\2\2\f^\3\2\2\2\16`\3\2\2\2\20b\3\2\2\2\22d\3\2\2\2\24f\3\2\2\2\26h"+
		"\3\2\2\2\30l\3\2\2\2\32n\3\2\2\2\34p\3\2\2\2\36t\3\2\2\2 x\3\2\2\2\"}"+
		"\3\2\2\2$\u0081\3\2\2\2&\u0086\3\2\2\2(\u008a\3\2\2\2*\u008c\3\2\2\2,"+
		"\u0097\3\2\2\2.\u0099\3\2\2\2\60\u009d\3\2\2\2\62\u00a1\3\2\2\2\64\u00a5"+
		"\3\2\2\2\66\u00a9\3\2\2\28\u00ad\3\2\2\2:\u00b1\3\2\2\2<\u00b5\3\2\2\2"+
		">\u00bd\3\2\2\2@\u00bf\3\2\2\2B\u00c1\3\2\2\2D\u00c5\3\2\2\2F\u00c9\3"+
		"\2\2\2H\u00ce\3\2\2\2J\u00d2\3\2\2\2L\u00d6\3\2\2\2N\u00d8\3\2\2\2P\u00da"+
		"\3\2\2\2R\u00de\3\2\2\2T\u00e2\3\2\2\2V\u00e5\3\2\2\2XY\7*\2\2Y\7\3\2"+
		"\2\2Z[\7+\2\2[\t\3\2\2\2\\]\7~\2\2]\13\3\2\2\2^_\7-\2\2_\r\3\2\2\2`a\7"+
		"A\2\2a\17\3\2\2\2bc\7,\2\2c\21\3\2\2\2de\7\60\2\2e\23\3\2\2\2fg\n\2\2"+
		"\2g\25\3\2\2\2hi\7}\2\2ij\3\2\2\2jk\b\n\2\2k\27\3\2\2\2lm\5V*\2m\31\3"+
		"\2\2\2no\5T)\2o\33\3\2\2\2pq\5P\'\2qr\3\2\2\2rs\b\r\3\2s\35\3\2\2\2tu"+
		"\5R(\2uv\3\2\2\2vw\b\16\3\2w\37\3\2\2\2xy\7]\2\2yz\7`\2\2z{\3\2\2\2{|"+
		"\b\17\4\2|!\3\2\2\2}~\7]\2\2~\177\3\2\2\2\177\u0080\b\20\4\2\u0080#\3"+
		"\2\2\2\u0081\u0082\7\177\2\2\u0082\u0083\3\2\2\2\u0083\u0084\b\21\5\2"+
		"\u0084%\3\2\2\2\u0085\u0087\t\3\2\2\u0086\u0085\3\2\2\2\u0087\u0088\3"+
		"\2\2\2\u0088\u0086\3\2\2\2\u0088\u0089\3\2\2\2\u0089\'\3\2\2\2\u008a\u008b"+
		"\7.\2\2\u008b)\3\2\2\2\u008c\u008d\7\177\2\2\u008d\u008e\3\2\2\2\u008e"+
		"\u008f\b\24\5\2\u008f+\3\2\2\2\u0090\u0098\5.\26\2\u0091\u0098\5\60\27"+
		"\2\u0092\u0098\5\62\30\2\u0093\u0098\5\64\31\2\u0094\u0098\5\66\32\2\u0095"+
		"\u0098\58\33\2\u0096\u0098\5:\34\2\u0097\u0090\3\2\2\2\u0097\u0091\3\2"+
		"\2\2\u0097\u0092\3\2\2\2\u0097\u0093\3\2\2\2\u0097\u0094\3\2\2\2\u0097"+
		"\u0095\3\2\2\2\u0097\u0096\3\2\2\2\u0098-\3\2\2\2\u0099\u009b\7N\2\2\u009a"+
		"\u009c\t\4\2\2\u009b\u009a\3\2\2\2\u009b\u009c\3\2\2\2\u009c/\3\2\2\2"+
		"\u009d\u009f\7O\2\2\u009e\u00a0\t\5\2\2\u009f\u009e\3\2\2\2\u009f\u00a0"+
		"\3\2\2\2\u00a0\61\3\2\2\2\u00a1\u00a3\7P\2\2\u00a2\u00a4\t\6\2\2\u00a3"+
		"\u00a2\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4\63\3\2\2\2\u00a5\u00a7\7R\2\2"+
		"\u00a6\u00a8\t\7\2\2\u00a7\u00a6\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\65"+
		"\3\2\2\2\u00a9\u00ab\7\\\2\2\u00aa\u00ac\t\b\2\2\u00ab\u00aa\3\2\2\2\u00ab"+
		"\u00ac\3\2\2\2\u00ac\67\3\2\2\2\u00ad\u00af\7U\2\2\u00ae\u00b0\t\t\2\2"+
		"\u00af\u00ae\3\2\2\2\u00af\u00b0\3\2\2\2\u00b09\3\2\2\2\u00b1\u00b3\7"+
		"E\2\2\u00b2\u00b4\t\n\2\2\u00b3\u00b2\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4"+
		";\3\2\2\2\u00b5\u00b6\7K\2\2\u00b6\u00b7\7u\2\2\u00b7\u00b9\3\2\2\2\u00b8"+
		"\u00ba\t\13\2\2\u00b9\u00b8\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb\u00b9\3"+
		"\2\2\2\u00bb\u00bc\3\2\2\2\u00bc=\3\2\2\2\u00bd\u00be\5V*\2\u00be?\3\2"+
		"\2\2\u00bf\u00c0\5T)\2\u00c0A\3\2\2\2\u00c1\u00c2\5P\'\2\u00c2\u00c3\3"+
		"\2\2\2\u00c3\u00c4\b \3\2\u00c4C\3\2\2\2\u00c5\u00c6\5R(\2\u00c6\u00c7"+
		"\3\2\2\2\u00c7\u00c8\b!\3\2\u00c8E\3\2\2\2\u00c9\u00ca\7]\2\2\u00ca\u00cb"+
		"\7`\2\2\u00cb\u00cc\3\2\2\2\u00cc\u00cd\b\"\4\2\u00cdG\3\2\2\2\u00ce\u00cf"+
		"\7]\2\2\u00cf\u00d0\3\2\2\2\u00d0\u00d1\b#\4\2\u00d1I\3\2\2\2\u00d2\u00d3"+
		"\7_\2\2\u00d3\u00d4\3\2\2\2\u00d4\u00d5\b$\5\2\u00d5K\3\2\2\2\u00d6\u00d7"+
		"\7/\2\2\u00d7M\3\2\2\2\u00d8\u00d9\n\f\2\2\u00d9O\3\2\2\2\u00da\u00db"+
		"\7^\2\2\u00db\u00dc\7r\2\2\u00dc\u00dd\7}\2\2\u00ddQ\3\2\2\2\u00de\u00df"+
		"\7^\2\2\u00df\u00e0\7R\2\2\u00e0\u00e1\7}\2\2\u00e1S\3\2\2\2\u00e2\u00e3"+
		"\7^\2\2\u00e3\u00e4\t\r\2\2\u00e4U\3\2\2\2\u00e5\u00e6\7^\2\2\u00e6\u00e7"+
		"\t\16\2\2\u00e7W\3\2\2\2\21\2\3\4\5\u0088\u0097\u009b\u009f\u00a3\u00a7"+
		"\u00ab\u00af\u00b3\u00b9\u00bb\6\7\3\2\7\4\2\7\5\2\6\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}