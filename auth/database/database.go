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

func (d *Database) Ctx(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx != nil {
		return context.WithTimeout(ctx, d.timeout)
	} else {
		return context.WithTimeout(context.Background(), d.timeout)
	}
}

func (d *Database) Login(email, password string) (int, error) {
	query := `select id, password from users where email = $1;`

	ctx, cancel := d.Ctx(context.Background())
	defer cancel()

	row := d.Pool.QueryRow(ctx, query, email)

	var id int
	var pass string
	if err := row.Scan(&id, &pass); err != nil {
		return 0, err
	}

	return id, bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
}

func (d *Database) Register(email, username, password string) error {
	query := `insert into users (email, username, password) values ($1, $2, $3);`

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), -1)
	if err != nil {
		return err
	}

	ctx, cancel := d.Ctx(context.Background())
	defer cancel()

	tx, err := d.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query, email, username, passHash)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
