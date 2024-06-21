package database

import (
	"context"
	"counterproto"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//using the same postgres instance but the auth and counter microservices won't have access
//to each other's data so it's the same anyways

type Database struct {
	timeout time.Duration
	pool    *pgxpool.Pool
	cache   *Cache
}

func New(timeout time.Duration) (*Database, error) {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:password@counterdb:5433/counter")
	if err != nil {
		return nil, err
	}
	cache, err := NewCache()
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
	d.pool = pool
	d.cache = cache
	return d, nil
}

func (d *Database) CreateUser(ctx context.Context, uid, value int) error {
	query := `insert into data values ($1, $2);`
	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query, uid, value)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (d *Database) GetValue(ctx context.Context, uid int) (int, error) {
	cacheVal, err := d.cache.Get(uid)
	if err == nil {
		return cacheVal, nil
	}

	query := `select value from data where userid = $1;`

	row := d.pool.QueryRow(ctx, query, uid)
	var val int

	err = row.Scan(&val)
	if err != nil {
		err = d.cache.Set(uid, val)
	}
	return val, err
}

func (d *Database) UpdateValue(ctx context.Context, uid, value int, action counterproto.Action) (int, error) {
	var query string
	switch action {
	case counterproto.Action_Decrement:
		query = `update data set value = value - $1 where userid = $2 returning value;`
	case counterproto.Action_Increment:
		query = `update data set value = value + $1 where userid = $2 returning value;`
	case counterproto.Action_SetValue:
		query = `update data set value = $1 where userid = $2 returning value;`
	}

	tx, err := d.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, err
	}

	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, query, value, uid)
	var val int

	if err := row.Scan(&val); err != nil {
		return val, err
	}

	_ = d.cache.Set(uid, val)
	return val, tx.Commit(ctx)
}
