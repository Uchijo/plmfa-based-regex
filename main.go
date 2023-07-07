package main

import (
	"errors"
	"fmt"

	"github.com/Uchijo/plmfa-based-regex/model"
)

func main() {
	regex := model.RegStar{
		Content: model.RegString{
			Content: "hoge",
		},
	}
	states, start, _ := model.CreateCompleteStates(regex)
	input := InputBuffer{
		Input: "hoge",
	}
	matched := search(states, input, start)
	fmt.Printf("input: %v, match: %v", input.Input, matched)
}

// search returns true if plmfa accepts input
func search(st model.StateList, input InputBuffer, currentId string) bool {
	// fixme: エラー処理直す
	curState, _ := st.StateById(currentId)

	// 終状態でない & 行先がない -> マッチ失敗
	if !curState.IsEnd && len(curState.Moves) == 0 {
		return false
	}

	// 終状態 & 入力文字がない -> マッチ成功
	if curState.IsEnd && input.Len() == 0 {
		return true
	}

	// 入力文字がない & イプシロンでたどり着けるゴールがある -> マッチ
	goalReachable := searchEps(st, currentId)
	if goalReachable && input.Len() == 0 {
		return true
	}

	for _, v := range curState.Moves {
		switch v.MType {
		case model.Epsilon:
			hasGoal := search(st, input, v.MoveTo)
			if hasGoal {
				return true
			}

		case model.Consumption:
			if input.CanConsume(v.Input) {
				input.Consume(v.Input)
				hasGoal := search(st, input, v.MoveTo)
				if hasGoal {
					return true
				}
			}

		case model.PosMem:
		case model.CapMem:
		case model.Ref:
		}
	}

	return false
}

// searchEps returns true if goal state is reachable with only epsilon transitions
func searchEps(st model.StateList, currentId string) bool {
	curSt, _ := st.StateById(currentId)
	if curSt.IsEnd {
		return true
	}

	eps := curSt.ExtractMove(model.Epsilon)
	for _, v := range eps {
		hasGoal := searchEps(st, v.MoveTo)
		if hasGoal {
			return true
		}
	}

	return false
}

type InputBuffer struct {
	Input string
}

func (ib InputBuffer) Len() int {
	return len(ib.Input)
}

func (ib InputBuffer) CanConsume(matcher string) bool {
	prefixLen := len(matcher)
	inputLen := len(ib.Input)
	if inputLen < prefixLen {
		return false
	}
	prefix := ib.Input[:prefixLen]
	return prefix == matcher
}

func (ib *InputBuffer) Consume(matcher string) error {
	if !ib.CanConsume(matcher) {
		return errors.New("cannot consume")
	}
	ib.Input = ib.Input[len(matcher):]
	return nil
}
