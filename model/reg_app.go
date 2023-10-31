package model

type RegApp struct {
	Contents []RegExp
}

func (ra RegApp) States(startId string) ([]State, string, error) {
	// 空の連接の場合何もせずに終わり
	if len(ra.Contents) == 0 {
		return []State{}, startId, nil
	}

	states := []State{}
	currentStart := startId
	for _, v := range ra.Contents {
		stFragment, currentEnd, err := v.States(currentStart)
		if err != nil {
			panic("something went wrong!")
		}
		states = append(states, stFragment...)
		currentStart = currentEnd
	}
	return states, currentStart, nil
}
