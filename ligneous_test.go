package ligneous

import (
	"testing"
)

func TestLogging(t *testing.T) {
	log := AddLogger("app","debug","logs")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Critical("critical")
}
