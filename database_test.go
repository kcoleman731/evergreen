package evergreen

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type User struct {
	Identifier string
	FirstName  string
	LastName   string
	Email      string
}

var user = User{FirstName: "Test", LastName: "Tester", Email: "test@tester.com"}

type Place struct {
	Identifier string
	Country    string
	City       string
	Telcode    string
}

var place = Place{Country: "USA", City: "Seattle", Telcode: "206"}

var databaseUser = "kevincoleman"
var databasePassword = ""
var databaseName = "kevincoleman"
var databaseType = "postgres"
var database = Database{databaseUser, databasePassword, databaseName, databaseType, nil}

// Perform Setup Here
func TestMain(m *testing.M) {
	HydrateTestData(&database)
	os.Exit(m.Run())
}

func TestCreatingDatabase(t *testing.T) {
	testDatabase := Database{databaseUser, databasePassword, databaseName, databaseType, nil}
	if testDatabase.User != databaseUser {
		t.Fail()
	}

	if testDatabase.Password != databasePassword {
		t.Fail()
	}

	if testDatabase.Name != databaseUser {
		t.Fail()
	}

	if testDatabase.Driver != databaseType {
		t.Fail()
	}
}

func TestOpeningDatabase(t *testing.T) {
	testDatabase := Database{databaseUser, databasePassword, databaseName, databaseType, nil}
	err := testDatabase.Open()
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}
}

func TestInsertingData(t *testing.T) {
	sql := fmt.Sprint("INSERT INTO person(first_name, last_name, email) VALUES($1,$2,$3)")
	result, err := database.Execute(sql, user.FirstName, user.LastName, user.Email)
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}
	fmt.Printf("Result %v", result)
}

func TestUpdatingData(t *testing.T) {
	sql := fmt.Sprint("INSERT INTO person(first_name, last_name, email) VALUES($1,$2,$3)")
	_, err := database.Execute(sql, user.FirstName, user.LastName, user.Email)
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}

	sql = fmt.Sprintf("UPDATE person SET name = $1 WHERE database_identifer = $2")
	_, err = database.Execute(sql, "NewTest")
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}
}

func TestDeletingData(t *testing.T) {
	sql := fmt.Sprint("INSERT INTO person(first_name, last_name, email) VALUES($1,$2,$3)")
	_, err := database.Execute(sql, user.FirstName, user.LastName, user.Email)
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}

	sql = fmt.Sprintf("DELETE FROM person WHERE database_identifer = $1")
	_, err = database.Execute(sql, "NewTest")
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}
}

func TestFetchingData(t *testing.T) {
	firstName := "Tester"
	lastName := "Testy"
	email := "test@test.com"

	sql := fmt.Sprint("INSERT INTO person(first_name, last_name, email) VALUES($1,$2,$3)")
	_, err := database.Execute(sql, firstName, lastName, email)
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}

	sql = fmt.Sprintf("SELECT * FROM person")
	_, err = database.Query(sql)
	if err != nil {
		t.Errorf("Failed with error: %v", err)
	}
}

func TestCreatingDatabaseTable(t *testing.T) {
	_, err := database.Execute("CREATE TABLE test ()")
	if err != nil {
		fmt.Printf("Failed creating test table with error %v\n", err)
		return
	}
}

func TestParseFile(t *testing.T) {
	f, err := os.Open("server.sql")
	if err != nil {
		fmt.Printf("Failed loading file with err %v", err)
	}
	defer f.Close()
	//fmt.Printf("Loaded File %v", f)
	Load(f)
}

func Load(r io.Reader) {
	// scanner := &dotsql.Scanner{}
	// queries := scanner.Run(bufio.NewScanner(r))

	//fmt.Printf("Queries %v", queries)
}

//-------------------
// Fake Data Creation
//-------------------

var UserSQL = `
CREATE TABLE IF NOT EXISTS users (
	database_identifier SERIAL PRIMARY KEY
    	first_name text,
    	last_name text,
    	email text
);`
var PlaceSQL = `
CREATE TABLE IF NOT EXISTS places (
	database_identifier SERIAL PRIMARY KEY
    	country text,
    	city text NULL,
    	telcode integer
)`

func HydrateTestData(d *Database) {
	err := d.Open()
	if err != nil {
		fmt.Printf("Failed opening database connection with error %v\n", err)
		return
	}
	_, err = d.Execute(UserSQL)
	if err != nil {
		fmt.Printf("Failed creating test person table with error %v\n", err)
		return
	}

	_, err = d.Execute(PlaceSQL)
	if err != nil {
		fmt.Printf("Failed creating test place table with error %v\n", err)
		return
	}
}
