package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// DSNType ...
// ENUM(sqlite3,mysql,postgres,oracle,mssql,redis,mongodb,elasticsearch,couchbase,cassandra,tidb,clickhouse,max)
type DSNType int

type DSNAction interface {
	DriverName() string
	DSN() string
	Func() dsnFunc
}

type dsnFunc = func(driverName string, path string, name string, debug bool) (string, bool, error)

type action struct {
	dsn        string
	dsnFunc    dsnFunc
	driverName string
}

var actions = [DSNTypeMax]*action{
	DSNTypeSqlite3: &action{
		driverName: DSNTypeSqlite3.String(),
		dsn:        "file:%v?cache=shared&_journal=WAL&_fk=1",
		dsnFunc:    openSqlite3,
	},
	DSNTypeMysql: &action{
		dsnFunc: func(dsn string, path string, name string, debug bool) (string, bool, error) {
			panic("not implemented")
		},
	},
}

func OpenDSN(t DSNType, path string, name string, debug bool) (string, bool, error) {
	if t < 0 || t >= DSNTypeMax {
		return "", false, fmt.Errorf("invalid dsn type:%v", t)
	}
	if t != DSNTypeSqlite3 {
		return "", false, fmt.Errorf("now not support dsn type:%v", t)
	}
	return actions[t].dsnFunc(actions[t].dsn, path, name, debug)
}

func openSqlite3(dsn string, path string, name string, debug bool) (string, bool, error) {
	fp := filepath.Join(path, dsnSharness(name))
	dsn = fmt.Sprintf(dsn, filepath.ToSlash(fp))
	if FileNotExists(path) {
		_ = os.MkdirAll(path, 0755)
		//return dsn, false, nil
	}
	if debug {
		log.Println("[DBC] check file exists", "filepath", fp, "exist", !FileNotExists(fp))
	}
	return dsn, !FileNotExists(fp), nil
}
