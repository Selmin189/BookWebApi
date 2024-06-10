package db

import (
	"database/sql"
	"fmt"
	"strings"
)

func Insert(table string, fields []string, values []interface{}) (sql.Result, error) {
	placeholders := make([]string, len(fields))
	for i := range fields {
		placeholders[i] = "?"
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(fields, ", "), strings.Join(placeholders, ", "))
	return db.Exec(query, values...)
}

func Select(table string, dest interface{}) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	return db.Query(query)
}

func Update(table string, fields []string, values []interface{}, id int64) (sql.Result, error) {
	set := make([]string, len(fields))
	for i, field := range fields {
		set[i] = fmt.Sprintf("%s = ?", field)
	}
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", table, strings.Join(set, ", "))
	values = append(values, id)
	return db.Exec(query, values...)
}

func Delete(table string, id int64) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", table)
	return db.Exec(query, id)
}
