package kv

import (
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

var conn *sqlite.Conn
var err error

func init() {

	// Open an in-memory database.
	conn, err = sqlite.OpenConn(":memory:", sqlite.OpenReadWrite)
	if err != nil {
		panic(err)
	}
	err = sqlitex.Execute(conn, "create table storage (token text,key text,value text)", &sqlitex.ExecOptions{})
	if err != nil {
		panic(err)
	}
}

func GetValue(token string, key string) string {
	stmt := conn.Prep("select value from storage where token = $token and key = $key")
	stmt.SetText("$token", token)
	stmt.SetText("$key", key)
	var hasRow bool
	if hasRow, err = stmt.Step(); hasRow {
		return stmt.GetText("value")
	}
	return ""
}

func SetValue(token string, key string, value string) error {
	stmt := conn.Prep("select value from storage where token = $token and key = $key")
	stmt.SetText("$token", token)
	stmt.SetText("$key", key)
	var hasRow bool
	if hasRow, err = stmt.Step(); hasRow {
		// 更新
		stmt.Finalize()
		stmt = conn.Prep("update storage set value = $value where token = $token and key = $key")
		stmt.SetText("$token", token)
		stmt.SetText("$key", key)
		stmt.SetText("$value", value)
		if _, err = stmt.Step(); err != nil {
			return err
		}

	} else {
		// 插入
		stmt.Finalize()
		stmt = conn.Prep("insert into storage (token,key,value) values ($token,$key,$value)")
		stmt.SetText("$token", token)
		stmt.SetText("$key", key)
		stmt.SetText("$value", value)
		stmt.Step()
		if _, err = stmt.Step(); err != nil {
			return err
		}
	}
	return nil
}
