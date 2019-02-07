package trace

import (
  "io"
  "fmt"
)

type Tracer interface {
  Trace(...interface{})
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}){}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
  out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func Off() Tracer {
  return &nilTracer{}
}
