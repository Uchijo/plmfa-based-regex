# PLMFAベースの正規表現の実装

[On Lookaheads in Regular Expressions with Backreferences](https://www.jstage.jst.go.jp/article/transinf/E106.D/5/E106.D_2022EDP7098/_article/-char/ja/)にて提案されている、古典正規表現に肯定先読みと後方参照を追加したオートマトンであるplmfaを用いて正規表現のマッチを行うプログラムです。

## 使い方

```bash
go run main.go input pattern
```

## 参考等

- 文法ファイル
  - [https://github.com/antlr/grammars-v4/blob/master/pcre/PCRE.g4](https://github.com/antlr/grammars-v4/blob/master/pcre/PCRE.g4)
