package orm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"platcont/src/database/connection"
	"platcont/src/libraries/library"
	"strconv"
	"strings"
)

type Querys struct {
	Query     string
	rowSql    *sql.Rows
	colSql    []string
	TableName string
	db        *sql.DB
	tx        *sql.Tx
}

type Config_Query struct {
	Cloud bool
}

func NewQuerys(name string) *Querys {
	return &Querys{TableName: name}
}

func (q *Querys) Connect(config ...Config_Query) *Querys {
	cloud := false
	if len(config) >= 1 {
		if config[0].Cloud {
			cloud = true
		}
	}

	if cloud {
		q.db = connection.ConnectionCloud()
	} else {
		q.db = connection.Connection()
	}
	var err error
	q.tx, err = q.db.Begin()

	if err != nil {
		fmt.Println("Error SQL:", err.Error())
	}

	return q
}

func (q *Querys) SetQuery(query string) *Querys {
	q.Query = query
	return q
}

func (q *Querys) GetQuery() string {
	return q.Query
}

func (q *Querys) Select(fields ...string) *Querys {
	if len(fields) == 0 {
		q.Query = "SELECT * FROM " + q.TableName
	} else {
		q.Query = "SELECT " + strings.Join(fields, ",") + " FROM " + q.TableName
	}

	return q
}

func (q *Querys) Where(where string, args ...interface{}) *Querys {
	q.Query += " WHERE " + where
	q.Query += args[0].(string)
	q.Query += library.InterfaceToString(args[1], true)

	return q
}

func (q *Querys) And(where string, args ...interface{}) *Querys {
	q.Query += " AND " + where
	q.Query += args[0].(string)
	if len(args) <= 2 {
		q.Query += library.InterfaceToString(args[1], true)

	} else {
		q.Query += library.InterfaceToString(args[1], args[2].(bool))
	}
	return q
}
func (q *Querys) Or(where string, args ...interface{}) *Querys {
	q.Query += " OR " + where
	q.Query += args[0].(string)
	if len(args) <= 2 {
		q.Query += library.InterfaceToString(args[1], true)

	} else {
		q.Query += library.InterfaceToString(args[1], args[2].(bool))
	}

	return q
}
func (q *Querys) Like(field string, value string) *Querys {
	q.Query += " WHERE " + field + " LIKE " + "'" + value + "'"
	return q
}
func (q *Querys) AndLike(field string, value string) *Querys {
	q.Query += " AND " + field + " LIKE " + "'" + value + "'"
	return q
}
func (q *Querys) OrLike(field string, value string) *Querys {
	q.Query += " OR " + field + " LIKE " + "'" + value + "'"
	return q
}

func (q *Querys) OrderBy(order ...interface{}) *Querys {
	q.Query += " ORDER BY " + order[0].(string)
	if len(order) > 1 && order[1].(string) == "DESC" {
		q.Query += " " + order[1].(string)
	}

	return q
}

func (q *Querys) Offset(offset int) *Querys {
	q.Query += " OFFSET " + strconv.Itoa(offset)
	return q
}
func (q *Querys) GroupBy(group string) *Querys {
	q.Query += " GROUP BY " + group
	return q
}
func (q *Querys) Having(having string, args ...interface{}) *Querys {
	q.Query += " HAVING " + having

	return q
}

func (q *Querys) AndBetween(field string, value ...interface{}) *Querys {
	query := "AND " + field + " BETWEEN"
	q.Query += fmt.Sprintf(" %s %v AND %v", query, value[0], value[1])
	return q
}
func (q *Querys) Top(top int) *Querys {
	cadena := ""
	cadena = q.Query + " LIMIT " + strconv.Itoa(top)
	q.Query = cadena
	return q
}
func (q *Querys) Limit(limit ...int) *Querys {
	if len(limit) == 2 {
		q.Query += " LIMIT " + strconv.Itoa(limit[0]) + " OFFSET " + strconv.Itoa(limit[1])
	} else if len(limit) == 1 {
		q.Query += " LIMIT " + strconv.Itoa(limit[0])
	} else {
		q.Query += " LIMIT 1"
	}

	return q
}

func (q *Querys) Distinct() *Querys {
	q.Query += " DISTINCT"
	return q
}
func (q *Querys) InnerJoin(table string, on string) *Querys {
	q.Query += " INNER JOIN " + table + " ON " + on
	return q
}
func (q *Querys) LeftJoin(table string, on string) *Querys {
	q.Query += " LEFT JOIN " + table + " ON " + on
	return q
}

func (q *Querys) RightJoin(table string, on string) *Querys {
	q.Query += " RIGHT JOIN " + table + " ON " + on
	return q
}
func (q *Querys) FullJoin(table string, on string) *Querys {
	q.Query += " FULL JOIN " + table + " ON " + on
	return q
}

func (q *Querys) Exec(config ...Config_Query) *Querys {

	cloud := false
	if len(config) >= 1 {

		if config[0].Cloud {
			cloud = true
		}

	}
	var db *sql.DB
	if cloud {
		db = connection.ConnectionCloud()
	} else {
		db = connection.Connection()
	}

	ctx := context.Background()
	err := db.PingContext(ctx)
	defer db.Close()
	if err != nil {
		fmt.Println("Error SQL:", err.Error())
	}
	// fmt.Println("query:", q.Query)
	rows, err := db.QueryContext(ctx, q.Query)
	if err != nil {
		fmt.Println("Error SQL:", err.Error())
	}

	cols, _ := rows.Columns()

	q.rowSql = rows
	q.colSql = cols

	return q
}

func (q *Querys) ExecTx() *Querys {
	ctx := context.Background()
	rows, err := q.tx.QueryContext(ctx, q.Query)
	if err != nil {
		fmt.Println("Error SQL:", err.Error())
	}
	cols, _ := rows.Columns()

	q.rowSql = rows
	q.colSql = cols

	return q
}

func (q *Querys) One() map[string]interface{} {
	m := make(map[string]interface{})
	for q.rowSql.Next() {
		columns := make([]interface{}, len(q.colSql))
		columnPointers := make([]interface{}, len(q.colSql))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := q.rowSql.Scan(columnPointers...); err != nil {
			log.Fatal("Scan93:", err)
		}
		for i, colName := range q.colSql {
			val := columnPointers[i].(*interface{})
			l := *val
			if l != nil {
				m[colName] = l
			} else {
				m[colName] = l
			}
		}
		break
	}

	defer q.rowSql.Close()
	return m
}
func (q *Querys) Text(columna string) interface{} {
	m := make(map[string]interface{})
	for q.rowSql.Next() {
		columns := make([]interface{}, len(q.colSql))
		columnPointers := make([]interface{}, len(q.colSql))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := q.rowSql.Scan(columnPointers...); err != nil {
			log.Fatal(err)
		}

		for i, colName := range q.colSql {
			val := columnPointers[i].(*interface{})

			l := *val
			if l != nil {
				m[colName] = l
			} else {
				m[colName] = l
			}

		}

		break
	}
	defer q.rowSql.Close()
	return m[columna]
}

func (q *Querys) All() []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for q.rowSql.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.

		columns := make([]interface{}, len(q.colSql))
		columnPointers := make([]interface{}, len(q.colSql))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := q.rowSql.Scan(columnPointers...); err != nil {
			log.Fatal(err)
		}

		//Crea nuestro mapa y recupera el valor de cada columna del segmento de punteros, almacenÃ¡ndolo en el mapa con el nombre de la columna como clave.
		m := make(map[string]interface{})
		for i, colName := range q.colSql {
			val := columnPointers[i].(*interface{})
			l := *val
			if l != nil {

				m[colName] = l

			} else {
				m[colName] = l
			}
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		result = append(result, m)
	}
	defer q.rowSql.Close()
	return result
}

func (q *Querys) Close() {
	q.tx.Rollback()
	q.db.Close()
}
