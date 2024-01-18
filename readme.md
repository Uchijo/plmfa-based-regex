# PLMFAベースの正規表現の実装

[On Lookaheads in Regular Expressions with Backreferences](https://www.jstage.jst.go.jp/article/transinf/E106.D/5/E106.D_2022EDP7098/_article/-char/ja/)にて提案されている、古典正規表現に肯定先読みと後方参照を追加したオートマトンであるplmfaを用いて正規表現のマッチを行うプログラムです。

## 使い方

正規表現と入力を与え、文字列全体がマッチするか確認する

```shell-session
$ go run main.go input pattern
```

PCRE2におけるテストケースをパースし、json形式で保存し直す

```shell-session
$ cd cmd/parse_testsuits/

# main.go内で対象とする入力ファイル名を書き換える
$ vim main.go

$ go run main.go > output.json
```

json形式で保存されているテストを実施する

```shell-session
$ cd cmd/try_tests/

$ go run main.go path/to/input.json > result.json
```

## 参考等

- 文法ファイル
  - [https://github.com/antlr/grammars-v4/blob/master/pcre/PCRE.g4](https://github.com/antlr/grammars-v4/blob/master/pcre/PCRE.g4)
