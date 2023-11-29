package eval

import (
	"fmt"
	"testing"

	"github.com/uchijo/plmfa-based-regex/model"
)

func TestCanConsume(t *testing.T) {
	testCases := []struct {
		desc       string
		buf        string
		container  model.CharContainer
		canConsume bool
		toConsume  string
	}{
		{
			desc: "test a from aiueo",
			buf: "aiueo",
			container: model.SingleChar{Char: 'a'},
			canConsume: true,
			toConsume: "a",
		},
		{
			desc: "test あ from あ",
			buf: "あ",
			container: model.SingleChar{Char: 'あ'},
			canConsume: true,
			toConsume: "あ",
		},
		{
			desc: "test あ from あいうえお",
			buf: "あいうえお",
			container: model.SingleChar{Char: 'あ'},
			canConsume: true,
			toConsume: "あ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ib := InputBuffer{tC.buf}
			canConsume, toConsume := ib.CanConsume(tC.container)
			if canConsume != tC.canConsume {
				t.Errorf("expected to be %v, but got %v in canConsume result", tC.canConsume, canConsume)
			}
			if toConsume != tC.toConsume {
				t.Errorf("expected to be %v, but got %v in toConsume result", tC.toConsume, toConsume)
			}
		})
	}
}

func TestContainerForConsumed(t *testing.T) {
	tests := []struct {
		name      string
		buf       InputBuffer
		toConsume string
		remaining string
	}{
		{
			name:      "subtract a from aiueo",
			buf:       InputBuffer{"aiueo"},
			toConsume: "a",
			remaining: "iueo",
		},
		{
			name:      "consume あ from あ",
			buf:       InputBuffer{"あ"},
			toConsume: "あ",
			remaining: "",
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf("ContainerForConsumed: %s", td.name), func(t *testing.T) {
			buf, err := td.buf.Consumed(td.toConsume)
			if err != nil {
				t.Error("unexpected error occurred: ", err)
			}
			if buf.Input != td.remaining {
				t.Errorf("expected remaining: %v, but got %v", td.remaining, buf.Input)
			}
		})
	}
}
