package ligneous

const (
	nocolor = 0
	red     = 31 // error critical
	green   = 32
	yellow  = 33 // warn
	blue    = 36 // info
	grey    = 37 // debug
)

func isValidLevel(level string) bool {
	levels := [6]string{"debug", "trace", "info", "critical", "error", "warn"}
	for i := range levels {
		if levels[i] == level {
			return true
		}
	}
	return false
}

// https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit
// https://github.com/cihub/seelog/wiki/Log-levels
func getConfig(level string) string {
	return `
        <seelog minlevel="` + level + `">
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
        </seelog>`
}
