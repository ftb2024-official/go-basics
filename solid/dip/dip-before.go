package dip

import "fmt"

type MySQL struct{}
type PostgreSQL struct{}

func (db MySQL) QueryMySQL() []string { return []string{"alex", "john", "mike"} }
func (db PostgreSQL) QueryPostgreSQL() map[string]string {
	return map[string]string{
		"a3f69c2b-d153-48fd-b10c-5b641657477b": "alex",
		"a4f69c2b-d153-48fd-b10c-5b641657477a": "john",
		"a5f69c2b-d153-48fd-b10c-5b641657477c": "mike",
	}
}

type UsersRepo struct {
	db MySQL
	// db PostgreSQL
}

func (repo UsersRepo) GetUsers() []string {
	res := repo.db.QueryMySQL()
	return res
	// res := repo.db.QueryPostgreSQL()
	// var users []string
	// for _, user := range res {
	// 	users = append(users, user)
	// }
	// return users
}

func Example1() {
	mySqlDB := MySQL{}
	// myPostgresDB := PostgreSQL{}
	repo := UsersRepo{mySqlDB}
	// repo := UsersRepo{myPostgresDB}

	fmt.Println(repo.GetUsers())
}
