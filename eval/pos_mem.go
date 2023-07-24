package eval

type PosMemoryList []PosMemory

// Appended appends given suffix to opened memory.
// it doesn't modify closed memory.
// returns modified copy and original data will be left unmodified.
func (pml PosMemoryList) Appended(suffix string) PosMemoryList {
	copy := append([]PosMemory{}, pml...)
	for i, v := range copy {
		copy[i] = v.appended(suffix)
	}
	return copy
}

// OpenedMem changes status of specified memory.
// if memory not found, it creates memory that satisfies given status.
// original data is unmodified, and modified copy will be returned.
func (pml PosMemoryList) OpenedMem(index int) PosMemoryList {
	copy := append([]PosMemory{}, pml...)
	for i, v := range copy {
		if v.Index == index {
			copy[i].IsOpen = true
			return copy
		}
	}
	copy = append(
		copy,
		PosMemory{
			Index:   index,
			IsOpen:  true,
			Content: "",
		},
	)
	return copy
}

// ClosedMem closes specified memory, clean ups that memory and return content of memory.
// it doesn't change original content and returns modified copy.
func (pml PosMemoryList) ClosedMem(index int) (PosMemoryList, string) {
	copy := append([]PosMemory{}, pml...)
	for i, v := range copy {
		if v.Index == index {
			content := v.Content
			copy[i].Content = ""
			copy[i].IsOpen = false

			return copy, content
		}
	}
	return copy, ""
}

func (pml PosMemoryList) Alike(subject PosMemoryList) bool {
	for i := range pml {
		if pml[i].alike(subject[i]) {
			continue
		} else {
			return false
		}
	}
	return true
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

func (pm PosMemory) alike(subject PosMemory) bool {
	return pm.IsOpen == subject.IsOpen && pm.Index == subject.Index && pm.Content == subject.Content
}
