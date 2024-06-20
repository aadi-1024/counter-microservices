package loggerclient

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Logger holds the connection to rabbitMQ and handles structuring and serializing logs
type Logger struct {
	pkg string
	ch *amqp.Channel
	conn *amqp.Connection
}

func NewLogger(pkg string, connString string) (*Logger, error) {
	l := &Logger{}
	l.pkg = pkg

	conn, err := amqp.Dial(connString)
	l.conn = conn

	if err != nil {
		return l, err
	}

	ch, err := conn.Channel()
	l.ch = ch
	return l, err
}

func (l *Logger) resetChannel() error {
	var err error
	for i := 0; i < 5; i++ {
		c, err := l.conn.Channel()
		if err == nil {
			l.ch = c
			break
		}
	}
	return err
}

func (l *Logger) Log(message string, level LogLevel) error {
	log := createLog(l.pkg, message, level)
	body, err := log.serialize()
	if err != nil {
		return err
	}

	err = l.ch.Publish("", "logs", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: body,
	})
	if err != nil {
		return errors.Join(err, l.resetChannel())
	}
	return nil
}