package log

import (
	"fmt"
	"log"
	"os"
	"strings"	

	"github.com/go-coder/logr"
	"github.com/spf13/pflag"
)

var (
	argV int
)

func init() {
	pflag.IntVar(&argV, "v", 2, "threshold for info output")
	pflag.Parse()
}

func New() logr.Logger {
	return &rlog{
		level: 0,
		name: "",
		fields: make([]interface{}, 0),
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *rlog) copy() *rlog {
	return &rlog{
		level: l.level,
		name: l.name,
		fields: append(make([]interface{},0), l.fields),
		logger: l.logger,
	}
}

type rlog struct {
	level int
	name string
	fields []interface{}
	logger *log.Logger
}

var _ logr.InfoLogger = (*rlog)(nil)
var _ logr.Logger = (*rlog)(nil)

func (l rlog) Enabled() bool {
	return l.level <= argV
}

func (l *rlog) Info(msg string, kvList ...interface{}) {
	if l.Enabled() {
		l.output(msg, kvList...)
	}
}

func (l *rlog) Error(err error, msg string, kvList ...interface{}) {
	log.Println(err, msg, kvList)
}

func (l *rlog) V(level int) logr.InfoLogger {
	out := l.copy()
	out.level = level
	return out
}

func (l *rlog) WithName(name string) logr.Logger {
	out := l.copy()
	if out.name == "" {
		out.name = name
	} else {
		out.name = out.name + "/" + name
	}
	return out
}

func (l *rlog) WithFields(kvList ...interface{}) logr.Logger {
	out := l.copy()
	out.fields = append(out.fields, kvList)
	return out
}

func (l *rlog) output(msg string, kvList ...interface{}) {
	if len(kvList) % 2 != 0 {
		panic("fields must be key-value pairs")
	}
	str := flatten(kvList...)
	l.logger.Println(msg, str)
}

func flatten(kvList ...interface{}) string {
	buf := strings.Builder{}
	sep := ""
	for i:=0; i<len(kvList); i+=2 {
		buf.WriteString(sep)
		sep = " "
		buf.WriteString(kvList[i].(string))
		buf.WriteString("=")
		buf.WriteString(pretty(kvList[i+1]))
	}
	return buf.String()
}

func pretty(v interface{}) string {
	return fmt.Sprintf("%v", v)
}