package backend

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-coder/log"
)

func Stderr() log.Outputer {
	return &outter{}
}

type outter struct {
}

var _ log.Outputer = (*outter)(nil)

func (o *outter) Output(e *log.Entry) {
	str := fmt.Sprintf("I%d %s %s:%d %s [%s] %s\n",
		e.Level, e.Time.Format("2006/1/2 15:04:05"), e.FileName, e.LineNum, e.Prefix, e.Message, flatten(e.Fields))
	os.Stderr.WriteString(str)
}

// flatten returns string of sortted key-value pair
func flatten(dict map[string]*log.TypedValue) string {
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := strings.Builder{}
	sep := ""
	for _, key := range keys {
		buf.WriteString(sep)
		sep = " "
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(dict[key].Value)
	}
	return buf.String()
}
