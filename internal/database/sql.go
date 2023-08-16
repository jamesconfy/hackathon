package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type db struct {
	conn *sql.DB
}

func (m *db) Ping() error {
	return m.conn.Ping()
}

func (m *db) Close() error {
	return m.conn.Close()
}

func (m *db) GetConn() *sql.DB {
	return m.conn
}

func New(connStr string) (*db, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &db{conn: conn}, nil
}
