package frontend

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/go-coder/log/types"
	"github.com/go-logr/logr"
	"github.com/spf13/pflag"
)

const (
	AutoGeneratedName = "<AutoGenerated>"
)

var (
	argV              int
	argShowErrorTrace bool
)

func init() {
	pflag.IntVar(&argV, "v", math.MaxInt32, "threshold for info output")
	pflag.BoolVar(&argShowErrorTrace, "show-error-trace", true, "show error trace")
	pflag.Parse()
}

func New(outter types.EntryWriter) logr.Logger {
	return &rlog{
		level:  0,
		name:   "",
		fields: make([]interface{}, 0),
		outter: outter,
	}
}

func (l *rlog) clone() *rlog {
	return &rlog{
		level:  l.level,
		name:   l.name,
		fields: append(make([]interface{}, 0), l.fields...),
		outter: l.outter,
	}
}

type rlog struct {
	level  int
	name   string
	fields []interface{}
	outter types.EntryWriter
}

var _ logr.InfoLogger = (*rlog)(nil)
var _ logr.Logger = (*rlog)(nil)

func (l *rlog) V(level int) logr.InfoLogger {
	if level < 0 {
		panic(fmt.Sprintf("log.V(level int) must received a non-negative parameter, but get %d", level))
	}
	out := l.clone()
	out.level = level
	return out
}

func (l *rlog) WithName(name string) logr.Logger {
	out := l.clone()
	if out.name == "" {
		out.name = name
	} else {
		out.name = l.name + "/" + name
	}
	return out
}

func (l *rlog) WithValues(kvList ...interface{}) logr.Logger {
	out := l.clone()
	out.fields = append(out.fields, kvList...)
	return out
}

func (l rlog) Enabled() bool {
	return l.level <= argV
}

func (l *rlog) Info(msg string, kvList ...interface{}) {
	if l.Enabled() {
		l.outter.WriteEntry(l.getEntry(msg, kvList))
	}
}

func (l *rlog) Error(err error, msg string, kvList ...interface{}) {
	entry := l.getEntry(msg, kvList)
	entry.Level = -1
	if err != nil {
		entry.Err = &types.ErrorRecord{Message: err.Error()}
		if argShowErrorTrace {
			entry.Err.StackTrace = string(debug.Stack())
		}
	}
	l.outter.WriteEntry(entry)
}

func (l *rlog) getEntry(msg string, kvList []interface{}) *types.Entry {
	if len(kvList)%2 != 0 {
		panic("fields must be key-value pairs")
	}
	// callDepth 0 represents current line
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = AutoGeneratedName
		line = 0
	}
	return &types.Entry{
		Level:    l.level,
		Time:     time.Now(),
		FileName: file,
		LineNum:  line,
		Prefix:   l.name,
		Message:  msg,
		// attention that must be append(kvList, l.fields...) rather than append(l.fields, kvList...)
		Fields: kvListToMap(append(kvList, l.fields...)),
	}
}

// kvList must be key-value pair
func kvListToMap(kvList []interface{}) map[string]*types.TypedValue {
	dict := make(map[string]*types.TypedValue)
	for i := 0; i < len(kvList); i += 2 {
		key := kvList[i].(string)
		value := getTypedValue(kvList[i+1])
		dict[key] = value
	}
	return dict
}

func getTypedValue(v interface{}) *types.TypedValue {
	if v == nil {
		return &types.TypedValue{
			Type:  types.TypeNil,
			Value: types.ValueNil,
		}
	}
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return &types.TypedValue{
				Type:  reflect.TypeOf(v).String(),
				Value: types.ValueNil,
			}
		}
		v = reflect.Indirect(reflect.ValueOf(v)).Interface()
	}
	if reflect.ValueOf(v).Kind() == reflect.Chan {
		return &types.TypedValue{
			Type:  types.TypeChan,
			Value: types.ValueChan,
		}
	}
	return &types.TypedValue{
		Type:  fmt.Sprintf("%T", v),
		Value: fmt.Sprintf("%v", v),
	}
}
