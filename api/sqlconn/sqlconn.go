package sqlconn

import (
   "fmt"
   "strings"
   "database/sql"

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

func NewDbLayer(driver, dsn string, fn func(*sql.DB)) (*DbLayer, error) {
   db, err := sql.Open(driver, dsn)
   if err != nil {
	   return nil, err
   }

   if fn != nil { fn(db) }

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
	   args []interface{}
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

/*
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "widget@user", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "DESKTOP-IF9CHOJ\\SQLEXPRESS", "the database server")
	user          = flag.String("user", "widgetuser", "the database user")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")
}
*/