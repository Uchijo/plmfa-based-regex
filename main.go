package main

import (
	"errors"
	"fmt"

	"github.com/Uchijo/plmfa-based-regex/model"
)

func main() {
	// regex := model.RegApp{
	// 	Contents: []model.RegExp{
	// 		model.RegPosLa{
	// 			Content: model.RegString{
	// 				Content: "h",
	// 			},
	// 			MemIndex: 1,
	// 		},
	// 		model.RegString{
	// 			Content: "h",
	// 		},
	// 	},
	// }
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
	input := InputBuffer{
		Input: "g",
	}

	for _, v := range states.States() {
		fmt.Printf("%+v\n", v)
	}

	matched := search(states, input, start, PosMemoryList{})
	fmt.Printf("input: %v, match: %v", input.Input, matched)
}

// search returns true if plmfa accepts input
func search(st model.StateList, input InputBuffer, currentId string, posMem PosMemoryList) bool {
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
			} else {
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

func (ib InputBuffer) Consumed(matcher string) (InputBuffer, error) {
	if !ib.CanConsume(matcher) {
		return InputBuffer{}, errors.New("cannot consume")
	}
	ib.Input = ib.Input[len(matcher):]
	return ib, nil
}

func (ib InputBuffer) Appended(prefix string) (InputBuffer, error) {
	ib.Input = prefix + ib.Input
	return ib, nil
}

type PosMemoryList []PosMemory

// Append appends given suffix to opened memory
// it doesn't change closed memory
func (pml PosMemoryList) Appended(suffix string) PosMemoryList {
	for i, v := range pml {
		pml[i] = v.appended(suffix)
	}
	return pml
}

// ChangeStatus changes status of specified memory
// if memory not found, it creates memory that satisfies given status
func (pml PosMemoryList) OpenedMem(index int) PosMemoryList {
	for i, v := range pml {
		if v.Index == index {
			pml[i].IsOpen = true
			return pml
		}
	}
	pml = append(
		pml,
		PosMemory{
			Index:   index,
			IsOpen:  true,
			Content: "",
		},
	)
	return pml
}

// CloseMem closes specified memory, clean ups that memory and return content of memory
func (pml PosMemoryList) ClosedMem(index int) (PosMemoryList, string) {
	for i, v := range pml {
		if v.Index == index {
			content := v.Content
			pml[i].Content = ""
			pml[i].IsOpen = false

			return pml, content
		}
	}
	return pml, ""
}

type PosMemory struct {
	Index   int
	Content string
	IsOpen  bool
}

// append appends given suffix to memory if memory is opened
func (pm PosMemory) appended(suffix string) PosMemory {
	if pm.IsOpen {
		pm.Content = pm.Content + suffix
	}
	return pm
}
