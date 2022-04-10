package manager

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xtrace"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xlog"

	"gitlab.pri.ibanyu.com/middleware/seaweed/internal/breaker"

	"github.com/jmoiron/sqlx"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xconfig"
)

// DBX mapper dbrouter DB
type DBX struct {
	*sqlx.DB
}

var dbManager *Manager

const (
	DB_TYPE_MYSQL = "mysql"
)

func init() {
	var err error
	dbManager, err = NewManager()
	if err != nil {
		panic(err)
	}
}

// GetDB  generalization of GetDB in Manager
func GetDB(ctx context.Context, cluster, table string) (*DB, error) {
	return dbManager.GetDB(ctx, cluster, table)
}

// InitDBConf init db dynamic conf
func InitDBConf(ctx context.Context, confCenter xconfig.ConfigCenter) error {

	return dbManager.InitConf(ctx, confCenter)
}

// ReloadDBConf reload db conf when change
func ReloadDBConf(ctx context.Context, confCenter xconfig.ConfigCenter, event xconfig.ChangeEvent) error {

	return dbManager.ReloadConf(ctx, confCenter, event)
}

// SqlExec generalization of SqlExec in Manager
func SqlExec(ctx context.Context, cluster string, query func(*DBX, []interface{}) error, tables ...string) error {
	fun := "XSQL.SqlExec -->"

	span, ctx := xtrace.StartSpanFromContext(ctx, "xsql.SqlExec")
	defer span.Finish()

	if len(tables) <= 0 {
		return fmt.Errorf("tables is empty")
	}
	table := tables[0]

	span.LogFields(
		xtrace.String(spanLogKeyCluster, cluster),
		xtrace.String(spanLogKeyTable, table))

	// check breaker
	if !breaker.Entry(cluster, table) {
		xlog.Errorf(ctx, "%s trigger tidb breaker, because too many timeout sqls, cluster: %s, table: %s", fun, cluster, table)
		return errors.New("sql cause breaker, because too many timeout")
	}

	db, err := GetDB(ctx, cluster, table)
	if err != nil {
		return err
	}

	dbx := buildDBX(db.db)
	var tmptables []interface{}
	for _, item := range tables {
		tmptables = append(tmptables, item)
	}
	now := time.Now()
	err = query(dbx, tmptables)
	_metricReqDur.With("cluster", cluster, "table", table, "command", "sqlExec").Observe(float64(time.Since(now) / time.Millisecond))
	statMetricReqErrTotal(cluster, table, "sqlExec", err)
	// record breaker
	breaker.StatBreaker(cluster, table, err)

	return err
}

// NamedExecWrapper ...
func (db *DBX) NamedExecWrapper(tables []interface{}, query string, arg interface{}) (sql.Result, error) {
	query = fmt.Sprintf(query, tables...)
	return db.DB.NamedExec(query, arg)
}

// NamedQueryWrapper ...
func (db *DBX) NamedQueryWrapper(tables []interface{}, query string, arg interface{}) (*sqlx.Rows, error) {
	query = fmt.Sprintf(query, tables...)
	return db.DB.NamedQuery(query, arg)
}

// SelectWrapper ...
func (db *DBX) SelectWrapper(tables []interface{}, dest interface{}, query string, args ...interface{}) error {
	query = fmt.Sprintf(query, tables...)
	return db.DB.Select(dest, query, args...)
}

// ExecWrapper ...
func (db *DBX) ExecWrapper(tables []interface{}, query string, args ...interface{}) (sql.Result, error) {
	query = fmt.Sprintf(query, tables...)
	return db.DB.Exec(query, args...)
}

// NamedQueryWrapper ...
func (db *DBX) QueryRowxWrapper(tables []interface{}, query string, args ...interface{}) *sqlx.Row {
	query = fmt.Sprintf(query, tables...)
	return db.DB.QueryRowx(query, args...)
}

// QueryxWrapper ...
func (db *DBX) QueryxWrapper(tables []interface{}, query string, args ...interface{}) (*sqlx.Rows, error) {
	query = fmt.Sprintf(query, tables...)
	return db.DB.Queryx(query, args...)
}

// GetWrapper ...
func (db *DBX) GetWrapper(tables []interface{}, dest interface{}, query string, args ...interface{}) error {
	query = fmt.Sprintf(query, tables...)
	return db.DB.Get(dest, query, args...)
}

func buildDBX(oriDB *sql.DB) *DBX {
	return &DBX{
		DB: sqlx.NewDb(oriDB, DB_TYPE_MYSQL),
	}
}
