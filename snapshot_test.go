package wag

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/tsavola/wag/internal/sexp"
	"github.com/tsavola/wag/runner"
)

func TestSnapshot(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/snapshot.wast")
	if err != nil {
		t.Fatal(err)
	}

	expr, _ := sexp.ParsePanic(data)
	if expr == nil {
		t.Fatal()
	}

	m := loadModule(expr)

	const (
		maxRODataSize = 4096
		stackSize     = 4096
	)

	b, err := runner.NewBuffer(maxRODataSize)
	if err != nil {
		t.Fatal(err)
	}

	text, _, globals, data, funcMap, callMap := m.Code(b.Imports, b.RODataAddr(), b.ROData)

	b.Seal()

	p, err := b.NewProgram(text, globals, data, funcMap, callMap, m.FuncTypes(), m.FuncNames())
	if err != nil {
		t.Fatal(err)
	}

	r1, err := p.NewRunner(m.Memory.MinSize, m.Memory.MaxSize, stackSize)
	if err != nil {
		t.Fatal(err)
	}

	var printBuf bytes.Buffer

	result, err := r1.Run(0, m.ImportTypes(), &printBuf)

	r1.Close()

	t.Logf("snapshots taken: %d", len(r1.Snapshots))

	if printBuf.Len() > 0 {
		t.Logf("print output:\n%s", string(printBuf.Bytes()))
	}

	if result < 0 {
		t.Fatal(result)
	}

	s := r1.Snapshots[int(result)]

	r2, err := s.NewRunner(m.Memory.MaxSize, stackSize)
	if err != nil {
		t.Fatal(err)
	}

	printBuf.Reset()

	result, err = r2.Run(0, m.ImportTypes(), &printBuf)

	r2.Close()

	t.Logf("snapshots taken: %d", len(r2.Snapshots))

	if printBuf.Len() > 0 {
		t.Logf("print output:\n%s", string(printBuf.Bytes()))
	}

	if result != -1 {
		t.Fatal(result)
	}

	p.Close()
	b.Close()
}