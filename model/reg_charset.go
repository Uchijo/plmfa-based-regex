package model

type RegCharSet struct {
	// 先頭に ^ がついていない -> true
	WhiteList bool

	// 対象文字
	Content IsInner
}

func (rc RegCharSet) States(startId string) ([]State, string, error) {
	return []State{}, "", nil
}

type IsInner interface {
	// IsIn tests if [c] is in the specified character set or not.
	IsIn(c rune) bool
}

type CharRange struct {
	Start rune
	End   rune
}

func (cr CharRange) IsIn(c rune) bool {
	return cr.Start <= c && c <= cr.End
}

type CharList struct {
	Chars []rune
}

func (cl CharList) IsIn(c rune) bool {
	for _, val := range cl.Chars {
		if val == c {
			return true
		}
	}
	return false
}
