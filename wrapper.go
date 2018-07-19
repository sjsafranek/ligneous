package ligneous

import (
	"fmt"

	seelog "github.com/cihub/seelog"
)

type Log struct {
	seelog.LoggerInterface
}

type SeelogWrapper struct {
	Log   seelog.LoggerInterface
	Level string
	// isValidLevel() bool
}

func (self *SeelogWrapper) init() (err error) {
	if "" == self.Level || !self.isValidLevel(self.Level) {
		self.Level = DEFAULT_LEVEL
	}

	self.Log = seelog.Disabled

	logConfig := getConfig(self.Level)

	self.Log, err = seelog.LoggerFromConfigAsBytes([]byte(logConfig))
	return
}

func (self *SeelogWrapper) isValidLevel(level string) bool {
	return isValidLevel(level)
}

func (self *SeelogWrapper) SetLevel(level string) error {
	if !self.isValidLevel(level) {
		return fmt.Errorf("Not a valid logging level")
	}
	self.Level = level
	return self.init()
}

func New() SeelogWrapper {
	logger := SeelogWrapper{Level: "debug"}
	logger.init()
	return logger
}

func NewLogger() seelog.LoggerInterface {
	wrapper := New()
	return wrapper.Log
}
