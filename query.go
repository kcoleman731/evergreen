package evergreen

//------------
// SQL Helpers
//------------

type Query struct {
	SQL  string
	Args []interface{}
}

func (q *Query) Select(values []string) *Query {
	// if values. > 0 {
	// 	return fmt.Sprintf("SELECT * ")
	// } else {
	// 	return fmt.Sprintf("SELECT * ")
	// }
	return q
}

func (q *Query) From(value string) *Query {
	// return fmt.Sprintf("FROM %v", value)
	return q
}

func (q *Query) Insert(table string) *Query {
	// return fmt.Sprintf("INSERT INTO %v", m.Name)
	return q
}

func (q *Query) Collums(values []string) *Query {
	// return fmt.Sprintf("(%v)", collums)
	return q
}

func (q *Query) Values(values []interface{}) *Query {
	// return fmt.Sprintf("VALUES(%v)", values)
	return q
}

func (q *Query) Return(value string) *Query {
	// return fmt.Sprintf("RETURNING %v", value)
	return q
}
