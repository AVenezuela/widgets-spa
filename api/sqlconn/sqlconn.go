package sqlconn

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	SQL_INSERT = "INSERT INTO %s (%s) VALUES (%s)"
	SQL_UPDATE = "UPDATE %s SET %s WHERE %s"
	SQL_DELETE = "DELETE FROM %s WHERE %s"
)

type DbRow map[string]interface{}

type DbWrapper interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type DbLayer struct {
	Db interface{}
}

func NewWidgetDB(fn func(*sql.DB)) (*DbLayer, error) {
	return NewDbLayer("mssql", "server=NBBV022082\\SQLEXPRESS;user id=userwidget;password=widget@123;encrypt=disable", fn)
}

func NewDbLayer(driver, dsn string, fn func(*sql.DB)) (*DbLayer, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if fn != nil {
		fn(db)
	}

	p := WrapDbLayer(db)
	return p, nil
}

func WrapDbLayer(db interface{}) *DbLayer {
	p := new(DbLayer)
	p.Db = db
	return p
}

func dbScan(rows *sql.Rows) DbRow {
	r := DbRow{}

	cols, _ := rows.Columns()
	c := len(cols)
	vals := make([]interface{}, c)
	valPtrs := make([]interface{}, c)

	for i := range cols {
		valPtrs[i] = &vals[i]
	}

	rows.Scan(valPtrs...)

	for i := range cols {
		if val, ok := vals[i].([]byte); ok {
			r[cols[i]] = string(val)
		} else {
			r[cols[i]] = vals[i]
		}
	}

	return r
}

func (p *DbLayer) Close() {
	if db, ok := p.Db.(*sql.DB); ok {
		db.Close()
	}
}

func (p *DbLayer) Transaction(fn func(*DbLayer) error) error {
	if db, ok := p.Db.(*sql.DB); ok {
		if tx, err := db.Begin(); err != nil {
			return err
		} else {
			if err = fn(WrapDbLayer(tx)); err != nil {
				tx.Rollback()
				return err
			} else {
				tx.Commit()
			}
		}
	}
	return nil
}

func (p *DbLayer) Exec(sql string) error {
	_, err := p.Db.(DbWrapper).Exec(sql)
	return err
}

func (p *DbLayer) Insert(table string, row DbRow) (int64, error) {
	var (
		fields []string
		values []string
		args   []interface{}
	)
	for field, value := range row {
		fields = append(fields, field)
		values = append(values, "?")
		args = append(args, value)
	}

	code := fmt.Sprintf(SQL_INSERT, table, strings.Join(fields, ", "), strings.Join(values, ", "))

	res, err := p.Db.(DbWrapper).Exec(code, args...)
	if err != nil {
		return 0, err
	}

	r, _ := res.LastInsertId()
	return r, nil
}

func (p *DbLayer) Update(table string, row DbRow, condition string, args ...interface{}) (int64, error) {
	var (
		fields []string
		values []interface{}
	)
	for field, value := range row {
		fields = append(fields, fmt.Sprintf("%s = ?", field))
		values = append(values, value)
	}
	args = append(values, args...)

	code := fmt.Sprintf(SQL_UPDATE, table, strings.Join(fields, ", "), condition)

	res, err := p.Db.(DbWrapper).Exec(code, args...)
	if err != nil {
		return 0, err
	}

	r, _ := res.RowsAffected()
	return r, nil
}

func (p *DbLayer) Delete(table, condition string, args ...interface{}) (int64, error) {
	code := fmt.Sprintf(SQL_DELETE, table, condition)

	res, err := p.Db.(DbWrapper).Exec(code, args...)
	if err != nil {
		return 0, err
	}

	r, _ := res.RowsAffected()
	return r, nil
}

func (p *DbLayer) One(code string, args ...interface{}) (DbRow, error) {
	rows, err := p.Db.(DbWrapper).Query(code, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rows.Next()

	return dbScan(rows), nil
}

func (p *DbLayer) All(code string, args ...interface{}) ([]DbRow, error) {
	rows, err := p.Db.(DbWrapper).Query(code, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := make([]DbRow, 0)

	for rows.Next() {
		r = append(r, dbScan(rows))
	}

	return r, nil
}

func (p *DbLayer) Scalar(code string, args ...interface{}) (interface{}, error) {
	rows, err := p.Db.(DbWrapper).Query(code, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rows.Next()

	var r interface{}
	if err = rows.Scan(&r); err != nil {
		return nil, err
	}

	if val, ok := r.([]byte); ok {
		return string(val), nil
	}

	return r, nil
}
