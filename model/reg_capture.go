package model

import "github.com/google/uuid"

type RegCapture struct {
	RegExp
	// captureする正規表現
	Content RegExp
	// メモリのindex
	MemoryIndex int
}

func (rc RegCapture) States(startId string) ([]State, string, error) {
	// fixme: エラー処理直す
	contentSId, _ := uuid.NewRandom()
	entryState := State{
		Id: startId,
		Moves: []Move{
			{
				MType: Mem,
				Input:     "",
				MoveTo:    contentSId.String(),
				CInst: CaptureInstr{
					Inst:     Open,
					MemIndex: rc.MemoryIndex,
				},
			},
		},
		IsEnd: false,
	}

	cs, endId, _ := rc.Content.States(contentSId.String())

	exitId, _ := uuid.NewRandom()
	exitState := State{
		Id:    endId,
		IsEnd: false,
		Moves: []Move{
			{
				MType: Mem,
				Input:     "",
				MoveTo:    exitId.String(),
				CInst: CaptureInstr{
					Inst:     Close,
					MemIndex: rc.MemoryIndex,
				},
			},
		},
	}

	cs = append(cs, entryState, exitState)

	return cs, exitId.String(), nil
}
