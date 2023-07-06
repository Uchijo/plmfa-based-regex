package model

// 各状態から行える操作を示す
type Move struct {
	IsEpsilon bool
	Input     string
	MoveTo    string
	CInst     CaptureInstr
}

// Capturing Groupへの操作を示す
type CaptureInstr struct {
	Inst     MemInst
	MemIndex int
}

// メモリ操作種別を示すenum
type MemInst int

const (
	// Keep is default
	Keep MemInst = iota
	Close
	Open
)
