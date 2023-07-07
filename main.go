package main

import (
	"errors"
	"fmt"

	"github.com/Uchijo/plmfa-based-regex/model"
)

func main() {
	fmt.Println("hello, world")

	hoge := InputBuffer{Input: "hoge"}
	fmt.Printf("can consume ho: %v\n", hoge.CanConsume("ho"))
	hoge.Consume("ho")
	fmt.Printf("input: %v", hoge.Input)
}

// search returns true if plmfa accepts input
func search(st model.StateList, input, currentId string) bool {
	// fixme: エラー処理直す
	curState, _ := st.StateById(currentId)

	// 終状態でない & 行先がない -> マッチ失敗
	if !curState.IsEnd && len(curState.Moves) == 0 {
		return false
	}

	// 終状態 & 入力文字がない -> マッチ成功
	if curState.IsEnd && len(input) == 0 {
		return true
	}

	// 入力文字がない & イプシロンでたどり着けるゴールがある -> マッチ
	goalReachable := searchEps(st, currentId)
	if goalReachable && len(input) == 0 {
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
			if len(input) >= 1 && input[:1] == v.Input {
				hasGoal := search(st, input[1:], v.MoveTo)
				if hasGoal {
					return true
				}
			}

		case model.PosMem:
		case model.CapMem:
		case model.Ref:
			// hasGoal := search(st, )
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

func (ib InputBuffer) CanConsume(matcher string) bool {
	prefixLen := len(matcher)
	inputLen := len(ib.Input)
	fmt.Printf("pl: %v, il: %v\n", prefixLen, inputLen)
	if inputLen < prefixLen {
		return false
	}
	prefix := ib.Input[prefixLen:]
	fmt.Printf("prefix: %v, matcher: %v\n", prefix, matcher)
	return prefix == matcher
}

func (ib *InputBuffer) Consume(matcher string) error {
	if !ib.CanConsume(matcher) {
		return errors.New("cannot consume.")
	}
	ib.Input = ib.Input[:len(matcher)]
	return nil
}
