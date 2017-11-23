package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbPassword = "password"
	dbName     = "postgres"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) OpenRemote(host, port string) (err error) {
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, dbUser, dbPassword, dbName)
	p.db, err = sql.Open("postgres", info)
	if err != nil {
		return errors.Wrap(err, "can't open db connection")
	}
	return nil
}

func (p *Postgres) Open() (err error) {
	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	p.db, err = sql.Open("postgres", info)
	if err != nil {
		return errors.Wrap(err, "can't open db connection")
	}
	return nil
}

func (p *Postgres) Close() {
	p.db.Close()
}

func (p *Postgres) Exec(query string) (*Table, error) {
	out := &Table{}
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "can't run query")
	}
	defer rows.Close()
	out.Header, err = rows.Columns()
	if err != nil {
		return nil, errors.Wrap(err, "can't get result columns")
	}
	for rows.Next() {
		r := fetchRow(rows, len(out.Header))
		out.Rows = append(out.Rows, r)
	}
	return out, nil
}

func fetchRow(rows *sql.Rows, colsCnt int) []string {
	sl := make([]interface{}, colsCnt)
	for i := 0; i < colsCnt; i++ {
		sl[i] = new(sql.RawBytes)
	}
	rows.Scan(sl...)

	var out []string
	for i := range sl {
		rb, _ := sl[i].(*sql.RawBytes)
		out = append(out, string(*rb))
		*rb = nil
	}
	return out
}
