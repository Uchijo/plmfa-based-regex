package model

type State struct {
	Id    string
	Moves []Move
	IsEnd bool
}

func (st State) ExtractMove(moveType MoveType) []Move {
	moves := []Move{}
	for _, v := range st.Moves {
		if v.MType == moveType {
			moves = append(moves, v)
		}
	}

	return moves
}
