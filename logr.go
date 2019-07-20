package log

import (
	"github.com/go-coder/log/pkg/api"
	"github.com/go-coder/log/pkg/impl/stderr"
	"github.com/go-logr/logr"
)

var (
	// log.NewLogger is the shortcut to api.New
	NewLogger = api.New

	syslog logr.Logger = NewLogger(stderr.New())

	// log.V is the shortcut to syslog.V
	V = syslog.V
	// log.WithName is the shortcut to syslog.WithName
	WithName = syslog.WithName
	// log.WithValues is the shortcut to syslog.WithValues
	WithValues = syslog.WithValues
	// log.Info is the shortcut to syslog.Info
	Info = syslog.Info
	// log.Error is the shortcut to syslog.Error
	Error = syslog.Error
)
