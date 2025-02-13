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

func GetData(key int) model.User {
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
