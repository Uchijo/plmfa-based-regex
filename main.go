package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/uchijo/plmfa-based-regex/eval"
	"github.com/uchijo/plmfa-based-regex/model"
	"github.com/uchijo/plmfa-based-regex/parser"
	gen "github.com/uchijo/plmfa-based-regex/parser/gen"
)

type Output struct {
	Match bool   `json:"match"`
	Regex string `json:"regex"`
	Input string `json:"input"`
}

type Args struct {
	input string
	regex string
}

func main() {
	args := Args{
		input: os.Args[1],
		regex: os.Args[2],
	}

	is := antlr.NewInputStream(args.regex)
	lexer := gen.NewPCRELexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := gen.NewPCREParser(stream)
	p.BuildParseTrees = true
	tree := p.Pcre()
	regex := tree.Accept(&parser.RegexBuilder{}).(model.RegExp)

	states, start, _ := model.CreateCompleteStates(regex)
	input := eval.InputBuffer{
		Input: args.input,
	}

	matched := search(states, input, start, eval.PosMemoryList{}, eval.CapMemoryList{}, 0)
	result := Output{
		Match: matched,
		Input: args.input,
		Regex: args.regex,
	}
	byteText, _ := json.Marshal(result)
	fmt.Println(string(byteText))
}

var logs = []eval.Log{}

func isInLoop(log eval.Log) bool {
	for _, v := range logs {
		if v.Alike(log) {
			return true
		}
	}
	return false
}

// search returns true if plmfa accepts input
func search(
	st model.StateList,
	input eval.InputBuffer,
	currentId string,
	posMem eval.PosMemoryList,
	capMem eval.CapMemoryList,
	depth int,
) bool {
	// 無限ループ検知
	currentLog := eval.Log{
		Input:          input.Input,
		CurrentId:      currentId,
		PositiveMemory: posMem,
		CaptureMemory:  capMem,
	}
	if isInLoop(currentLog) {
		return false
	}
	logs = append(logs, currentLog)

	for i := 0; i < depth; i++ {
		// fmt.Printf("  ")
	}
	// fmt.Printf("buffer: %v, state: %v, pos_memory: %+v, cap_memory: %+v\n", input.Input, currentId, posMem, capMem)
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
			hasGoal := search(st, input, v.MoveTo, posMem, capMem, depth+1)
			if hasGoal {
				return true
			}

		case model.Consumption:
			if input.CanConsume(v.Input) {
				consumed, _ := input.Consumed(v.Input)
				appendedPos := posMem.Appended(v.Input)
				appendedCap := capMem.Appended(v.Input)
				hasGoal := search(st, consumed, v.MoveTo, appendedPos, appendedCap, depth+1)
				if hasGoal {
					return true
				}
			}

		case model.PosMem:
			if v.PLInst.Inst == model.Open {
				opened := posMem.OpenedMem(v.PLInst.MemIndex)
				hasGoal := search(st, input, v.MoveTo, opened, capMem, depth+1)
				if hasGoal {
					return true
				}
			} else if v.PLInst.Inst == model.Close {
				closed, memContent := posMem.ClosedMem(v.PLInst.MemIndex)
				appended, _ := input.Appended(memContent)
				hasGoal := search(st, appended, v.MoveTo, closed, capMem, depth+1)
				if hasGoal {
					return true
				}
			}

		case model.CapMem:
			if v.CInst.Inst == model.Open {
				opened := capMem.OpenedMem(v.CInst.MemIndex)
				hasGoal := search(st, input, v.MoveTo, posMem, opened, depth+1)
				if hasGoal {
					return true
				}
			} else if v.CInst.Inst == model.Close {
				closed := capMem.ClosedMem(v.CInst.MemIndex)
				hasGoal := search(st, input, v.MoveTo, posMem, closed, depth+1)
				if hasGoal {
					return true
				}
			}

		case model.Ref:
			mem := capMem.Content(v.RefIndex)
			if input.CanConsume(mem) {
				consumed, _ := input.Consumed(mem)
				appendedPos := posMem.Appended(v.Input)
				appendedCap := capMem.Appended(v.Input)
				hasGoal := search(st, consumed, v.MoveTo, appendedPos, appendedCap, depth+1)
				if hasGoal {
					return true
				}
			}

		case model.ArbitraryConsumption:
			if input.Len() >= 1 {
				toConsume := input.Input[:1]
				consumed, _ := input.Consumed(toConsume)
				appendedPos := posMem.Appended(toConsume)
				appendedCap := capMem.Appended(toConsume)
				hasGoal := search(st, consumed, v.MoveTo, appendedPos, appendedCap, depth+1)
				if hasGoal {
					return true
				}
			}
		}
	}

	for i := 0; i < depth; i++ {
		// fmt.Printf("  ")
	}
	// fmt.Println("backtrack!")
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
