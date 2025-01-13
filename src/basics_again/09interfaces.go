package basicsagain

import "fmt"

func InterfacesFunc() {
	mock := MockDB{}
	fmt.Println(GetData(&mock))
}

type Database interface {
	Query(query string) (string, error)
}

type RealDB struct{}

type MockDB struct{}

func (db *RealDB) Query(query string) (string, error) {
	return "Real Data", nil
}

func (db *MockDB) Query(quer string) (string, error) {
	return "Mock Data", nil
}

func GetData(db Database) string {
	data, _ := db.Query("SELECT * FROM users")
	return data
}
