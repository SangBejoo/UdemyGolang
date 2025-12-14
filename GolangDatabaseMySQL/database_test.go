package GolangDatabaseMySQL

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

// ini untuk execusi seperti insert delet update
func TestExectContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES ('subagiya', 'ayub', 'subagiya12@gmail.com', 100000, 5.0, '1999-9-9', false)"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert New Customer")
}

// ini untuk query seperti select
func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT id, name FROM customer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// cara iterasi dan menampilka data

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}
	defer rows.Close()

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at  FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=======================")
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthDate.Valid {
			fmt.Println("BirthDate :", birthDate.Time)
		}
		fmt.Println("Married :", married)
	}

}

func TestInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin11"
	password := "admin12"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' Limit 1" // ini jangan gampang sql injection
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login")
	} else {
		fmt.Println("gagal login")
	}
}

// gunakan sql dengan parameter
func TestSqlWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin';"
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1" // gunakan ini agar aman dari sql injection
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("sukses login")
	} else {
		fmt.Println("gagal login")
	}
}

func TestExcecWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; drop table user; #"
	password := "admin"

	script := "INSERT INTO user (username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert New Customer")
}

// ini untuk auto increment
func TestSqlAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "ayub@gmail.com"
	comment := "admin"

	script := "INSERT INTO comments(email, comment) values (?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert New comment with id :", insertId)
}

// gunakan prepare statement untuk menggunakan query yang sama dengan parameter berbeda
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) values (?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "ayubs" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar :" + strconv.Itoa(i) + "dan komentar"
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Succes Insert New comment with id :", insertId)
	}
}

// database transaction
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) values (?, ?)"

	// lakukan transaction
	for i := 0; i < 10; i++ {
		email := "ayubs" + strconv.Itoa(i) + "@gmail.com"
		comment := "komentar :" + strconv.Itoa(i) + "dan komentar"
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Succes Insert New comment with id :", insertId)
	}

	err = tx.Rollback() // bisa roll back dan commit
	if err != nil {
		panic(err)
	}
}
