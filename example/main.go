package main

import (
	"errors"

	"github.com/go-coder/log"
	"github.com/go-coder/log/backend"
)

func main() {
	logr := log.New(backend.Stderr())

	logr.V(1).Info("msg", "uint", 112, "int", 211, "nil", nil)
	var typedNil *int
	logr.V(3).Info("msg", "float", 2.33, "typedNil", typedNil)

	err := errors.New("myerr")
	logr.Error(err, "msggg", "map", map[string]int{"a": 12})

	ch := make(chan int)
	logr.Info("mmseg", "array", [...]int{1, 0, 2, 4}, "channel", ch)

	a := 3
	err = errors.New("field err")
	logr.Info("message", "error", err, "*int", &a)
}
