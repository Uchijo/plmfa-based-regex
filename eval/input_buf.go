package eval

import "errors"

type InputBuffer struct {
	Input string
}

func (ib InputBuffer) Len() int {
	return len(ib.Input)
}

func (ib InputBuffer) CanConsume(matcher string) bool {
	prefixLen := len(matcher)
	inputLen := len(ib.Input)
	if inputLen < prefixLen {
		return false
	}
	prefix := ib.Input[:prefixLen]
	return prefix == matcher
}

func (ib InputBuffer) Consumed(matcher string) (InputBuffer, error) {
	if !ib.CanConsume(matcher) {
		return InputBuffer{}, errors.New("cannot consume")
	}
	ib.Input = ib.Input[len(matcher):]
	return ib, nil
}

func (ib InputBuffer) Appended(prefix string) (InputBuffer, error) {
	ib.Input = prefix + ib.Input
	return ib, nil
}
