package infrastructure

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
}

const dbTimeout = time.Second * 3

type MysqlRepository struct {
	Conn *sqlx.DB
}

func NewMysqlRepository(pool *sqlx.DB) *MysqlRepository {
	return &MysqlRepository{
		Conn: pool,
	}
}

// sqlContextGetter is an interface provided both by transaction and standard db connection
type sqlContextGetter interface {
	GetContext(ctx context.Context, dest any, query string, args ...any) error
}

func NewMysqlConnection() (*sqlx.DB, error) {
	config := mysql.NewConfig()

	config.Net = "tcp"
	config.Addr = os.Getenv("MYSQL_ADDR")
	config.User = os.Getenv("MYSQL_USER")
	config.Passwd = os.Getenv("MYSQL_PASSWORD")
	config.DBName = os.Getenv("MYSQL_DATABASE")
	config.ParseTime = true // with that parameter, we can use time.Time in mysqlHour.Hour

	db, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		return nil, errors.Join(err, errors.New("cannot connect to MySQL"))
	}

	return db, nil
}
