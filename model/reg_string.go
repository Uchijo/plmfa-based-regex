package model

import (
	"fmt"

	"github.com/google/uuid"
)

type RegString struct {
	Content string
}

func (rs RegString) States(startId string) ([]State, string, error) {
	states := []State{}

	var currentId, nextId string
	currentId = startId
	for _, c := range rs.Content {
		rawId, err := uuid.NewRandom()
		if err != nil {
			return nil, "", err
		}
		nextId = rawId.String()
		moves := []Move{
			{
				MType:  Consumption,
				Input:  SingleChar{c},
				MoveTo: nextId,
			},
		}
		states = append(states, State{
			Id:    currentId,
			Moves: moves,
			IsEnd: false,
		})
		currentId = nextId
	}

	return states, currentId, nil
}

// 1文字だけ、[]なしでも使えるやつ
type SingleChar struct {
	Char rune
}

var _ CharContainer = (*SingleChar)(nil)

func (sc SingleChar) WholeMatches(c string) bool {
	fmt.Println("whole match debug print 1: " + string(sc.Char))
	fmt.Println("whole match debug print 2: " + c)
	return string(sc.Char) == c
}

func (sc SingleChar) Len() int { return 1 }
