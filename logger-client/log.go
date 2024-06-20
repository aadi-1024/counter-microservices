package loggerclient

import (
	"encoding/json"
	"time"
)

// Log holds the structure for a log entry and handles the (de)serialization
type log struct {
	Origin    string    `json:"origin"`
	Message   string    `json:"message"`
	createdAt time.Time `json:"createdAt"`
}

type LogLevel int

const (
	Info    LogLevel = 1
	Warning LogLevel = 2
	Error   LogLevel = 3
)

func createLog(origin, message string, level LogLevel) *log {
	l := &log{}

	var m string
	switch level {
	case Info:
		m = "INFO: "
	case Warning:
		m = "WARNING: "
	case Error:
		m = "ERROR: "
	default:
		m = "UNDEFINED: "
	}

	l.Origin = origin
	l.Message = m + message
	l.createdAt = time.Now()

	return l
}

func (l *log) serialize() ([]byte, error) {
	return json.Marshal(l)
}