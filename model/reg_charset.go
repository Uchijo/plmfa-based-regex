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
	// IsIn tests if [c] is in the specified character set or not.
	Contains(c rune) bool
}

// [a-z]タイプの文字集合
type CharRange struct {
	Start rune
	End   rune

	// ^が先頭にない -> true
	WhiteList bool
}

func (cr CharRange) Contains(c rune) bool {
	return cr.Start <= c && c <= cr.End
}

// [abc]タイプの文字集合
type CharList struct {
	Chars []rune

	// ^が先頭にない -> true
	WhiteList bool
}

func (cl CharList) Contains(c rune) bool {
	for _, val := range cl.Chars {
		if val == c {
			return true
		}
	}
	return false
}

// 1文字だけ、[]なしでも使えるやつ
type SingleChar struct {
	Char rune
}

func (sc SingleChar) Contains(c rune) bool {
	return sc.Char == c
}
