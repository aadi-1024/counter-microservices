package database

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

// schema for cassandra instance
// Keyspace: auth
// Table users:
// - Username Text
// - Email Text PRIMARY KEY
// - Password Text

type Database struct {
	Session *gocql.Session
}

func New() (*Database, error) {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "auth"

	var session *gocql.Session
	var err error
	for i := 0; i < 5; i++ {
		session, err = cluster.CreateSession()
		if err != nil {
			log.Println(err.Error())
			time.Sleep(5 * time.Second)
		}
	}

	d := &Database{}
	d.Session = session

	return d, err
}

func (d *Database) Login(email, password string) error {
	var passHash string
	if err := d.Session.Query(`select password from users where email = $1;`, email).Scan(&passHash); err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password))
}

func (d *Database) Register(email, username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), -1)
	if err != nil {
		return err
	}
	return d.Session.Query(`insert into users (username, email, password) values ($1, $2, $3)`, username, email, hash).Exec()
}
