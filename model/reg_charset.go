package model

// 文字集合を表す構造体
type RegCharSet struct {
	// 対象文字
	Content CharContainer
}

func (rc RegCharSet) States(startId string) ([]State, string, error) {
	return []State{}, "", nil
}

type CharContainer interface {
	// 渡されたcが完全に一致する
	WholeMatches(c string) bool

	// CharContainerに含まれる文字を具体化したときの長さ
	Len() int
}

// [a-z]タイプの文字集合
type CharRange struct {
	Start rune
	End   rune

	// ^が先頭にない -> true
	WhiteList bool
}

var _ CharContainer = (*CharRange)(nil)

func (cr CharRange) WholeMatches(c string) bool {
	r := []rune(c)

	// 文字集合自体は1文字としかマッチしない
	if len(r) != 1 {
		return false
	}
	if cr.WhiteList {
		return cr.Start <= r[0] && r[0] <= cr.End
	} else {
		return !(cr.Start <= r[0] && r[0] <= cr.End)
	}
}

func (cr CharRange) Len() int { return 1 }

// [abc]タイプの文字集合
type CharList struct {
	Chars []rune

	// ^が先頭にない -> true
	WhiteList bool
}

var _ CharContainer = (*CharList)(nil)

func (cl CharList) WholeMatches(c string) bool {
	runed := []rune(c)

	// 文字集合自体は1文字としかマッチしない
	if len(runed) != 1 {
		return false
	}

	for _, val := range cl.Chars {
		if val == runed[0] {
			return true
		}
	}
	return false
}

func (cl CharList) Len() int { return 1 }
