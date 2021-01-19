package logger

type Level uint32

// These are the different logging levels. You can set the logging level to
// log on your instance of logger, obtained with `New()`.
const (
	// Fatal level. Logs and then calls `os.Exit(1)`. It should execute
	// regardless of the set logging level.
	Fatal Level = 0
	// Panic level, highest level of severity. Logs and then calls
	// panic with the message passed to enable recovery.
	Panic Level = 10
	// Error level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	Error Level = 20
	// Warn level. Non-critical entries that deserve eyes.
	Warn Level = 30
	// Info level. General operational entries about what's going on inside the
	// application.
	Info Level = 40
	// Debug level. Usually only enabled when debugging. Very verbose logging.
	Debug Level = 50
	// Trace level. Usually only enabled when debugging. Extremely verbose logging, generally including a full stack trace.
	Trace Level = 60
)

// Fields defines a map of fields to be passed to WithFields()
type Fields map[string]interface{}
