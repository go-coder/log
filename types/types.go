package types

import (
	"time"
)

const (
	TypeNil   = "nil"
	ValueNil  = "<nil>"
	TypeError = "error"
	TypeChan  = "chan"
	ValueChan = "<?>"
)

type EntryWriter interface {
	WriteEntry(*Entry)
}

type Entry struct {
	Level    int
	Time     time.Time
	FileName string
	LineNum  int
	Prefix   string
	Message  string
	Fields   map[string]*TypedValue
	Err      *ErrorRecord
}

type TypedValue struct {
	Type  string
	Value string
}

type ErrorRecord struct {
	Message    string
	StackTrace string
}
