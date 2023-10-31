package model

import "github.com/google/uuid"

// "任意の1文字"に当たるもの
type RegArb struct {
}

func (ra RegArb) States(startId string) ([]State, string, error) {
	nextId, _ := uuid.NewRandom()
	state := State{
		Id:    startId,
		IsEnd: false,
		Moves: []Move{
			{
				MType:  ArbitraryConsumption,
				Input:  "",
				MoveTo: nextId.String(),
			},
		},
	}
	return []State{state}, nextId.String(), nil
}
