package ligneous

import (
	"testing"
)

func TestLogging(t *testing.T) {
	l := New()
	log := l.Log
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Critical("critical")

	err := l.SetLevel("critical")
	if nil != err {
		panic(err)
	}
	log = l.Log
	// should only show critical
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Critical("critical")
}
