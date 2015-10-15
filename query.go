package evergreen

import (
	"fmt"
	"strconv"
)

//------------
// SQL Helpers
//------------

const (
	Insert = "INSERT INTO "
	Select = "SELECT "
	Update = "UPDATE "
	DELETE = "DELETE "
)

type Query struct {
	Action string
	Table  string
	Collum []string
	Value  []interface{}
	Return string
	sql    string
	SQL    string
	Args   []interface{}
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) Select(values []string) *Query {
	// if values > 0 {
	// 	q.SQL = fmt.Sprintf("SELECT * ")
	// } else {
	// 	q.SQL = fmt.Sprintf("SELECT * ")
	// }
	return q
}

func (q *Query) Insert(table string) *Query {
	q.SQL = fmt.Sprintf("INSERT INTO %v", table)
	return q
}

func (q *Query) From(value string) *Query {
	// return fmt.Sprintf("FROM %v", value)
	return q
}

func Collums(q *Query) string {
	collums := ToString(q.Collum, ", ")
	return "(" + collums + ")"
}

func Values(q *Query) string {
	values := ""
	length := (len(q.Value) - 1)
	for i := 0; i < length; i++ {
		values = values + "$" + strconv.Itoa(i+1) + ","
	}
	values = values + "$" + strconv.Itoa(len(q.Value))
	return " VALUES(" + values + ")"
}

func Return(q *Query) string {
	return " RETURNING " + q.Return
}

func (q *Query) Compile() *Query {
	sql := ""
	switch q.Action {
	case Insert:
		sql = Insert + q.Table
	}
	if q.Collum != nil {
		sql = sql + Collums(q)
	}

	if q.Value != nil {
		sql = sql + Values(q)
	}

	if q.Return != "" {
		sql = sql + Return(q)
	}
	q.Args = q.Value
	q.SQL = sql
	return q
}

func ToString(ary []string, split string) string {
	str := ""
	length := len(ary) - 1
	for i := 0; i < length; i++ {
		str = str + ary[i] + split
	}
	str = str + ary[length]
	return str
}
