package main

import (
"fmt"

_ "github.com/go-sql-driver/mysql"
"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    	`db:"qr_id"`
	Username string 	`db:"appid"`
	QrType 	 string 	`db:"qr_type"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "dsxuser_test:BTg5ONP106WH7!LL@tcp(47.99.200.12:3306)/hawkeye_test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		defer Db.Close()  // 注意这行代码要写在上面err判断的下面
		return
	}

	Db = database
}

func main() {

	var person []Person
	err := Db.Select(&person, "select qr_id,appid,qr_type from tb_qrcode_wxa limit ?", 10)
	//fmt.Printf("%s",err)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	//for k,v := range(person) {
	//	fmt.Println("select succ: %s,%s", k,v)
	//	fmt.Println("s", v.Username)
	//}

	fmt.Println("select succ:", person[0].Username)
}