package eval

import (
	"fmt"
	"testing"
)

func TestContainerForConsumed(t *testing.T) {
	tests := []struct {
		name      string
		buf       InputBuffer
		toConsume string
		remaining string
	}{
		{
			name:      "",
			buf:       InputBuffer{"aiueo"},
			toConsume: "a",
			remaining: "iueo",
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf("ContainerForConsumed: %s", td.name), func(t *testing.T) {
			buf, err := td.buf.Consumed(td.toConsume)
			if err != nil {
				t.Error("unexpected error occurred.")
			}
			if buf.Input != td.remaining {
				t.Errorf("expected remaining: %v, but got %v", td.remaining, buf.Input)
			}
		})
	}
}
