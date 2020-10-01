package util

import (
	"context"
	"database/sql"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xsql/manager"
)

var mgr *manager.Manager

type DB struct {
	Cluster string
	Table   string
}

func NewDB(cluster string, table string) *DB {
	return &DB{Cluster: cluster, Table: table}
}

var GetDB = manager.GetDB

func (d *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	db, err := GetDB(ctx, d.Cluster, d.Table)
	if err != nil {
		return nil, err
	}
	return db.QueryContext(ctx, query, args...)
}

func (d *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	db, err := GetDB(ctx, d.Cluster, d.Table)
	if err != nil {
		var res sql.Result
		return res, err
	}
	return db.ExecContext(ctx, query, args...)
}

func (d *DB) Begin(ctx context.Context) (*manager.Tx, error) {
	db, err := GetDB(ctx, d.Cluster, d.Table)
	if err != nil {
		return nil, err
	}
	return db.Begin(ctx)
}
