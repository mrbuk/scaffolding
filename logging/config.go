package logging

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// DefaultLogLevel logrus logger level (INFO)
const DefaultLogLevel = log.InfoLevel

// LoggerSetings allows to set the exposed logrus logger settings
type Logger struct {
	Formatter log.Formatter
	Output    io.Writer
	Level     log.Level
}

// DefaultLogger creates a default logrus logger with
//  { JSON, Stdout, DefaultLogLevel }
var DefaultLogger = Logger{
	Formatter: &log.JSONFormatter{},
	Output:    os.Stdout,
	Level:     DefaultLogLevel,
}

// Init initilizes the logger. Should be called in the init() method
func (l *Logger) Init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(l.Formatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(l.Output)

	// Only log the warning severity or above.
	log.SetLevel(LevelFromEnv(l.Level))
}

// LevelFromEnv extracts the log level from the environment variable LOG_LEVEL
// and sets the level to default log level if environment variable is not
// provided or empty
func LevelFromEnv(defaultLevel log.Level) log.Level {

	var level = defaultLevel

	customLogLevel := os.Getenv("LOG_LEVEL")

	if customLogLevel != "" {
		var err error
		level, err = log.ParseLevel(customLogLevel)
		if err != nil {
			log.Warnf("Couldn't parse LOG_LEVEL=%s", customLogLevel)
			return log.InfoLevel
		}
	}
	return level
}
