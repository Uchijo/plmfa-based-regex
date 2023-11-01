package model

import (
	"github.com/google/uuid"
)

// 文字集合を表す構造体
type RegCharSet struct {
	// 対象文字
	Content CharContainer
}

func (rc RegCharSet) States(startId string) ([]State, string, error) {
	nextId, err := uuid.NewRandom()
	if err != nil {
		return []State{}, "", err
	}
	return []State{
		{
			Id:    startId,
			IsEnd: false,
			Moves: []Move{
				{
					MType:  Consumption,
					Input:  rc.Content,
					MoveTo: nextId.String(),
				},
			},
		},
	}, nextId.String(), nil
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

// [a-zABC]タイプに対応する用。 CharRange, CharList等の複合
type CompositeContainer struct {
	Contents  []CharContainer
	WhiteList bool
}

var _ CharContainer = (*CompositeContainer)(nil)

// fixme: ↓eval内にあるstringContainerとかと合わせて使うと死ぬ
func (cc CompositeContainer) Len() int { return 1 }
func (cc CompositeContainer) WholeMatches(c string) bool {
	// 全体としてホワイトリストなら、どれかにマッチするなら進めてok
	if cc.WhiteList {
		for _, v := range cc.Contents {
			if v.WholeMatches(c) {
				return true
			}
		}
		return false
	} else { // ブラックリスト方式なら、Contentsのうち一個でもアウト(=マッチする)なら全体としてアウト
		ok := true
		for _, v := range cc.Contents {
			// マッチしない -> ok
			ok = ok && !v.WholeMatches(c)
		}
		return ok
	}
}

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

	if cl.WhiteList {
		for _, val := range cl.Chars {
			if val == runed[0] {
				return true
			}
		}
		return false
	} else {
		for _, val := range cl.Chars {
			if val == runed[0] {
				return false
			}
		}
		return true
	}
}

func (cl CharList) Len() int { return 1 }
