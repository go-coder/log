package common

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-coder/log/pkg/api"
)

// NewStringWriteFunc returns a function, which should be used to implement an EntryWriter
func NewStringWriteFunc(w io.StringWriter) func(*api.Entry) {
	buffer := &strings.Builder{}
	buffer.Grow(bufferSize)
	return func(e *api.Entry) {
		buffer.Reset()
		if e.Level < 0 {
			buffer.WriteString("Er ")
		} else {
			buffer.WriteString(fmt.Sprintf("I%d ", e.Level))
		}
		w.WriteString(fmt.Sprintf("%s %s:%d %s [%s] %s\n",
			e.Time.Format(TimeLayout), shorten(e.FileName), e.LineNum, e.Prefix, e.Message, flatten(e.Fields)))
		if e.Err != nil {
			buffer.WriteString(fmt.Sprintf("   %s %s\n", e.Err.Message, e.Err.StackTrace))
		}
		w.WriteString(buffer.String())
	}
}
