package manager

import (
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	//
	_ "github.com/go-sql-driver/mysql"
	"gitlab.pri.ibanyu.com/middleware/seaweed/internal/dbrouter"
)

const (
	defaultDriver = "mysql"
	defaultPort   = 3306
	cDSNFormat    = "%s%s=%s&"
)

//Setting is the signature of setting function
type Setting func(string) string

func boolSetting(source, param string, ok bool) string {
	return fmt.Sprintf(cDSNFormat, source, param, strconv.FormatBool(ok))
}

func timeSetting(source, param string, t time.Duration) string {
	//make sure 1ms<=t<24h
	if t < time.Millisecond || t >= 24*time.Hour {
		return ""
	}
	return fmt.Sprintf(cDSNFormat, source, param, t)
}

func stringSetting(source, param, value string) string {
	if "" == value {
		return ""
	}
	return fmt.Sprintf(cDSNFormat, source, param, value)
}

//SetCharset Sets the charset used for client-server interaction
func SetCharset(v string) Setting {
	return func(source string) string {
		return stringSetting(source, "charset", v)
	}
}

//SetLoc Sets the location for time.Time values (when using parseTime=true). "Local" sets the system's location. See time.LoadLocation for details.
func SetLoc(v string) Setting {
	return func(source string) string {
		return stringSetting(source, "loc", v)
	}
}

//SetCollation Sets the collation used for client-server interaction on connection. In contrast to charset, collation does not issue additional queries. If the specified collation is unavailable on the target server, the connection will fail.
func SetCollation(v string) Setting {
	return func(source string) string {
		return stringSetting(source, "collation", v)
	}
}

//SetAllowCleartextPasswords allowCleartextPasswords=true allows using the cleartext client side plugin if required by an account, such as one defined with the PAM authentication plugin. Sending passwords in clear text may be a security problem in some configurations.
func SetAllowCleartextPasswords(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "allowCleartextPasswords", ok)
	}
}

//SetAllowNativePasswords allows the usage of the mysql native password method
func SetAllowNativePasswords(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "allowNativePasswords", ok)
	}
}

//SetAutoCommit set it to true if you know what you are doing
func SetAutoCommit(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "autocommit", ok)
	}
}

//SetParseTime parseTime=true changes the output type of DATE and DATETIME values to time.Time instead of []byte / string
func SetParseTime(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "parseTime", ok)
	}
}

//SetAllowAllFiles allowAllFiles=true disables the file Whitelist for LOAD DATA LOCAL INFILE and allows all files. Might be insecure!
func SetAllowAllFiles(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "allowAllFiles", ok)
	}
}

//SetClientFoundRows clientFoundRows=true causes an UPDATE to return the number of matching rows instead of the number of rows changed.
func SetClientFoundRows(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "clientFoundRows", ok)
	}
}

//SetColumnsWithAlias When columnsWithAlias is true, calls to sql.Rows.Columns() will return the table alias and the column name separated by a dot.
func SetColumnsWithAlias(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "columnsWithAlias", ok)
	}
}

//SetInterpolateParams If interpolateParams is true, placeholders (?) in calls to db.Query() and db.Exec() are interpolated into a single query string with given parameters. This reduces the number of roundtrips, since the driver has to prepare a statement, execute it with given parameters and close the statement again with interpolateParams=false.
func SetInterpolateParams(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "interpolateParams", ok)
	}
}

//SetStrict strict=true enables the strict mode in which MySQL warnings are treated as errors.
func SetStrict(ok bool) Setting {
	return func(source string) string {
		return boolSetting(source, "strict", ok)
	}
}

//SetTimeout Driver side connection timeout. The value must be a decimal number with an unit suffix ( "ms", "s", "m", "h" ), such as "30s", "0.5m" or "1m30s". To set a server side timeout, use the parameter wait_timeout.
func SetTimeout(timeout time.Duration) Setting {
	return func(source string) string {
		return timeSetting(source, "timeout", timeout)
	}
}

//SetReadTimeout I/O read timeout. The value must be a decimal number with an unit suffix ( "ms", "s", "m", "h" ), such as "30s", "0.5m" or "1m30s".
func SetReadTimeout(timeout time.Duration) Setting {
	return func(source string) string {
		return timeSetting(source, "readTimeout", timeout)
	}
}

//SetWriteTimeout I/O write timeout. The value must be a decimal number with an unit suffix ( "ms", "s", "m", "h" ), such as "30s", "0.5m" or "1m30s".
func SetWriteTimeout(timeout time.Duration) Setting {
	return func(source string) string {
		return timeSetting(source, "writeTimeout", timeout)
	}
}

//Option stands for a series of options for creating a DB
type Option struct {
	driver   string
	dbName   string
	user     string
	password string
	host     string
	port     int
	settings []Setting
}

//DBInstance ...
type DBInstance struct {
	insName         string
	group           string
	dbType          string
	dbName          string
	dbAddr          string
	timeOut         time.Duration
	userName        string
	passWord        string
	db              *sql.DB
	dynamicConfiger *DynamicConfiger
}

//GetType ...
func (m *DBInstance) GetType() string {
	return m.dbType
}

//Close ...
func (m *DBInstance) Close() error {
	err := m.db.Close()
	if err != nil {
		return fmt.Errorf("xsql Instance Close err: %v", err)
	}

	return nil
}

//Reload ...
func (m *DBInstance) Reload() error {
	dynamicConf := m.dynamicConfiger.loadDynamicConf(m.insName)
	m.db.SetMaxIdleConns(dynamicConf.MaxIdleConns)
	m.db.SetMaxOpenConns(dynamicConf.MaxOpenConns)
	m.db.SetConnMaxLifetime(dynamicConf.MaxLifeTimeSec)

	return nil
}

//NewDBInstance 实例化DB实例
func NewDBInstance(insName, group string, config *dbrouter.Config, dynamicConfiger *DynamicConfiger) (*DBInstance, error) {
	dynamicConf := dynamicConfiger.loadDynamicConf(insName)
	host := strings.Split(config.DBAddr[0], ":")
	port, _ := strconv.Atoi(host[1])
	var err error
	db, err := New(config.DBName, config.UserName, config.PassWord, host[0]).Set(
		SetCharset("utf8mb4"),
		SetCollation("utf8mb4_unicode_ci"),
		SetAllowCleartextPasswords(true),
		SetInterpolateParams(true),
		SetParseTime(true),
		SetLoc(url.QueryEscape("Asia/Shanghai")),
		SetTimeout(dynamicConf.Timeout),
		SetReadTimeout(dynamicConf.ReadTimeout),
		SetWriteTimeout(dynamicConf.WriteTimeout)).Port(port).Open(true)
	if err != nil {
		fmt.Printf("new db instance error:%+v\n", err.Error())
		return nil, err
	}
	db.SetMaxIdleConns(dynamicConf.MaxIdleConns)
	db.SetMaxOpenConns(dynamicConf.MaxOpenConns)
	db.SetConnMaxLifetime(dynamicConf.MaxLifeTimeSec)
	instance := &DBInstance{
		insName:         insName,
		group:           group,
		dbType:          config.DBType,
		dbName:          config.DBName,
		dbAddr:          config.DBAddr[0],
		userName:        config.UserName,
		passWord:        config.PassWord,
		timeOut:         dynamicConf.Timeout,
		dynamicConfiger: dynamicConfiger,
		db:              db,
	}

	return instance, err
}

//New returns an Option
func New(dbName, user, password, host string) *Option {
	return &Option{
		dbName:   dbName,
		user:     user,
		password: password,
		host:     host,
		port:     defaultPort,
		driver:   defaultDriver,
	}
}

//Port sets the server port,default 3306
func (o *Option) Port(port int) *Option {
	o.port = port
	return o
}

//Driver sets the driver, default mysql
func (o *Option) Driver(driver string) *Option {
	o.driver = driver
	return o
}

//Set receives a series of Set*-like functions
func (o *Option) Set(sets ...Setting) *Option {
	o.settings = append(o.settings, sets...)
	return o
}

//Open is used for creating a *sql.DB
//Use it at the last
func (o *Option) Open(ping bool) (*sql.DB, error) {
	db, err := open(o)
	if nil != err {
		return nil, err
	}
	if ping {
		err = db.Ping()
	}
	return db, err
}

func concatDSN(settings []Setting) string {
	s := ""
	for _, f := range settings {
		s = f(s)
	}
	return strings.TrimRight(s, "&")
}

func realDSN(info *Option) string {
	format := "%s:%s@tcp(%s:%d)/%s?%s"
	return strings.TrimRight(fmt.Sprintf(format, info.user, info.password, info.host, info.port, info.dbName, concatDSN(info.settings)), "?")
}

func open(o *Option) (*sql.DB, error) {
	return sql.Open(o.driver, realDSN(o))
}
