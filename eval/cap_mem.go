package eval

import (
	"errors"
	"fmt"
)

type CapMemoryList []CapMemory

// epsilon semantics
func (cml CapMemoryList) Content(index int, epsilonSem bool) (string, error) {
	for _, v := range cml {
		if v.Index == index {
			return v.Content, nil
		}
	}
	if epsilonSem {
		return "", nil
	} else {
		// fixme: ここの動作おかしいので直す
		message := fmt.Sprintf("unassigned reference was used; accessed %v memory but not found", index)
		return "", errors.New(message)
	}
}

func (cml CapMemoryList) OpenedMem(index int) CapMemoryList {
	copy := append([]CapMemory{}, cml...)
	for i, v := range copy {
		if v.Index == index {
			copy[i].IsOpen = true
			copy[i].Content = ""
			return copy
		}
	}
	copy = append(copy,
		CapMemory{
			Index:   index,
			IsOpen:  true,
			Content: "",
		},
	)
	return copy
}

func (cml CapMemoryList) ClosedMem(index int) CapMemoryList {
	copy := append([]CapMemory{}, cml...)
	for i, v := range copy {
		if v.Index == index {
			copy[i].IsOpen = false
			return copy
		}
	}
	return copy
}

func (cml CapMemoryList) Appended(suffix string) CapMemoryList {
	copy := append([]CapMemory{}, cml...)
	for i, v := range copy {
		copy[i] = v.appended(suffix)
	}
	return copy
}

func (cml CapMemoryList) Alike(subject CapMemoryList) bool {
	for i := range cml {
		if cml[i].alike(subject[i]) {
			continue
		} else {
			return false
		}
	}
	return true
}

type CapMemory struct {
	Index   int
	Content string
	IsOpen  bool
}

func (cm CapMemory) appended(suffix string) CapMemory {
	if cm.IsOpen {
		cm.Content = cm.Content + suffix
	}
	return cm
}

func (cm CapMemory) alike(subject CapMemory) bool {
	return cm.IsOpen == subject.IsOpen && cm.Index == subject.Index && cm.Content == subject.Content
}
