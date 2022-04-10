package manager

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xconfig"

	"gitlab.pri.ibanyu.com/middleware/seaweed/internal/breaker"
	"gitlab.pri.ibanyu.com/middleware/seaweed/internal/dbrouter"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xlog"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xtrace"
)

const (
	spanLogKeyCluster = "cluster"
	spanLogKeyTable   = "table"
	spanLogKeySQL     = "sql"
)

var bCheckTableName = true

// XDB should work with database/sql.DB and database/sql.Tx, 方便在生成的dao函数里使用事务Tx和非事务DB，XDB自身没有事务相关接口
type XDB interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// DB 实现了XDB接口，同时可以通过GetTx获取一个Tx句柄并进行提交
type DB struct {
	db      *sql.DB
	cluster string
	table   string
}

// Tx wrapper of sql.Tx
type Tx struct {
	tx      *sql.Tx
	cluster string
	table   string
}

// Manager 管理数据库实例和连接池
type Manager struct {
	configer  dbrouter.Configer
	instances *dbrouter.InstanceManager
	// 动态配置
	dynamicConfiger *DynamicConfiger
}

// ExecContext exec insert/update/delete and so on.
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	fun := "xsql.DB.ExecContext"
	// check breaker
	if !breaker.Entry(db.cluster, db.table) {
		xlog.Errorf(ctx, "%s trigger tidb breaker, because too many timeout sqls, cluster: %s, table: %s", fun, db.cluster, db.table)
		return nil, errors.New("sql cause breaker, because too many timeout")
	}
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, fun)
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, db.cluster),
		xtrace.String(spanLogKeyTable, db.table),
		xtrace.String(spanLogKeySQL, query))

	now := time.Now()
	res, err := db.db.ExecContext(ctx, query, args...)
	_metricReqDur.With("cluster", db.cluster, "table", db.table, "command", "exec").Observe(float64(time.Since(now) / time.Millisecond))
	// stat breaker
	breaker.StatBreaker(db.cluster, db.table, err)
	statMetricReqErrTotal(db.cluster, db.table, "exec", err)
	return res, err
}

// QueryContext executes a query that returns rows, typically a SELECT.
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	fun := "xsql.DB.QueryContext"
	// check breaker
	if !breaker.Entry(db.cluster, db.table) {
		xlog.Errorf(ctx, "%s trigger tidb breaker, because too many timeout sqls, cluster: %s, table: %s", fun, db.cluster, db.table)
		return nil, errors.New("sql cause breaker, because too many timeout")
	}
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, fun)
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, db.cluster),
		xtrace.String(spanLogKeyTable, db.table),
		xtrace.String(spanLogKeySQL, query))

	now := time.Now()
	res, err := db.db.QueryContext(ctx, query, args...)
	_metricReqDur.With("cluster", db.cluster, "table", db.table, "command", "query").Observe(float64(time.Since(now) / time.Millisecond))
	// stat breaker
	breaker.StatBreaker(db.cluster, db.table, err)
	statMetricReqErrTotal(db.cluster, db.table, "query", err)
	return res, err
}

// QueryRowContext executes a query that is expected to return at most one row.
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	fun := "xsql.DB.QueryRowContext"
	// check breaker
	if !breaker.Entry(db.cluster, db.table) {
		xlog.Errorf(ctx, "%s trigger tidb breaker, because too many timeout sqls, cluster: %s, table: %s", fun, db.cluster, db.table)
		return nil
	}
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, fun)
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, db.cluster),
		xtrace.String(spanLogKeyTable, db.table),
		xtrace.String(spanLogKeySQL, query))

	now := time.Now()
	res := db.db.QueryRowContext(ctx, query, args...)
	_metricReqDur.With("cluster", db.cluster, "table", db.table, "command", "query row").Observe(float64(time.Since(now) / time.Millisecond))
	return res
}

// GetTx 获取sql.Tx
// Deprecated: 无法进行trace、打点，建议使用下述db.Begin()
func (db *DB) GetTx() (*sql.Tx, error) {
	return db.db.Begin()
}

// GetSQLDB 获取sql.DB
// Deprecated: 无法进行trace、打点，无事务场景建议直接使用db.QueryContext等函数
func (db *DB) GetSQLDB() *sql.DB {
	return db.db
}

//SetSQLDB mock时使用
func (db *DB) SetSQLDB(outdb *sql.DB) {
	db.db = outdb
}

// Begin return Tx, wrapper of sql.Tx
func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	var err error
	tx := &Tx{cluster: db.cluster, table: db.table}
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, "xsql.DB.Begin")
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, tx.cluster),
		xtrace.String(spanLogKeyTable, tx.table))

	now := time.Now()
	tx.tx, err = db.db.Begin()
	_metricReqDur.With("cluster", tx.cluster, "table", tx.table, "command", "begin").Observe(float64(time.Since(now) / time.Millisecond))
	statMetricReqErrTotal(tx.cluster, tx.table, "begin", err)
	return tx, err
}

// ExecContext exec insert/update/delete and so on.
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	fun := "xsql.Tx.ExecContext"
	// check breaker
	if !breaker.Entry(tx.cluster, tx.table) {
		xlog.Errorf(ctx, "%s trigger tidb breaker, because too many timeout sqls, cluster: %s, table: %s", fun, tx.cluster, tx.table)
		return nil, errors.New("sql cause breaker, because too many timeout")
	}
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, fun)
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, tx.cluster),
		xtrace.String(spanLogKeyTable, tx.table),
		xtrace.String(spanLogKeySQL, query))

	now := time.Now()
	res, err := tx.tx.ExecContext(ctx, query, args...)
	_metricReqDur.With("cluster", tx.cluster, "table", tx.table, "command", "exec").Observe(float64(time.Since(now) / time.Millisecond))
	// stat breaker
	breaker.StatBreaker(tx.cluster, tx.table, err)
	statMetricReqErrTotal(tx.cluster, tx.table, "exec", err)
	return res, err
}

// QueryContext executes a query that returns rows, typically a SELECT.
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	fun := "xsql.Tx.QueryContext"
	// check breaker
	if !breaker.Entry(tx.cluster, tx.table) {
		xlog.Errorf(ctx, "%s trigger tidb breaker, because too many timeout sqls, cluster: %s, table: %s", fun, tx.cluster, tx.table)
		return nil, errors.New("sql cause breaker, because too many timeout")
	}
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, fun)
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, tx.cluster),
		xtrace.String(spanLogKeyTable, tx.table),
		xtrace.String(spanLogKeySQL, query))

	now := time.Now()
	res, err := tx.tx.QueryContext(ctx, query, args...)
	_metricReqDur.With("cluster", tx.cluster, "table", tx.table, "command", "query").Observe(float64(time.Since(now) / time.Millisecond))
	// stat breaker
	breaker.StatBreaker(tx.cluster, tx.table, err)
	statMetricReqErrTotal(tx.cluster, tx.table, "query", err)
	return res, err
}

// QueryRowContext executes a query that is expected to return at most one row.
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	fun := "xsql.Tx.QueryRowContext"
	// check breaker
	if !breaker.Entry(tx.cluster, tx.table) {
		xlog.Errorf(ctx, "%s trigger tidb breaker, because too many timeout sqls, cluster: %s, table: %s", fun, tx.cluster, tx.table)
		return nil
	}
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, fun)
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, tx.cluster),
		xtrace.String(spanLogKeyTable, tx.table),
		xtrace.String(spanLogKeySQL, query))

	now := time.Now()
	res := tx.tx.QueryRowContext(ctx, query, args...)
	_metricReqDur.With("cluster", tx.cluster, "table", tx.table, "command", "query row").Observe(float64(time.Since(now) / time.Millisecond))
	return res
}

// Commit wrapper of sql.Tx commit
func (tx *Tx) Commit(ctx context.Context) error {
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, "xsql.Tx.Commit")
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, tx.cluster),
		xtrace.String(spanLogKeyTable, tx.table))

	now := time.Now()
	err := tx.tx.Commit()
	_metricReqDur.With("cluster", tx.cluster, "table", tx.table, "command", "commit").Observe(float64(time.Since(now) / time.Millisecond))
	statMetricReqErrTotal(tx.cluster, tx.table, "commit", err)
	return err
}

// Rollback wrapper of sql.Tx rollback
func (tx *Tx) Rollback(ctx context.Context) error {
	// trace
	span, ctx := xtrace.StartSpanFromContext(ctx, "xsql.Tx.Commit")
	defer span.Finish()
	span.LogFields(xtrace.String(spanLogKeyCluster, tx.cluster),
		xtrace.String(spanLogKeyTable, tx.table))

	now := time.Now()
	err := tx.tx.Rollback()
	_metricReqDur.With("cluster", tx.cluster, "table", tx.table, "command", "rollback").Observe(float64(time.Since(now) / time.Millisecond))
	statMetricReqErrTotal(tx.cluster, tx.table, "rollback", err)
	return err
}

//NewManager 实例化DB路由对象
func NewManager() (*Manager, error) {
	var dbChangeChan = make(chan dbrouter.ConfigChange)
	var closeInsChan = make(chan dbrouter.ChangeIns)
	var reloadInsChan = make(chan dbrouter.ChangeIns)
	etcdConfiger, err := dbrouter.NewEtcdConfiger(context.TODO(), dbChangeChan)
	if err != nil {
		return nil, err
	}
	dynamicConfiger := &DynamicConfiger{
		closeInsChan:  closeInsChan,
		reloadInsChan: reloadInsChan,
	}

	factory := func(configer dbrouter.Configer, dynamicConfiger *DynamicConfiger) func(ctx context.Context, key, group string) (in dbrouter.Instancer, err error) {
		return func(ctx context.Context, key, group string) (in dbrouter.Instancer, err error) {
			return factory(ctx, key, group, configer, dynamicConfiger)
		}
	}(etcdConfiger, dynamicConfiger)

	return &Manager{
		configer:        etcdConfiger,
		dynamicConfiger: dynamicConfiger,
		instances:       dbrouter.NewInstanceManager(factory, dbChangeChan, closeInsChan, reloadInsChan, etcdConfiger.GetAllGroups(context.TODO())),
	}, nil
}

// GetDB return manager.DB without transaction
func (m *Manager) GetDB(ctx context.Context, cluster, table string) (*DB, error) {
	if bCheckTableName && !dbrouter.BValidExpress(table) {
		return nil, fmt.Errorf("forbid use table prefix")
	}
	instance, err := m.getInstance(ctx, cluster, table)
	if err != nil {
		return nil, err
	}
	db := new(DB)
	db.cluster = cluster
	db.table = table
	db.db = instance.db
	return db, nil
}

// IgnoreTableNameCheck ignore check table name rule
func (m *Manager) IgnoreTableNameCheck() {
	bCheckTableName = false
}

//获取DB实例
func (m *Manager) getInstance(ctx context.Context, cluster, table string) (dbInstance *DBInstance, err error) {
	insName := m.configer.GetInstanceName(ctx, cluster, table)
	in := m.instances.Get(ctx, generateKey(insName))
	if in == nil {
		err = fmt.Errorf("db instance not find: instance:%s", insName)
		return nil, err
	}

	dbInstance, ok := in.(*DBInstance)
	if !ok {
		err = fmt.Errorf("db instance type error: instance:%s, dbtype:%s", insName, in.GetType())

		return nil, err
	}

	return dbInstance, err
}

// InitConf ...
func (m *Manager) InitConf(ctx context.Context, confCenter xconfig.ConfigCenter) error {
	if confCenter == nil {
		return fmt.Errorf("init xsql conf err: configcenter nil")
	}
	mysqlConf := new(cfg)
	err := confCenter.UnmarshalWithNamespace(ctx, MysqlConfNamespace, mysqlConf)
	if err != nil {
		return err
	}
	m.dynamicConfiger.mysqlConf = mysqlConf
	return nil
}

// ReloadConf ...
func (m *Manager) ReloadConf(ctx context.Context, confCenter xconfig.ConfigCenter, event xconfig.ChangeEvent) error {
	if event.Namespace != MysqlConfNamespace {
		return nil
	}
	if confCenter == nil {
		return fmt.Errorf("reload xsql conf err: configcenter nil")
	}
	c := new(cfg)
	err := confCenter.UnmarshalWithNamespace(ctx, MysqlConfNamespace, c)
	if err != nil {
		return err
	}
	m.dynamicConfiger.reloadConf(c)
	var closeMap = make(map[string]bool)
	var reloadMap = make(map[string]bool)
	for k, v := range event.Changes {
		if v != nil {
			parts := strings.Split(k, KeySep)
			if len(parts) < 3 {
				return fmt.Errorf("mysql conf key:%s fmt err", k)
			}
			if bCloseConn(k) {
				closeMap[parts[1]] = true
			} else {
				reloadMap[parts[1]] = true
			}
		}
	}
	dbrouter.SendInsChange(closeMap, m.dynamicConfiger.closeInsChan)
	dbrouter.SendInsChange(reloadMap, m.dynamicConfiger.reloadInsChan)

	return nil
}

func bCloseConn(key string) bool {
	if strings.Contains(key, TimeoutMsecKey) || strings.Contains(key, ReadTimeoutMsecKey) || strings.Contains(key, WriteTimeoutMsecKey) || strings.Contains(key, MaxLifeTimeSecKey) {
		return true
	}

	return false
}
