package model

type RegQuantifier struct {
	RegExp
	Content RegExp

	// 最低繰り返し数
	Min int

	// 0のときは上限なしとみなす
	Max int
}

func (rq RegQuantifier) States(startId string) ([]State, string, error) {
	if rq.Max < rq.Min || rq.Max < 0 || rq.Min < 0 {
		panic("invalid quantifier specified.")
	}

	// 上限なし、下限なしは * と同じ
	if rq.Max == 0 && rq.Min == 0 {
		return RegStar{Content: rq.Content}.States(startId)
	}

	nextId := startId
	states := []State{}
	var s []State
	var err error
	// 上限下限が同じ -> その回数だけ繰り返す
	if rq.Max == rq.Min {
		for i := 0; i < rq.Max; i++ {
			s, nextId, err = rq.Content.States(nextId)
			if err != nil {
				panic("something went wrong")
			}

			states = append(states, s...)
		}
		return states, nextId, nil
	}

	// 上限なし -> 最低その回数繰り返した後 * につなぐ
	if rq.Max == 0 {
		for i := 0; i < rq.Min; i++ {
			s, nextId, err = rq.Content.States(nextId)
			if err != nil {
				panic("something went wrong")
			}

			states = append(states, s...)
		}

		s, nextId, err := RegStar{Content: rq.Content}.States(nextId)
		if err != nil {
			panic("something went wrong")
		}
		states = append(states, s...)

		return states, nextId, nil
	}

	// 上限下限あり -> 下限は一列につなぎ、上限ゾーンに入ったらいつでも終われる
	// ブランチを繋いでいく
	// 下限分を作る
	for i := 0; i < rq.Min; i++ {
		s, nextId, err = rq.Content.States(nextId)
		if err != nil {
			panic("something went wrong")
		}

		states = append(states, s...)
	}

	// 上限部分と上段のε遷移
	union := mayRepeat(rq.Content, rq.Max-rq.Min)
	s, nextId, err = union.States(nextId)
	if err != nil {
		panic("something went wrong")
	}
	states = append(states, s...)

	return states, nextId, nil
}

func mayRepeat(content RegExp, n int) RegUnion {
	if n < 1 {
		panic("n in mayRepeat(RegExp, n) needs to be bigger than 0")
	}
	if n == 1 {
		return RegUnion{
			Left:  RegSkip{},
			Right: content,
		}
	}

	return RegUnion{
		Left:  mayRepeat(content, n-1),
		Right: content,
	}
}
