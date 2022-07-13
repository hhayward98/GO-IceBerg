/*MySQL Database*/

/*Intalling package*/
/* go get -u github.com/go-sql-driver/mysql */

/*Connecting to a MySQL database*/
/*

import "database/sql"
import _ "go-sql-driver/mysql"

// configure database connection 

db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=ture")

// init first connection to database and check error

err := db.Ping()

*/

/* Creating database table*/
/*
query := `
	CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`

// Executes the sql query in our database and check error
_, err := db.Exec(query)

*/

/*Inserting users*/
/*
import "time"

username := "hunter"
password := "Password1"
createdAt := time.Now()

//Inserts data into the users table
result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES(?, ?, ?)`, username, password, createdAt)
*/

/*Grab the recently created id for users*/

// userID, err := result.LastInsertId()



/*Querying tables*/
/*
var (
	id int
	username string
	password string
	createdAt time.Time
)

// Query the database and scan values into out variables.
query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAT)

*/

/*Querying all users*/
/*
type user struct {
	id 		int
	username 	string
	password 	string
	createdAT time.Time
}

rows, err := dbQuery(`SELECT id, username, password, created_at FROM users`)
defer rows.Close()

var users []user
for rows.Next() {
	var u user
	err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
	users = append(users, u)
}
err := rows.Err() 
*/

/*Deleting a user from table*/
// _, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)



/*Main*/

