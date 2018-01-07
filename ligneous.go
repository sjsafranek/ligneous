package ligneous

import (
	"fmt"
	"os"
	"strings"

	seelog "github.com/cihub/seelog"
)

var (
	Verbose bool = false
	Logger  seelog.LoggerInterface
	Level   string = "debug"
)

const (
	nocolor = 0
	red     = 31 // error critical
	green   = 32
	yellow  = 33 // warn
	blue    = 36 // info
	grey    = 37 // debug
)

func formatter(message []interface{}) string {
	var text []string
	for i := range message {
		text = append(text, fmt.Sprintf("%v", message[i]))
	}
	return strings.Join(text, " ")
}

func Trace(message ...interface{}) {
	Logger.Trace(formatter(message))
}

func Debug(message ...interface{}) {
	Logger.Debug(formatter(message))
}

func Info(message ...interface{}) {
	Logger.Info(formatter(message))
}

func Warn(message ...interface{}) {
	Logger.Warn(formatter(message))
}

func Error(message ...interface{}) {
	Logger.Error(formatter(message))
}

func Critical(message ...interface{}) {
	Logger.Critical(formatter(message))
}

// https://github.com/cihub/seelog/wiki/Custom-formatters
func pidLogFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel, context seelog.LogContextInterface) interface{} {
		var pid = os.Getpid()
		return fmt.Sprintf("%v", pid)
	}
}

func initLogging() {
	if Verbose {
		Level = "trace"
	}

	Logger = seelog.Disabled

	// https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit
	// https://github.com/cihub/seelog/wiki/Log-levels
	appConfig := `
<seelog minlevel="` + Level + `">
    <outputs formatid="stdout">
	<filter levels="debug,trace">
		<console formatid="debug"/>
	</filter>
    <filter levels="info">
        <console formatid="info"/>
    </filter>
	<filter levels="critical,error">
        <console formatid="error"/>
    </filter>
	<filter levels="warn">
        <console formatid="warn"/>
    </filter>
    </outputs>
    <formats>
		<format id="stdout"   format="%Date %Time [%LEVEL] [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />

		<format id="debug"   format="%Date %Time %EscM(37)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
		<format id="info"    format="%Date %Time %EscM(36)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
		<format id="warn"    format="%Date %Time %EscM(33)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
		<format id="error"   format="%Date %Time %EscM(31)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />

	</formats>
</seelog>
`

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}
	Logger = logger
}

func init() {
	seelog.RegisterCustomFormatter("pidLogFormatter", pidLogFormatter)
	initLogging()
}
