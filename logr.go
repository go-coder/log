package log

import (
	"github.com/go-coder/log/backend"
	"github.com/go-coder/log/frontend"
	"github.com/go-logr/logr"
)

var (
	NewLogger = frontend.New

	syslog logr.Logger = NewLogger(backend.Stderr())

	V          = syslog.V
	WithName   = syslog.WithName
	WithValues = syslog.WithValues
	Info       = syslog.Info
	Error      = syslog.Error
)
