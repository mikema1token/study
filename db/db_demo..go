package main

import (
	"database/sql"
	"log"
	"time"
)

func openDB() (*sql.DB,error){
	db,err:=sql.Open("mysql","root:secret@tcp(65.49.222.239:33063)/go_web?parseTime=true")
	if err!=nil{
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db,nil
}
func createUserTable()error{
	query := `CREATE TABLE users (
       id INT AUTO_INCREMENT,
       username TEXT NOT NULL,
       password TEXT NOT NULL,
       created_at DATETIME,
       PRIMARY KEY (id)
   );`
	db, err := openDB()
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec(query)
	return err
}

func insertUser()error{
	insert := `INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)`
	db, err := openDB()
	defer db.Close()
	if err != nil {
		return err
	}
	_, err = db.Exec(insert,"my","secret",time.Now())
	return err
}

func queryUser(id int)(interface{}, error){
	type Student struct {
		id int
		username string
		password string
		createdAt time.Time
	}
	db, err := openDB()
	if err != nil {
		return nil,err
	}
	defer db.Close()
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	s := Student{}
	err = db.QueryRow(query, id).Scan(&id, &s.username, &s.password, &s.createdAt)
	return s,err
}
