package mock

import (
	"context"
	"database/sql"

	"github.com/zhao520a1a/go-base/test/mock/db"
)

type DBManager struct {
	BeginFn      func(ctx context.Context) (*sql.Tx, error)
	BeginInvoked bool

	GetDBFn      func(ctx context.Context) (*sql.DB, error)
	GetDBInvoked bool
}

func (m *DBManager) Begin(ctx context.Context) (*sql.Tx, error) {
	m.BeginInvoked = true
	return m.BeginFn(ctx)
}

func (m *DBManager) GetDB(ctx context.Context) (*sql.DB, error) {
	m.GetDBInvoked = true
	return m.GetDBFn(ctx)
}

func NewDBManagerMock(cluster, table string, sdb *sql.DB) db.DBManager {
	begin := func(ctx context.Context) (tx *sql.Tx, err error) {
		// TODO
		return
	}

	getDB := func(ctx context.Context) (db *sql.DB, err error) {
		// TODO
		return
	}

	return &DBManager{
		BeginFn: begin,
		GetDBFn: getDB,
	}
}
