package logger

// Logger defines a common logger interface for packages that accept an
// injected logger. It fully implements the standard 'pkg/log' log package
// logging methods as well Info, Warn, Error, and Debug level equivalents.
// The Info* methods are expected to mirror the Print* methods.
//
// Packages that implement this interface should support level codes and
// expect usage as follows:
//
// 	- Panic, highest level of severity. Log and then call panic.
// 	- Fatal, log and then call `os.Exit(<code>)` with a non-zero value.
// 	- Error, used for errors that should definitely be noted and addressed.
// 	- Warn, non-critical information about undesirable behavior that needs
// 	  to be addressed in some way.
// 	- Info, general operational information about what's happening inside an
// 	  application. 'Print' methods should output Info-level information.
// 	- Debug, usually only enabled when debugging. Add anything useful, but
// 	  still exclude PII or sensitive data...
//
// Each level should include all the log levels above it. To manage the output,
// this interface also implements a SetLevel(level uint) method that defines
// the log level of this logger.
//
// Compatible logger packages include:
//
// 	- "github.com/bdlm/log"
// 	- "github.com/sirupsen/logrus"
type Logger interface {
	// WithFields adds a map of key/value data to the log entry.
	WithFields(Fields)

	// SetLevel sets the log level of this logger. Accepts any unsigned
	// integer.
	SetLevel(level Level)

	// Fatal methods should call os.Exit() with a non-zero exit status.
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	// Level 1
	// Panic methods should end with a call to the panic() builtin to
	// allow recovery, regardless of log level, but should not attempt to
	// write to any logs independent of that if the log level is below
	// PanicLevel.
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	// Level 2
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	// Level 3
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	// Level 4
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	// pkg/log compatibility, alias to Info* methods.
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})

	// Level 5
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
}
