package evergreen

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

var testPerson = `
CREATE TABLE IF NOT EXISTS person (
    first_name text,
    last_name text,
    email text
);`
var testPlace = `
CREATE TABLE IF NOT EXISTS place (
    country text,
    city text NULL,
    telcode integer
)`

const (
	databaseUser     = "kevincoleman"
	databasePassword = ""
	databaseName     = "kevincoleman"
)

const TestUser = databaseUser
const TestPassword = databasePassword
const TestName = databaseName
const TestDriver = "postgres"

// Perform Setup Here
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCreatingDatabase(t *testing.T) {
	database := Database{TestUser, TestPassword, TestName, TestDriver, nil}
	if database.User != TestUser {
		t.Fail()
	}
}

func TestOpeningDatabase(t *testing.T) {
	database := Database{TestUser, TestPassword, TestName, TestDriver, nil}
	err := database.Open()
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}
}

func TestCreatingDatabaseTable(t *testing.T) {
	database := Database{TestUser, TestPassword, TestName, TestDriver, nil}
	err := database.Open()
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}
	_, err = database.Execute("CREATE TABLE test ()")
	if err != nil {
		fmt.Printf("Failed creating test table with error %v\n", err)
		return
	}
}

func ParseFile() {
	f, err := os.Open("schema.sql")
	if err != nil {

	}
	defer f.Close()

	Load(f)
}

func Load(r io.Reader) {
	scanner := &Scanner{}
	queries := scanner.Run(bufio.NewScanner(r))

	fmt.Printf("%v", queries)
}

//-----------------
// Hydrate
//-----------------

func HydrateTestData(d *Database) {
	err := d.Open()
	if err != nil {
		fmt.Printf("Failed opening database connection with error %v\n", err)
		return
	}
	_, err = d.Execute(testPerson)
	if err != nil {
		fmt.Printf("Failed creating test person table with error %v\n", err)
		return
	}

	_, err = d.Execute(testPlace)
	if err != nil {
		fmt.Printf("Failed creating test place table with error %v\n", err)
		return
	}
}
