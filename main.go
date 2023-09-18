package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

var id int
var name string
var domain string

func main() {
	var choice int
	db := connectpostgresDB()
	for {
		fmt.Println("choose\n1.insert data\n2.read data\n3.update data\n4.delete data\n5.exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			insert(db)
	    case 2:
			read(db)
		case 3:
			update(db)
		case 4:
			delete(db)
		case 5:
			os.Exit(0)

		}
	}
}

func connectpostgresDB() *sql.DB {
	connstring := "user=postgres dbname=stddb password='4686' host=localhost port=5432 sslmode=disable"
	db, err:=sql.Open("postgres",connstring)
	if err !=nil {
     fmt.Println(err)
	}
return db
}

func insert(db *sql.DB){
	fmt.Print("enter the id ")
	fmt.Scan(&id)
	fmt.Print("enter the name ")
	fmt.Scan(&name)
	fmt.Print("enter the domain ")
	fmt.Scan(&domain)
	insertIntoPostgres(db, id, name ,domain)

}
func  insertIntoPostgres(db *sql.DB, id int,name,domain string){
	_, err := db.Exec("INSERT INTO student(id,name,domain) VALUES($1,$2,$3)",id ,name,domain)
    if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println("value inserted")
	}
}

func read(db *sql.DB){
	rows,err := db.Query("SELECT * FROM student")
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("id name domain")
		for rows.Next(){
			rows.Scan(&id,&name,&domain)
			fmt.Printf("%d - %s - %s \n",id, name,domain)
		}
	}
}

func update(db *sql.DB) {
	fmt.Print("enter the id ")
	fmt.Scan(&id)
	fmt.Print("enter the name ")
	fmt.Scan(&name)
	_,err := db.Exec("UPDATE student SET name=$1 WHERE id=$2",name,id)
	if err != nil{
		fmt.Print(err)
	}else{
		fmt.Println("data updated")
	}
}


func delete(db *sql.DB) {
	fmt.Print("enter the id ")
	fmt.Scan(&id)
	_,err := db.Exec("DELETE FROM student WHERE id=$1",id)
	if err != nil{
		fmt.Print(err)
	} else {
		fmt.Println("data deleted") 
	}
}