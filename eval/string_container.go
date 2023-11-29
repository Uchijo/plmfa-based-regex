package eval

import "github.com/uchijo/plmfa-based-regex/model"

// 型の辻褄を合わせるためだけの構造体
// model.CharContainerを実装する
type stringContainer string

var _ model.CharContainer = (*stringContainer)(nil)

func (cc stringContainer) Len() int {
	return len([]rune(cc))
}
func (cc stringContainer) WholeMatches(c string) bool {
	return string(cc) == c
}
