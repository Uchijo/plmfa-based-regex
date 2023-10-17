package model

type RegSkip struct {
	RegExp
}

func (rs RegSkip) States(nextId string) ([]State, string, error) {
	return []State{}, nextId, nil
}
