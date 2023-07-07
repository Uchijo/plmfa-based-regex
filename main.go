package main

import (
	"fmt"

	"github.com/Uchijo/plmfa-based-regex/eval"
	"github.com/Uchijo/plmfa-based-regex/model"
)

func main() {
	regex :=
		model.RegApp{
			Contents: []model.RegExp{
				model.RegPosLa{
					Content: model.RegUnion{
						Left: model.RegString{
							Content: "g",
						},
						Right: model.RegString{
							Content: "h",
						},
					},
					MemIndex: 1,
				},
				model.RegUnion{
					Left: model.RegString{
						Content: "h",
					},
					Right: model.RegString{
						Content: "g",
					},
				},
			},
		}
	states, start, _ := model.CreateCompleteStates(regex)
	input := eval.InputBuffer{
		Input: "g",
	}

	for _, v := range states.States() {
		fmt.Printf("%+v\n", v)
	}

	matched := search(states, input, start, eval.PosMemoryList{})
	fmt.Printf("input: %v, match: %v", input.Input, matched)
}

// search returns true if plmfa accepts input
func search(st model.StateList, input eval.InputBuffer, currentId string, posMem eval.PosMemoryList) bool {
	fmt.Printf("buffer:%v , memory: %+v\n", input.Input, posMem)
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
			hasGoal := search(st, input, v.MoveTo, posMem)
			if hasGoal {
				return true
			}

		case model.Consumption:
			if input.CanConsume(v.Input) {
				consumed, _ := input.Consumed(v.Input)
				appended := posMem.Appended(v.Input)
				hasGoal := search(st, consumed, v.MoveTo, appended)
				if hasGoal {
					return true
				}
			}

		case model.PosMem:
			if v.PLInst.Inst == model.Open {
				opened := posMem.OpenedMem(v.PLInst.MemIndex)
				hasGoal := search(st, input, v.MoveTo, opened)
				if hasGoal {
					return true
				}
			} else if v.PLInst.Inst == model.Close {
				closed, memContent := posMem.ClosedMem(v.PLInst.MemIndex)
				appended, _ := input.Appended(memContent)
				hasGoal := search(st, appended, v.MoveTo, closed)
				if hasGoal {
					return true
				}
			}

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
