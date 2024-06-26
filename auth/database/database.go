package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Database struct {
	timeout time.Duration
	Pool    *pgxpool.Pool
}

func New(timeout time.Duration) (*Database, error) {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:password@authdb:5432/auth")
	if err != nil {
		return nil, err
	}
	for i := 0; i < 5; i++ {
		err = pool.Ping(context.Background())
		if err != nil {
			log.Println("ping failed, trying again")
			time.Sleep(timeout)
		}
	}
	if err != nil {
		return nil, err
	}
	d := &Database{}
	d.timeout = timeout
	d.Pool = pool
	return d, nil
}

func (d *Database) Login(ctx context.Context, email, password string) (int, error) {
	query := `select id, password from users where email = $1;`

	row := d.Pool.QueryRow(ctx, query, email)

	var id int
	var pass string
	if err := row.Scan(&id, &pass); err != nil {
		return 0, err
	}

	return id, bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
}

func (d *Database) Register(ctx context.Context, email, username, password string) (int, error) {
	query := `insert into users (email, username, password) values ($1, $2, $3) returning id;`

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), -1)
	if err != nil {
		return 0, err
	}

	tx, err := d.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	// _, err = tx.Exec(ctx, query, email, username, passHash)
	// if err != nil {
	// 	return err
	// }
	var id int
	if err := tx.QueryRow(ctx, query, email, username, passHash).Scan(&id); err != nil {
		return id, err
	}

	return id, tx.Commit(ctx)
}
