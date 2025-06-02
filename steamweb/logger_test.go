package steamweb_test

import (
	"github.com/13k/valve.go/steamweb"
)

type NoopLogger struct{}

var _ steamweb.Logger = (*NoopLogger)(nil)

func (l *NoopLogger) Errorf(format string, args ...interface{}) {}
func (l *NoopLogger) Warnf(format string, args ...interface{})  {}
func (l *NoopLogger) Debugf(format string, args ...interface{}) {}
