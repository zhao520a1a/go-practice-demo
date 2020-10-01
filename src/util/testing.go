package util

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/olebedev/config"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xsql/manager"
	"io/ioutil"
	"strings"
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
	file, err := ioutil.ReadFile("../../settings.yaml")
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

		user = cfg.UString("development.database.user", user)
		host = cfg.UString("development.database.host", host)
		password = cfg.UString("development.database.password", password)
		port = cfg.UString("development.database.port", port)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)

	testingMySQL, err := sql.Open("mysql", dsn)
	testingMySQL.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	testingMySQL.Close()

	testingDB, err := sql.Open("mysql", fmt.Sprintf("%s%s", dsn, dbName))
	if err != nil {
		panic(err)
	}

	GetDB = func(ctx context.Context, cluster string, table string) (*manager.DB, error) {
		db := &manager.DB{}
		db.SetSQLDB(testingDB)
		return db, nil
	}
	return testingDB
}
