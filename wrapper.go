package ligneous

import (
	seelog "github.com/cihub/seelog"
)

type Log seelog.LoggerInterface

func isValidLevel(level string) bool {
	levels := [6]string{"debug", "trace", "info", "critical", "error", "warn"}
	for i := range levels {
		if levels[i] == level {
			return true
		}
	}
	return false
}

func AddLogger(name, level, path string) seelog.LoggerInterface {
	if "" == level || !isValidLevel(level) {
		level = DEFAULT_LEVEL
	}

	logConfig := getConfig(name, level, path)

	log, err := seelog.LoggerFromConfigAsBytes([]byte(logConfig))
	if nil != err {
		panic(err)
	}
	return log
}

func New() seelog.LoggerInterface {
	return AddLogger("", "", "")
}

func NewLogger() seelog.LoggerInterface {
	return New()
}
