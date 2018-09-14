package main

import(
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
	return nil
}

type User struct{
	Id int `db:"id`
	Name string `db:"name"`
	Age int `db:"age"`
}
func testQueryData(){
	sqlstr := "select id, name, age from user where id = ?"
	row := DB.QueryRow(sqlstr, 2)
	err := row.Scan(&User.Id, &User.Name, &User.Age)
	if err != nil {
		fmt.Printf("scan err:%v\n", err)
		return
	}
	fmt.Printf("ID:%d, Name:%s, Age:%d\n",User.Id,User.Name,User.Age)
}

func main()  {
	// 初始化数据库
	err := initDB()
	if err != nil{
		fmt.Printf("init db failed, err:%v\n",err)
		return
	}
	testQueryData()
}