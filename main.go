package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Event struct {
	gorm.Model
	UserId int
	Summary string
	Dtstart string
	Dtend string
	Description string
	Year int
	Month int
	Day int
}
func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "testdb"
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()

	db.AutoMigrate(&Event{})
	db.Create(&Event{UserId: 100, Summary: "hogehoge"})

	// 構造体のインスタンス化
	// eventEx := Event{}
    // 指定したIDを元にレコードを１つ引っ張ってくる
	// db.First(&eventEx)
    // もしくはwhere句っぽく
	// db.First(&eventEx, "summary=?", "test")

	// eventEx.UserId = 2
	// eventEx.Day = 19
	// db.Create(&eventEx)

	// eventsEx := []Event{}
	// db.Find(&eventsEx, "user_id=?", 2)
}