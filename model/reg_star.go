package model

import "github.com/google/uuid"

type RegStar struct {
	Content RegExp
}

func (rs RegStar) States(startId string) ([]State, string, error) {
	nextId, err := uuid.NewRandom()
	if err != nil {
		panic("uuid generation error.")
	}
	contentStart, err := uuid.NewRandom()
	if err != nil {
		panic("uuid generation error.")
	}
	cs, cout, err := rs.Content.States(contentStart.String())
	if err != nil {
		return nil, "", err
	}
	entry := State{
		Id:    startId,
		IsEnd: false,
		Moves: []Move{
			{
				MType:  Epsilon,
				Input:  nil,
				MoveTo: contentStart.String(),
			},
			{
				MType:  Epsilon,
				Input:  nil,
				MoveTo: cout,
			},
		},
	}
	root := State{
		Id:    cout,
		IsEnd: false,
		Moves: []Move{
			{
				MType:  Epsilon,
				Input:  nil,
				MoveTo: contentStart.String(),
			},
			{
				MType:  Epsilon,
				Input:  nil,
				MoveTo: nextId.String(),
			},
		},
	}
	states := []State{}
	states = append(states, cs...)
	states = append(states, entry, root)
	return states, nextId.String(), nil
}
