package logging_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	logging "github.com/mrbuk/scaffolding/logging"
)

func newBufferLogger() logging.Logger {
	buf := &bytes.Buffer{}

	logger := logging.Logger{
		Formatter: &log.JSONFormatter{},
		Output:    buf,
		Level:     log.InfoLevel,
	}

	return logger
}

func getLogObj(l *logging.Logger) (map[string]interface{}, error) {
	var logObj map[string]interface{}
	buf, ok := l.Output.(*bytes.Buffer)

	if !ok {
		return nil, fmt.Errorf("Logger output is not a bytes.Buffer")
	}

	err := json.Unmarshal(buf.Bytes(), &logObj)
	if err != nil {
		return nil, fmt.Errorf("Couldn't unmarshal JSON: %v", err)
	}

	return logObj, nil
}

func TestInit(t *testing.T) {

	testLogger := newBufferLogger()
	testLogger.Init()

	log.Println("test")

	logObj, err := getLogObj(&testLogger)
	if err != nil {
		t.Error(err)
	}

	msg, ok := logObj["msg"]
	if msg != "test" || !ok {
		t.Errorf("Expected '%s' but got '%s' as part of %v", "test", msg, logObj)
	}
}

func TestLevelFromEnv(t *testing.T) {

	os.Setenv("LOG_LEVEL", "debug")

	testLogger := newBufferLogger()
	testLogger.Init()

	log.Debug("test")

	logObj, err := getLogObj(&testLogger)
	if err != nil {
		t.Error(err)
	}

	msg, ok := logObj["msg"]
	if msg != "test" || !ok {
		t.Errorf("Expected '%s' but got '%s' as part of %v", "test", msg, logObj)
	}

	level, ok := logObj["level"]
	if level != "debug" || !ok {
		t.Errorf("Expected level '%s' but got '%s' as part of %v", "debug", level, logObj)
	}

}
