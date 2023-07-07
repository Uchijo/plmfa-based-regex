package eval

type PosMemoryList []PosMemory

// Appended appends given suffix to opened memory.
// it doesn't modify closed memory.
// returns modified copy and original data will be left unmodified.
func (pml PosMemoryList) Appended(suffix string) PosMemoryList {
	for i, v := range pml {
		pml[i] = v.appended(suffix)
	}
	return pml
}

// OpenedMem changes status of specified memory.
// if memory not found, it creates memory that satisfies given status.
// original data is unmodified, and modified copy will be returned.
func (pml PosMemoryList) OpenedMem(index int) PosMemoryList {
	for i, v := range pml {
		if v.Index == index {
			pml[i].IsOpen = true
			return pml
		}
	}
	pml = append(
		pml,
		PosMemory{
			Index:   index,
			IsOpen:  true,
			Content: "",
		},
	)
	return pml
}

// ClosedMem closes specified memory, clean ups that memory and return content of memory.
// it doesn't change original content and returns modified copy.
func (pml PosMemoryList) ClosedMem(index int) (PosMemoryList, string) {
	for i, v := range pml {
		if v.Index == index {
			content := v.Content
			pml[i].Content = ""
			pml[i].IsOpen = false

			return pml, content
		}
	}
	return pml, ""
}

type PosMemory struct {
	Index   int
	Content string
	IsOpen  bool
}

// appended appends given suffix to memory if memory is opened.
func (pm PosMemory) appended(suffix string) PosMemory {
	if pm.IsOpen {
		pm.Content = pm.Content + suffix
	}
	return pm
}
