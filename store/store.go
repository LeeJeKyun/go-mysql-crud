package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-mysql-crud/model"
	"log"
	"time"
)

var mysql *sql.DB

func init() {
	var err error
	mysql, err = sql.Open("mysql", "root:leejk123@tcp(localhost:3306)/GOCRUD")
	if err != nil {
		for err == nil {
			log.Default().Print("#### DB Server Open Fail ####")
			time.Sleep(5000)
			mysql, err = sql.Open("mysql", "root:leejk123@tcp(localhost:3306)/GOCRUD")
		}
	}
	log.Println("#### DB Server Open Success ####")
}

func GetDetail(key int) model.User {
	var result model.User = model.User{}
	rows, err := mysql.Query("SELECT * FROM USER where id = ?", key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Name, &result.Email)
		if err != nil {
			log.Fatal(err)
		}
	}
	return result
}

func GetAll() []model.User {
	var results []model.User
	rows, err := mysql.Query("SELECT * FROM USER")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		result := model.User{}
		err := rows.Scan(&result.Id, &result.Name, &result.Email)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	return results
}

func InsertOne(data model.User) int64 {
	result, err := mysql.Exec("INSERT INTO USER(name, email) VALUES (?, ?)", data.Name, data.Email)
	if err != nil {
		log.Fatal(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return insertId
}

func UpdateOne(data model.User) int64 {
	stmt, err := mysql.Prepare("UPDATE USER set name=?, email=? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(data.Name, data.Email, data.Id)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func DeleteOne(data model.User) int64 {
	result, err := mysql.Exec("DELETE FROM USER WHERE id = ?", data.Id)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}
