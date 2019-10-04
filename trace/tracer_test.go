package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Fatal("Return form New shoudn't be null")
	}
	tracer.Trace("Hello trace package.")
	if buf.String() != "Hello trace package.\n" {
		t.Errorf("Trace shoundn't write '%s'", buf.String())
	}
}

func TestOff(t *testing.T) {
	tracer := Off()
	tracer.Trace("Don't trace")
}
