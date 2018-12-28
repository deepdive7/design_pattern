package behavioral

import "fmt"

const (
	INFO  = 1
	DEBUG = 2
	ERROR = 3
)

// Note: if Abstract impl the logMessage method, the nextLogger will all convert to Abstract and use it's write method

//Logger ...
type Logger interface {
	setNextLogger(Logger)
	logMessage(int, string)
	write(string)
}
type AbstractLogger struct {
	level      int
	nextLogger Logger
}

func (aL *AbstractLogger) write(message string) {
	fmt.Println("Console::Logger:(level:", aL.level, "), log message:", message)
}

func (aL *AbstractLogger) setNextLogger(nextLogger Logger) {
	aL.nextLogger = nextLogger
}

type InfoLogger struct {
	AbstractLogger
}

func (iL *InfoLogger) write(message string) {
	fmt.Println("Info Console::Logger:(log level:", iL.level, "):", message)
}

func (iL *InfoLogger) logMessage(level int, message string) {
	if iL.level <= level {
		iL.write(message)
		return
	}
	iL.nextLogger.logMessage(level, message)
}

type DebugLogger struct {
	AbstractLogger
}

func (dL *DebugLogger) write(message string) {
	fmt.Println("Debug Console::Logger(log level:", dL.level, "):", message)
}
func (dL *DebugLogger) logMessage(level int, message string) {
	if dL.level <= level {
		dL.write(message)
		return
	}
	dL.nextLogger.logMessage(level, message)
}

type ErrorLogger struct {
	AbstractLogger
}

func (eL *ErrorLogger) write(message string) {
	fmt.Println("Error Console::Logger:(log level:", eL.level, "):", message)
}

func (eL *ErrorLogger) logMessage(level int, message string) {
	if eL.level <= level {
		eL.write(message)
		return
	}
	eL.nextLogger.logMessage(level, message)
}
