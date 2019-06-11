package ligneous

import (
	"testing"
)

func TestLogging(t *testing.T) {
	log := NewLogger()
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Critical("critical")
}
