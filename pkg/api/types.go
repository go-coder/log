package api

import (
	"time"
)

const (
	// TypeNil is the Type used by TypedValue
	TypeNil = "nil"
	// ValueNil is the Value used by TypedValue
	ValueNil = "<nil>"
	// TypeError is the Type used by TypedValue
	TypeError = "error"
	// TypeChan is the Type used by TypedValue
	TypeChan = "chan"
	// ValueChan is the Value used by TypedValue
	ValueChan = "<?>"
)

// EntryWriter is the backend log interface
type EntryWriter interface {
	WriteEntry(*Entry)
}

// Entry was used to record a log-entry
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

// TypedValue was used to record values
type TypedValue struct {
	Type  string
	Value string
}

// ErrorRecord was used to record an error
type ErrorRecord struct {
	Message    string
	StackTrace string
}
