package model

import "github.com/google/uuid"

type RegCapRef struct {
	MemIndex int
}

func (rcr RegCapRef) States(startId string) ([]State, string, error) {
	exitId, _ := uuid.NewRandom()
	s := State{
		Id: startId,
		Moves: []Move{
			{
				MType:    Ref,
				RefIndex: rcr.MemIndex,
				MoveTo:   exitId.String(),
				Input:    nil,
			},
		},
	}

	return []State{s}, exitId.String(), nil
}
