package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// DSNType ...
//ENUM(sqlite3,mysql,postgres,oracle,mssql,redis,mongodb,elasticsearch,couchbase,cassandra,tidb,clickhouse,max)
type DSNType int

type dsnInitFunc = func(dsn string, path string, name string) (string, bool, error)

type dsnAction struct {
	dsn string
	fn  dsnInitFunc
}

var dsnActions = [DSNTypeMax]*dsnAction{
	DSNTypeSqlite3: &dsnAction{
		dsn: "file:%v?cache=shared&_journal=WAL&_fk=1",
		fn:  openSqlite3,
	},
}

func OpenDSN(t DSNType, path string, name string) (string, bool, error) {
	if t < 0 || t >= DSNTypeMax {
		return "", false, fmt.Errorf("invalid dsn type:%v", t)
	}
	if t != DSNTypeSqlite3 {
		return "", false, fmt.Errorf("now not support dsn type:%v", t)
	}
	return dsnActions[t].fn(dsnActions[t].dsn, path, name)

}

func openSqlite3(dsn string, path string, name string) (string, bool, error) {
	fp := filepath.Join(path, warpDSN(name))
	dsn = fmt.Sprintf(dsn, fp)
	if FileNotExists(path) {
		//ignore mkdir error
		_ = os.MkdirAll(path, 0755)
		return dsn, false, nil
	}
	return dsn, FileNotExists(fp), nil
}
