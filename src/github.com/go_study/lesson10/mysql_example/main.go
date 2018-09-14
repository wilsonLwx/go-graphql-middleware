package main

import (
	"database/sql"
	"fmt"
 	_ "github.com/Go-SQL-Driver/MySQL"
)


var DB *sql.DB
func initDB() error {
	var err error
	dsn := "root:mysql@tcp(localhost:3306)/golang_db"
	DB,err = sql.Open("mysql",dsn)
	if err != nil{
		return err
	}
	return nil
}

type User struct{
	Id int `db:"id"`
	Name string `db:"string"`
	Age int `db:"age"`
}
func testQueryData(){
	sqlstr := "select id,name,age from user where id=?"
	row := DB.QueryRow(sqlstr,3)

	var user User
	err := row.Scan(&user.Id,&user.Name,&user.Age)
	if err != nil{
		fmt.Printf("scan failed, err:%v\n",err)
	} 
	fmt.Printf("id:%d name:%s age:%d\n",user.Id,user.Name,user.Age)
}

func main(){
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n",err)
		return
	}
	testQueryData()
}