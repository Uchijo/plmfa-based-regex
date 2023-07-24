package eval

type Log struct {
	Input          string
	CurrentId      string
	PositiveMemory PosMemoryList
	CaptureMemory  CapMemoryList
}

func (l Log) Alike(log Log) bool {
	return l.Input == log.Input &&
		l.CurrentId == log.CurrentId &&
		l.PositiveMemory.Alike(log.PositiveMemory) &&
		l.CaptureMemory.Alike(log.CaptureMemory)
}
