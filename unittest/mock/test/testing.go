package test

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gitlab.pri.ibanyu.com/middleware/seaweed/xlog"

	"github.com/olebedev/config"
)

func Fixture(db *sql.DB, filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	qs := strings.Split(string(bytes), ";")
	for _, q := range qs {
		if strings.TrimSpace(q) != "" {
			_, err = db.Exec(q)
			if err != nil {
				panic(err)
			}
		}
	}
}

func SetupTestingMySQL(dbName string) *sql.DB {
	ctx := context.TODO()
	dir, _ := os.Getwd()
	xlog.Infof(ctx, "dir %s", dir)
	file, err := ioutil.ReadFile("./test/settings.yaml")
	user := "root"
	password := ""
	port := "3306"
	host := "localhost"
	if err == nil {
		yamlString := string(file)

		cfg, err := config.ParseYaml(yamlString)
		if err != nil {
			panic(err)
		}

		user = cfg.UString("database.user", user)
		host = cfg.UString("database.host", host)
		password = cfg.UString("database.password", password)
		port = cfg.UString("database.port", port)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)

	testingMySQL, err := sql.Open("mysql", dsn)
	testingMySQL.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	testingMySQL.Close()

	testingDB, err := sql.Open("mysql", fmt.Sprintf("%s%s?&charset=utf8mb4&parseTime=true", dsn, dbName))
	if err != nil {
		panic(err)
	}

	return testingDB
}
