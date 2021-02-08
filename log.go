package gatt

// LogFunction is a callback for log messages.
type LogFunction func(string, ...interface{})

var currentLogFunctionTrace LogFunction
var currentLogFunctionInfo LogFunction
var currentLogFunctionWarning LogFunction
var currentLogFunctionError LogFunction

// SetLogFunctionTrace will set the function used for trace messages.
func SetLogFunctionTrace(f LogFunction) {
	currentLogFunctionTrace = f
}

// SetLogFunctionInfo will set the function used for informational log messages.
func SetLogFunctionInfo(f LogFunction) {
	currentLogFunctionInfo = f
}

// SetLogFunctionWarning will set the function used for warning messages.
func SetLogFunctionWarning(f LogFunction) {
	currentLogFunctionWarning = f
}

// SetLogFunctionError will set the function used for errors.
func SetLogFunctionError(f LogFunction) {
	currentLogFunctionError = f
}

func logTrace(format string, args ...interface{}) {
	if currentLogFunctionTrace != nil {
		currentLogFunctionTrace(format, args...)
	}
}

func logInfo(format string, args ...interface{}) {
	if currentLogFunctionInfo != nil {
		currentLogFunctionInfo(format, args...)
	}
}

func logWarning(format string, args ...interface{}) {
	if currentLogFunctionWarning != nil {
		currentLogFunctionWarning(format, args...)
	}
}

func logError(format string, args ...interface{}) {
	if currentLogFunctionError != nil {
		currentLogFunctionError(format, args...)
	}
}
