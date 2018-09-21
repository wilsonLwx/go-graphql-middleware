package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDB() error {
	var err error
	dsn := "root:mysql@tcp(localhost:3306)/golang_db"
	DB, err = sql.Open("mysql",dsn)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct{
	Id 	  int     `db:"id"`
	Name  sql.NullString  `db:"name"`
	Age   int     `db:"age"`
}
func testQueryData(){
	sqlstr := "select id, name, age from user where id = ?"
	row := DB.QueryRow(sqlstr, 1)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Age)
	if err != nil {
		fmt.Printf("scan err:%v\n", err)
		return
	}
	fmt.Printf("ID:%d, Name:%v, Age:%d\n",user.Id,user.Name,user.Age)
}

func testQueryMultiData(){
	sqlstr := "select id, name, age from user where id > ?"
	rows, err := DB.Query(sqlstr,2)
	defer func(){
		if rows != nil{
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	for rows.Next(){
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan err:%v\n", err)
			return
		}
		fmt.Printf("ID:%d, Name:%v, Age:%d\n",user.Id,user.Name,user.Age)
	}

}

func main()  {
	// 初始化数据库
	err := initDB()
	if err != nil{
		fmt.Printf("init db failed, err:%v\n",err)
		return
	}
	// testQueryData()
	testQueryMultiData()
}