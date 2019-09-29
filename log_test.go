package log

import (
	"testing"
)

func Test_Log(t *testing.T) {
	logger := NewLogger("", nil)
	logger.LogV("Verbose log")
	logger.LogD("Debug log")
	logger.LogI("Info log")
	logger.LogW("Warning log")
	logger.LogE("Error log")

	SetDefaulLevel(LEVEL_VERBOSE)
	logger.LogV("Verbose log 2")
	logger.LogD("Debug log 2")
	logger.LogI("Info log 2")
	logger.LogW("Warning log 2")
	logger.LogE("Error log 2")
}
