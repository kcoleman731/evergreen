package evergreen

import (
	"fmt"
	"strconv"
)

//------------
// SQL Helpers
//------------

const (
	INSERT = "INSERT INTO"
	SELECT = "SELECT"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

type Query struct {
	Action  string
	Table   string
	Collums []string
	Values  []interface{}
	Where   map[string]interface{}
	Return  string
	SQL     string
	Args    []interface{}
}

func NewQuery() *Query {
	return &Query{}
}

func Select(collums []string, q *Query) {
	SQL := ""
	if collums != nil {
		SQL = fmt.Sprintf("%s %s", SELECT, CollumsToString(collums))
	} else {
		SQL = fmt.Sprintf("%s *", SELECT)
	}
	q.SQL = SQL
}

func From(table string, q *Query) {
	q.SQL = fmt.Sprintf("%s FROM %s", q.SQL, table)
}

func Where(where map[string]interface{}, q *Query) {
	q.SQL = fmt.Sprintf("%s WHERE", q.SQL)
}

func Insert(table string, q *Query) {
	q.SQL = fmt.Sprintf("INSERT INTO %s", table)
}

func Collums(collums []string, q *Query) {
	q.SQL = fmt.Sprintf("%s (%s)", q.SQL, CollumsToString(collums))
}

func Values(values []interface{}, q *Query) {
	q.SQL = fmt.Sprintf("%s VALUES (%s)", q.SQL, ValuesToString(values))
}

func Return(r string, q *Query) {
	q.SQL = fmt.Sprintf("%s RETURNING %s", q.SQL, r)
}

func CollumsToString(collums []string) string {
	return ToString(collums, ", ")
}

func ValuesToString(values []interface{}) string {
	valueString := ""
	length := (len(values) - 1)
	for i := 0; i < length; i++ {
		valueString = valueString + "$" + strconv.Itoa(i+1) + ","
	}
	valueString = valueString + "$" + strconv.Itoa(len(values))
	return valueString
}

func (q *Query) Compile() *Query {
	switch q.Action {
	case INSERT:
		Insert(q.Table, q)
		if q.Collums != nil {
			Collums(q.Collums, q)
		}
		if q.Values != nil {
			Values(q.Values, q)
		}
		if q.Return != "" {
			Return(q.Return, q)
		}
	case SELECT:
		Select(q.Collums, q)
		From(q.Table, q)
		if q.Where != nil {
			Where(q.Where, q)
		}
	}
	q.Args = q.Values
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
