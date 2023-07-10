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
	for i, v := range cml {
		if v.Index == index {
			cml[i].IsOpen = true
			cml[i].Content = ""
			return cml
		}
	}
	cml = append(cml,
		CapMemory{
			Index:   index,
			IsOpen:  true,
			Content: "",
		},
	)
	return cml
}

func (cml CapMemoryList) ClosedMem(index int) CapMemoryList {
	for i, v := range cml {
		if v.Index == index {
			cml[i].IsOpen = false
			return cml
		}
	}
	return cml
}

func (cml CapMemoryList) Appended(suffix string) CapMemoryList {
	for i, v := range cml {
		cml[i] = v.appended(suffix)
	}
	return cml
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
