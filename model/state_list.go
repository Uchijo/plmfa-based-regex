package model

import "errors"

type StateList struct {
	states  []State
	entryId string
}

func (sl StateList) States() []State { return sl.states }

func (sl StateList) StateById(id string) (State, error) {
	for _, v := range sl.states {
		if v.Id == id {
			return v, nil
		}
	}
	return State{}, errors.New("state with given id not found")
}
