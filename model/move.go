package model

// 各状態から行える操作を示す
type Move struct {
	MType    MoveType
	Input    string
	MoveTo   string
	CInst    CaptureInstr
	PLInst   PosMemInstr
	RefIndex int
}

// Capturing Groupへの操作
type CaptureInstr struct {
	Inst     MemInst
	MemIndex int
}

// Positive Lookahead Memoryへの操作
type PosMemInstr struct {
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

// Moveの種別を示すenum
type MoveType int

const (
	Epsilon MoveType = iota
	Mem
	Consumption
	Ref
)
