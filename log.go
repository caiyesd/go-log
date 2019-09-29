package log

import (
	"fmt"
	"log"
	"os"
)

const (
	LEVEL_VERBOSE = 1
	LEVEL_DEBUG   = 2
	LEVEL_INFO    = 3
	LEVEL_WARNING = 4
	LEVEL_ERROR   = 5
	LEVEL_FATAL   = 6
)

var gConfig = config{
	LEVEL_DEBUG,
	make(map[string]int),
	log.Lshortfile | log.LstdFlags | log.Lmicroseconds,
}

type config struct {
	defaultLogLevel int
	logLevelMap     map[string]int
	flags           int
}

func (c *config) isLoggable(name string, level int) bool {
	if l, ok := c.logLevelMap[name]; !ok {
		return level >= c.defaultLogLevel
	} else {
		return level >= l
	}
}

func SetLevel(name string, level int) {
	gConfig.logLevelMap[name] = level
}

func SetDefaulLevel(level int) {
	gConfig.defaultLogLevel = level
}

func SetFlags(flags int) {
	gConfig.flags = flags
}

type Logger interface {
	LogV(format string, values ...interface{})
	LogD(format string, values ...interface{})
	LogI(format string, values ...interface{})
	LogW(format string, values ...interface{})
	LogE(format string, values ...interface{})
	LogF(format string, values ...interface{})
}

func NewLogger(name string, file *os.File) Logger {
	if file == nil {
		file = os.Stderr
	}
	return &levelLogger{
		name,
		log.New(os.Stderr, "", gConfig.flags),
	}
}

type levelLogger struct {
	name   string
	logger *log.Logger
}

func (l *levelLogger) log(level int, format string, values ...interface{}) {
	if gConfig.isLoggable(l.name, level) {
		depth := 3
		message := fmt.Sprintf(format, values...)
		if l.name != "" {
			message = fmt.Sprintf("[%s] %s", l.name, message)
		}
		switch level {
		case LEVEL_VERBOSE:
			l.logger.Output(depth, fmt.Sprintf("V %s\n", message))
			break
		case LEVEL_DEBUG:
			l.logger.Output(depth, fmt.Sprintf("D %s\n", message))
			break
		case LEVEL_INFO:
			l.logger.Output(depth, fmt.Sprintf("I %s\n", message))
			break
		case LEVEL_WARNING:
			l.logger.Output(depth, fmt.Sprintf("W %s\n", message))
			break
		case LEVEL_ERROR:
			l.logger.Output(depth, fmt.Sprintf("E %s\n", message))
			break
		case LEVEL_FATAL:
			l.logger.Output(depth, fmt.Sprintf("F %s\n", message))
			panic(message)
			break
		default:
			panic(fmt.Sprintf("unsupported log level %d", level))
		}
	}

}

func (l *levelLogger) LogV(format string, values ...interface{}) {
	l.log(LEVEL_VERBOSE, format, values...)
}

func (l *levelLogger) LogD(format string, values ...interface{}) {
	l.log(LEVEL_DEBUG, format, values...)
}

func (l *levelLogger) LogI(format string, values ...interface{}) {
	l.log(LEVEL_INFO, format, values...)
}

func (l *levelLogger) LogW(format string, values ...interface{}) {
	l.log(LEVEL_WARNING, format, values...)
}

func (l *levelLogger) LogE(format string, values ...interface{}) {
	l.log(LEVEL_ERROR, format, values...)
}

func (l *levelLogger) LogF(format string, values ...interface{}) {
	l.log(LEVEL_FATAL, format, values...)
}
