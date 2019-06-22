package main

import (
	"errors"

	"github.com/go-coder/log"
)

func main() {
	logr := log.New()

	logr.V(1).Info("msg", "int", 211)
	logr.V(3).Info("msg", "float", 2.33)

	err := errors.New("myerr")
	logr.Error(err, "msggg", "map", map[string]int{"a": 12})
	logr.Info("mmseg", "array", [...]int{1, 0, 2, 4})
}
