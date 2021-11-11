package logger_rate_limiter

type Logger struct {
	nextAllowed map[string]int
}

func Constructor() Logger {
	return Logger{
		nextAllowed: make(map[string]int),
	}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if allowedTime, ok := l.nextAllowed[message]; !ok || ok && timestamp >= allowedTime {
		l.nextAllowed[message] = timestamp + 10
		return true
	}
	return false
}
