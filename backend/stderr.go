package backend

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-coder/log/types"
)

func Stderr() types.EntryWriter {
	out := &outter{
		entryChan: make(chan *types.Entry),
	}
	go func() {
		for e := range out.entryChan {
			doWrite(e)
		}
	}()
	return out
}

type outter struct {
	entryChan chan *types.Entry
}

var _ types.EntryWriter = (*outter)(nil)

func (o *outter) WriteEntry(e *types.Entry) {
	o.entryChan <- e
}

func doWrite(e *types.Entry) {
	var str string
	if e.Level < 0 {
		if e.Err != nil {
			str = fmt.Sprintf("Er %s %s:%d %s [%s] %s\n   %s %s\n",
				e.Time.Format("2006/1/2 15:04:05"), shorten(e.FileName), e.LineNum, e.Prefix, e.Message, flatten(e.Fields),
				e.Err.Message, e.Err.StackTrace)
		} else {
			str = fmt.Sprintf("Er %s %s:%d %s [%s] %s\n",
				e.Time.Format("2006/1/2 15:04:05"), shorten(e.FileName), e.LineNum, e.Prefix, e.Message, flatten(e.Fields))
		}
	} else {
		str = fmt.Sprintf("I%d %s %s:%d %s [%s] %s\n",
			e.Level, e.Time.Format("2006/1/2 15:04:05"), shorten(e.FileName), e.LineNum, e.Prefix, e.Message, flatten(e.Fields))
	}
	os.Stderr.WriteString(str)
}

func shorten(fileName string) string {
	index := strings.LastIndexByte(fileName, '/')
	if index > 0 {
		return fileName[index+1:]
	}
	return fileName
}

// flatten returns string of sortted key-value pair
func flatten(dict map[string]*types.TypedValue) string {
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
