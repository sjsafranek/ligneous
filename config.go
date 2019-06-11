package ligneous

import (
	"fmt"
)

const (
	nocolor = 0
	red     = 31 // error critical
	green   = 32
	yellow  = 33 // warn
	blue    = 36 // info
	grey    = 37 // debug
)

// https://en.wikipedia.org/wiki/ANSI_escape_code#3/4_bit
// https://github.com/cihub/seelog/wiki/Log-levels
func getConfig(name, level, path string) string {

	rollingfile := ""
	if "" != path {
		rollingfile = fmt.Sprintf(`<rollingfile type="size" filename="%v/%v.log" maxsize="10000000" maxrolls="5" />`, path, name)
	}

	formats := `
		<formats>
			<format id="stdout"  format="%Date %Time [%LEVEL] [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
			<format id="debug"   format="%Date %Time %EscM(37)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
			<format id="info"    format="%Date %Time %EscM(36)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
			<format id="warn"    format="%Date %Time %EscM(33)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
			<format id="error"   format="%Date %Time %EscM(31)[%LEVEL]%EscM(0) [PID-%pidLogFormatter] %File %FuncShort:%Line %Msg %n" />
		</formats>
	`

	return fmt.Sprintf(`
        <seelog minlevel="%v">
            <outputs formatid="stdout">
            <filter levels="debug,trace">
                <console formatid="debug"/>
				%v
            </filter>
            <filter levels="info">
                <console formatid="info"/>
				%v
            </filter>
            <filter levels="critical,error">
                <console formatid="error"/>
				%v
            </filter>
            <filter levels="warn">
                <console formatid="warn"/>
				%v
            </filter>
            </outputs>
            %v
        </seelog>`, level, rollingfile, rollingfile, rollingfile, rollingfile, formats)
}
