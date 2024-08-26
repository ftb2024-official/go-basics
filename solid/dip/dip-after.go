package dip

import "fmt"

type DBConn interface {
	Query() interface{}
}

type UsersRepo2 struct{ db DBConn }

func (repo UsersRepo2) GetUsers() []string {
	var users []string
	res := repo.db.Query()

	switch res.(type) {
	case map[string]string:
		for _, user := range res.(map[string]string) {
			users = append(users, user)
		}
		return users
	case []string:
		return res.([]string)
	}

	return []string{}
}

type MySQL2 struct{}
type PostgreSQL2 struct{}

func (db MySQL2) Query() interface{} { return []string{"alex", "john", "mike"} }
func (db PostgreSQL2) Query() interface{} {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

func Example2() {
	mySql := MySQL2{}
	myPostgres := PostgreSQL2{}
	repo1 := UsersRepo2{mySql}
	repo2 := UsersRepo2{myPostgres}

	fmt.Println(repo1.GetUsers())
	fmt.Println(repo2.GetUsers())
}
