package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct {
	Name       string `json:"name"`
	Age        uint16 `json:"age"`
	Money      int16
	Avg_grades float64
	Happiness  float64
	Hobbies    []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 18, 10000, 5.7, 6,
		[]string{"Football", "Tennis"}}
	templ, _ := template.ParseFiles("templates/home_page.html")
	templ.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)

	http.ListenAndServe(":7070", nil)
}

func dbConnection() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/*	insert, err := db.Query("INSERT INTO users(name,age) VALUES ('Bob',31)")
		if err != nil {
			panic(err)
		}
		defer insert.Close()*/

	res, err := db.Query("SELECT name,age FROM users")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}

	defer res.Close()

	fmt.Println("Connected to DB")
}

func main() {
	dbConnection()

	handleRequest()
}
