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

func TestBuildingQuery(t *testing.T) {
	query := Query{
		Action: Insert,
		Table:  "people",
		Collum: []string{"name", "age", "height"},
		Value:  []interface{}{"Kevin", "29", "6.0"},
	}
	query.Compile()
	fmt.Printf("Query: %v\n", query.SQL)
	if query.SQL != "INSERT INTO people(name, age, height) VALUES($1,$2,$3)" {
		t.Fail()
	}
}

func TestBuildingQueryWithReturnValue(t *testing.T) {
	query := Query{
		Action: Insert,
		Table:  "people",
		Collum: []string{"name", "age", "height"},
		Value:  []interface{}{"Kevin", "29", "6.0"},
		Return: "user_id",
	}
	query.Compile()
	fmt.Printf("Query: %v\n", query.SQL)
	if query.SQL != "INSERT INTO people(name, age, height) VALUES($1,$2,$3) RETURNING user_id" {
		t.Fail()
	}
}
