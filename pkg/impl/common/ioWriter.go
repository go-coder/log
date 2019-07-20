package common

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-coder/log/pkg/api"
)

// NewEntryWriteFunc returns a function, which should be used to implement an EntryWriter
func NewEntryWriteFunc(w io.Writer) func(*api.Entry) {
	buffer := &bytes.Buffer{}
	buffer.Grow(bufferSize)
	return func(e *api.Entry) {
		buffer.Reset()
		if e.Level < 0 {
			fmt.Fprintf(buffer, "Er ")
		} else {
			fmt.Fprintf(buffer, "I%d ", e.Level)
		}
		fmt.Fprintf(buffer, "%s %s:%d %s [%s] %s\n",
			e.Time.Format(TimeLayout), shorten(e.FileName), e.LineNum, e.Prefix, e.Message, flatten(e.Fields))
		if e.Err != nil {
			fmt.Fprintf(buffer, "   %s %s\n", e.Err.Message, e.Err.StackTrace)
		}
		w.Write(buffer.Bytes())
	}
}
