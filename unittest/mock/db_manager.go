package mock

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zhao520a1a/go-base.git/unittest/mock/db"
	"gitlab.pri.ibanyu.com/quality/dry.git/errors"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xsql/manager"
)

type DBManager struct {
	BeginFn      func(ctx context.Context) (*manager.Tx, error)
	BeginInvoked bool

	GetDBFn      func(ctx context.Context) (*manager.DB, error)
	GetDBInvoked bool
}

func (m *DBManager) Begin(ctx context.Context) (*manager.Tx, error) {
	m.BeginInvoked = true
	return m.BeginFn(ctx)
}

func (m *DBManager) GetDB(ctx context.Context) (*manager.DB, error) {
	m.GetDBInvoked = true
	return m.GetDBFn(ctx)
}

func NewDBManagerMock(cluster, table string, sdb *sql.DB) db.DBManager {
	begin := func(ctx context.Context) (*manager.Tx, error) {
		op := errors.Op("DB.Begin")
		db, err := manager.GetDB(ctx, cluster, table)
		if err != nil {
			err = fmt.Errorf("%s get db err %v", op, err)
			return nil, err
		}
		return db.Begin(ctx)
	}

	getDB := func(ctx context.Context) (db *manager.DB, err error) {
		db, err = manager.GetDB(ctx, cluster, table)
		if err != nil {
			return
		}
		db.SetSQLDB(sdb)
		return
	}

	return &DBManager{
		BeginFn: begin,
		GetDBFn: getDB,
	}
}
