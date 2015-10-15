package evergreen

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Database models a connection to an underlying database. The struct
// provides for open, querying and closing a database connection.
type Database struct {
	// The user accessing the database.
	User string

	// The password for the user accessing the database.
	Password string

	// The name of the underlying database.
	Name string

	// The name of the diver for the underlying database.
	Driver string

	// A database connection which is an instance of `sql.DB`
	Connection *sql.DB
}

func New(user string, password string, name string, driver string) Database {
	return Database{User: user, Password: password, Name: name, Driver: driver}
}

// Opens and establishes a connection to the underlying database. Upon establishing
// a successful connection, database.Connection will be hydrated.
//
// From Golang Docs - The returned DB is safe for concurrent use by multiple
// goroutines and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to close a DB.
func (d *Database) Open() error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", d.User, d.Password, d.Name)
	connection, err := sql.Open(d.Driver, dbinfo)
	if err != nil {
		return err
	}
	// Ensure DB has valid connection
	err = connection.Ping()
	if err != nil {
		return err
	}
	d.Connection = connection
	return err
}

func (d *Database) Close() error {
	return d.Connection.Close()
}

//---------------------------
// Executing Queries
//---------------------------

// Executes a query without returning any rows.
//
// Method should be used for persisting new and updated data.
func (d *Database) Execute(q *Query) (sql.Result, error) {
	q.Compile()
	stmt, err := d.Connection.Prepare(q.SQL)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(q.Args...)
}

func (d *Database) ExecuteSQL(SQL string, args []interface{}) (sql.Result, error) {
	stmt, err := d.Connection.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(args...)
}

// Executes the supplied query and returns the resulting rows.
//
// Method should be used with select statements.
func (d *Database) Query(q *Query) (*sql.Rows, error) {
	q.Compile()
	fmt.Printf("Executing Query: %v\n Args: %v ", q.SQL, q.Args)
	stmt, err := d.Connection.Prepare(q.SQL)
	if err != nil {
		return nil, err
	}
	return stmt.Query(q.Args...)
}

func (d *Database) QuerySQL(SQL string, args []interface{}) (*sql.Rows, error) {
	stmt, err := d.Connection.Prepare(SQL)
	if err != nil {
		return nil, err
	}
	return stmt.Query(args...)
}

//---------------------------
// Performing Transactions
//---------------------------

func (d *Database) PerformTransaction(f func(*sql.Tx) (bool, error)) error {
	transaction, err := d.Connection.Begin()
	if err != nil {
		return err
	}
	success, err := f(transaction)
	if success != true {
		transaction.Rollback()
	}
	return transaction.Commit()
}

// Example of a Transaction
func (d *Database) Test() error {
	return d.PerformTransaction(func(t *sql.Tx) (bool, error) {
		stmt, err := t.Prepare("")
		if err != nil {
			return false, err
		}
		_, err = stmt.Exec()
		if err != nil {
			return false, err
		}
		return true, nil
	})
}

//---------------------------
// Helpers
//---------------------------

func DatabaseIdentifier(rows *sql.Rows) (string, error) {
	var err error
	var id string
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return "", err
		}
	}
	if id == "" {
		errStr := fmt.Sprintf("Failed to get Database identifier with error: %v\n", err)
		err = errors.New(errStr)
	}
	return id, err
}

func ObjectsFromResult(rows *sql.Rows, object interface{}) []interface{} {
	// objects = []interce{}
	// columns, err := rows.Columns()
	// if err != nil {
	// 	fmt.Printf("Failed getting database columns with error - %+v\n", err)
	// }
	// // Figure out how to dynamically create an args ...
	// if rows.Next() {
	// 	err := rows.Scan(&company.Identifier, &company.Name, &company.Funding, &company.Website, &company.Created)
	// 	if err != nil {
	// 		fmt.Printf("Failed getting database identifier with error - %+v\n", err)
	// 	}
	// }
	return nil
}
