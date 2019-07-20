package common

import (
	"sort"
	"strings"

	"github.com/go-coder/log/pkg/api"
)

// TimeLayout was used to format time
const TimeLayout = "2006/1/2 15:04:05"

const bufferSize = 4 * 1024

func shorten(fileName string) string {
	index := strings.LastIndexByte(fileName, '/')
	if index > 0 {
		return fileName[index+1:]
	}
	return fileName
}

// flatten returns string of sortted key-value pair
func flatten(dict map[string]*api.TypedValue) string {
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
