package eval

type CapMemoryList []CapMemory

// epsilon semantics
func (cml CapMemoryList) Content(index int) string {
	for _, v := range cml {
		if v.Index == index {
			return v.Content
		}
	}
	return ""
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
