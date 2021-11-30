package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestDBDemo(t *testing.T) {
	db,err:=sql.Open("mysql","root:secret@tcp(65.49.222.239:33063)/go_web")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateUserTable(t *testing.T){
	err := createUserTable()
	if err != nil {
		t.Error(err)
	}
}

func TestInsertUser(t *testing.T){
	err := insertUser()
	if err != nil {
		t.Error(err)
	}
}

func TestQueryUser(t *testing.T){
	user, err := queryUser(1)
	t.Log(user,err)
}