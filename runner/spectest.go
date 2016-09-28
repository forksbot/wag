package runner

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/tsavola/wag/internal/types"
)

func spectestPrint(f io.Reader, sigs map[int64]types.Function, sigIndex int64, printer io.Writer) {
	sig, found := sigs[sigIndex]
	if !found {
		panic(fmt.Sprintf("0x%x", sigIndex))
	}

	args := make([]uint64, len(sig.Args))

	if err := binary.Read(f, byteOrder, args); err != nil {
		panic(err)
	}

	for i, t := range sig.Args {
		if i > 0 {
			if _, err := fmt.Fprint(printer, " "); err != nil {
				panic(err)
			}
		}

		x := args[len(args)-1-i] // arguments are laid out in reverse

		var err error

		switch t {
		case types.I32:
			_, err = fmt.Fprintf(printer, "%d", int32(x))

		case types.I64:
			_, err = fmt.Fprintf(printer, "%d", int64(x))

		case types.F32:
			_, err = fmt.Fprintf(printer, "%f", math.Float32frombits(uint32(x)))

		case types.F64:
			_, err = fmt.Fprintf(printer, "%f", math.Float64frombits(x))

		default:
			panic(t)
		}

		if err != nil {
			panic(err)
		}
	}

	if _, err := fmt.Fprintln(printer); err != nil {
		panic(err)
	}
}