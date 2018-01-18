package ligneous

import (
	"fmt"

	seelog "github.com/cihub/seelog"
)

type SeelogWrapper struct {
	Log   seelog.LoggerInterface
	Level string
	// isValidLevel() bool
}

func (self *SeelogWrapper) init() (err error) {
	if "" == self.Level {
		self.Level = "debug"
	}

	self.Log = seelog.Disabled

	logConfig := getConfig(Level)

	self.Log, err = seelog.LoggerFromConfigAsBytes([]byte(logConfig))
	return
}

func (self *SeelogWrapper) isValidLevel(level string) bool {
	return isValidLevel(level)
}

func (self *SeelogWrapper) SetLevel(level string) error {
	if !self.isValidLevel(level) {
		return fmt.Errorf("Mot a valid logging level")
	}
	self.Level = level
	return self.init()
}

func New() SeelogWrapper {
	logger := SeelogWrapper{Level: "debug"}
	logger.init()
	return logger
}
