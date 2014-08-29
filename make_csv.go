package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"encoding/csv"
	"os"
    "time"
    "strings"
)

var (
	//db         *sql.DB
	dbhostsip  = "127.0.0.1:3306"
	dbusername = "root"
	dbpassowrd = "123456"
	dbname     = "nodejs"
)

func mysql_open() *sql.DB {
	Odb, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)

	if err != nil {
		log.Fatalln(err)
		fmt.Println("链接失败")
		return nil
	}
	//defer Odb.Close()
	fmt.Println("链接成功")
	return Odb
}

func mysql_select(db *sql.DB, data string) {
	rows, err := db.Query("select id,name from user order by id asc")
	if err != nil {
		log.Fatalln(err)
	}
    time := time.Now().Format("20060102")
    str := []string{"./csv/",time,".xls"}
    s := strings.Join(str,"")
    f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	
	w.Write([]string{"编号", "姓名"})
	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		w.Write([]string{id, name})
		fmt.Println("content:", id, name)
	}
	w.Flush()
}

func main() {
	db := mysql_open()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	mysql_select(db, "select * from user")
}
