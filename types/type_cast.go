package types

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
)

type binWriter struct {
	w   io.Writer
	buf bytes.Buffer
	err error
}

func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	// switch v.(type) {
	// case string:
	// 	s := v.(string)
	// 	w.Write(int32(len(s)))
	// 	w.Write([]byte(s))
	// case int:
	// 	i := v.(int)
	// 	w.Write(int64(i))
	// default:
	// 	if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
	// 		w.size += int64(binary.Size(v))
	// 	}
	// }

	// type cast switch:
	switch x := v.(type) {
	case string:
		w.Write(int32(len(x)))
		w.Write([]byte(x))
	case int:
		w.Write(int64(x))
	default:
		w.err = binary.Write(&w.buf, binary.LittleEndian, v)
	}
}

// writing everything or nothing
func (w *binWriter) Flush() (int64, error) {
	if w.err != nil {
		return 0, w.err
	}
	return w.buf.WriteTo(w.w)
}

type Gopher struct {
	Name     string
	AgeYears int
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
	bw := &binWriter{w: w}
	bw.Write(g.Name)
	bw.Write(g.AgeYears)
	return bw.Flush()
}

func TestGopherWriter() {
	g := &Gopher{
		Name:     "Davidisawesome",
		AgeYears: 9999,
	}

	if _, err := g.WriteTo(os.Stdout); err != nil {
		log.Printf("DumpBinary: %v\n", err)
	}
}
