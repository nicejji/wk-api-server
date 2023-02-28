package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	Id       int       `json:"id"`
	Nickname string    `json:"nickname"`
	Password string    `json:"password"`
	Bio      *string   `json:"bio"`
	Image    *string   `json:"image"`
	Created  time.Time `json:"created"`
}

var (
	db     *sql.DB
	db_err error
)

func main() {
	// connect to sql
	db_url := "dbname=wk sslmode=disable"
	db, db_err = sql.Open("postgres", db_url)
	if db_err != nil {
		log.Fatalf("Couldn't init db: %v", db_err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Couldn't ping database: %v", err)
	}
	defer db.Close()
	// setup server
	http.HandleFunc("/users/", usersHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		u := &User{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := addUser(u)
		u = getUser(id)
		b, err := json.Marshal(u)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(b))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
	if err != nil {
		users := getAllUsers()
		b, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprint(w, string(b))
		return
	}
	user := getUser(id)
	if user == nil {
		fmt.Fprint(w, "Not found!")
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, string(b))
}

func getUser(id int) *User {
	stmt, err := db.Prepare("select * from \"user\" where id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	user := User{}
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Nickname, &user.Password, &user.Bio, &user.Image, &user.Created)
	if err != nil {
		return nil
	}
	return &user
}

func addUser(u *User) int {
	var id int
	err := db.QueryRow("insert into \"user\"(nickname, password, bio, image) values ($1, $2, $3, $4) returning id", u.Nickname, u.Password, u.Bio, u.Image).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func getAllUsers() []User {
	stmt, err := db.Prepare("select * from \"user\"")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Nickname, &user.Password, &user.Bio, &user.Image, &user.Created)
		if err != nil {
			log.Fatalf("Failed to read row: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Rows error: %v", err)
	}
	return users
}
