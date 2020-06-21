// package custom is custom current logger associate with logrus
package custom

import (
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sirupsen/logrus"
)

// custom library logrus in gokit
type logrusLogger struct {
	logrus.FieldLogger
}

var errMissingValue = errors.New("(MISSING)")

// NewLogrusLogger returns a go-kit log.Logger that sends log events to a Logrus logger.
func NewLogrusLogger(logger logrus.FieldLogger) log.Logger {
	return &logrusLogger{logger}
}

func (l logrusLogger) Log(keyvals ...interface{}) error {
	var description string
	var levelLog string

	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		if keyvals[i] == "description" { // unique key
			description = keyvals[i+1].(string) // assertion interface to string
			continue
		}

		if keyvals[i] == "level" { // unique key
			// assertion interface{} to level.Value type in library level go-kit
			// and take value name with method String()
			levelLog = keyvals[i+1].(level.Value).String()
			continue
		}

		if i+1 < len(keyvals) {
			fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
		} else {
			fields[fmt.Sprint(keyvals[i])] = errMissingValue
		}
	}

	switch levelLog { // case by value
	case level.InfoValue().String():
		l.WithFields(fields).Info(description)
	case level.ErrorValue().String():
		l.WithFields(fields).Error(description)
	case level.WarnValue().String():
		l.WithFields(fields).Warn(description)
	}

	return nil
}
