package evergreen

import (
	"fmt"
	"os"
	"testing"
)

// Perform Setup Here
func TestMain(m *testing.M) {
	HydrateTestData(&database)
	os.Exit(m.Run())
}

func TestInsertStmt(t *testing.T) {
	query := Query{
		Action:  INSERT,
		Table:   "people",
		Collums: []string{"name", "age", "height"},
		Values:  []interface{}{"Kevin", "29", "6.0"},
	}
	query.Compile()
	fmt.Printf("Query: %v\n", query.SQL)
	if query.SQL != "INSERT INTO people (name, age, height) VALUES ($1,$2,$3)" {
		t.Fail()
	}
}

func TestInsertStmtWithReturnValue(t *testing.T) {
	query := Query{
		Action:  INSERT,
		Table:   "people",
		Collums: []string{"name", "age", "height"},
		Values:  []interface{}{"Kevin", "29", "6.0"},
		Return:  "user_id",
	}
	query.Compile()
	fmt.Printf("Query: %v\n", query.SQL)
	if query.SQL != "INSERT INTO people (name, age, height) VALUES ($1,$2,$3) RETURNING user_id" {
		t.Fail()
	}
}

func TestSelectStmt(t *testing.T) {
	query := Query{
		Action: SELECT,
		Table:  "people",
	}
	query.Compile()
	fmt.Printf("Query: %v\n", query.SQL)
	if query.SQL != "SELECT * FROM people" {
		t.Fail()
	}
}

func TestSelectStmtWithWhere(t *testing.T) {
	query := Query{
		Action: SELECT,
		Table:  "people",
		Where: map[string]interface{}{
			"name": "kevin",
		},
	}
	query.Compile()
	fmt.Printf("Query: %v\n", query.SQL)
	if query.SQL != "SELECT * FROM people WHERE name = $1" {
		t.Fail()
	}
}
