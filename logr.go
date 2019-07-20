package log

import (
	"github.com/go-coder/log/pkg/api"
	"github.com/go-coder/log/pkg/impl/stderr"
	"github.com/go-logr/logr"
)

var (
	NewLogger = api.New

	syslog logr.Logger = NewLogger(stderr.Stderr())

	V          = syslog.V
	WithName   = syslog.WithName
	WithValues = syslog.WithValues
	Info       = syslog.Info
	Error      = syslog.Error
)
