package model

import "github.com/google/uuid"

type RegPosLa struct {
	Content  RegExp
	MemIndex int
}

func (rpl RegPosLa) States(startId string) ([]State, string, error) {
	csId, _ := uuid.NewRandom()
	entryState := State{
		Id:    startId,
		IsEnd: false,
		Moves: []Move{
			{
				MType:  PosMem,
				Input:  nil,
				MoveTo: csId.String(),
				PLInst: PosMemInstr{
					Inst:     Open,
					MemIndex: rpl.MemIndex,
				},
			},
		},
	}

	cs, endId, _ := rpl.Content.States(csId.String())

	exitId, _ := uuid.NewRandom()
	exitState := State{
		Id:    endId,
		IsEnd: false,
		Moves: []Move{
			{
				MType:  PosMem,
				Input:  nil,
				MoveTo: exitId.String(),
				PLInst: PosMemInstr{
					Inst:     Close,
					MemIndex: rpl.MemIndex,
				},
			},
		},
	}

	cs = append(cs, entryState, exitState)

	return cs, exitId.String(), nil
}
