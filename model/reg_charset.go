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

func (cr CharRange) Contains(c rune) bool {
	return cr.Start <= c && c <= cr.End
}

func (cr CharRange) Len() int { return 1 }

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

func (cl CharList) Len() int { return 1 }
