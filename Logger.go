package Logger

import(
	"os"

	"github.com/op/go-logging"
)

var log *logging.Logger
var syslogID = "INSERT_YOUR_SISLOGID_PLEASE"

func Init(customSyslogID, loggerName, formatString string) {

	if len(formatString) <= 0 {
		formatString = "%{module} %{shortfile} > %{level:.7s} > %{message}"
	}

	log = logging.MustGetLogger(loggerName)
	log.ExtraCalldepth++ // Increase +1 to avoid wrapper call from call stack

	format := logging.MustStringFormatter(formatString)

    //file to stdout
    stdLog          := logging.NewLogBackend(os.Stderr, "", 0)
    stdLogFormatter := logging.NewBackendFormatter(stdLog, format)

    //log to syslog
    syslogLogger, err := logging.NewSyslogBackend(syslogID)
    if err != nil {
    	panic(err)
    }
    syslogFormatter   := logging.NewBackendFormatter(syslogLogger, format)

    log.SetBackend(logging.MultiLogger(stdLogFormatter, syslogFormatter))
}

func GetSyslogID() string {
	return syslogID
}

func SetSyslogID(id string) {
	syslogID = id
}

func Debug(v ...interface{}) {
	if log != nil {
		log.Debug(v)
	}
}

func Error(v ...interface{}) {
	if log != nil {
		log.Error(v)
	}
}

func Warning(v ...interface{}) {
	if log != nil {
		log.Warning(v)
	}
}

func Fatal(v ...interface{}) {
	if log != nil {
		log.Fatal(v)
	}
}