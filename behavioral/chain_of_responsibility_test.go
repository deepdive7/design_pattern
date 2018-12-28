package behavioral

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	loggerChain := &ErrorLogger{AbstractLogger{level: ERROR}}
	debugLogger := &DebugLogger{AbstractLogger{level: DEBUG}}
	infoLogger := &InfoLogger{AbstractLogger{level: INFO}}
	loggerChain.setNextLogger(debugLogger)
	debugLogger.setNextLogger(infoLogger)

	loggerChain.logMessage(INFO, "This is an information.")
	loggerChain.logMessage(DEBUG, "This is an debug level information.")
	loggerChain.logMessage(ERROR, "This is an error information.")
}
