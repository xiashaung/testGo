package pkg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	DB *gorm.DB
)

//type config struct {
//	user string
//	password string
//	ip string
//	port int
//	dbname string
//	maxIdleConns int
//	maxOpenConns int
//}

type dbConfig struct {
	db  interface{}
}

func init() {
	var err error
	sec,_ := cfg.GetSection("mysql")
	var str = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sec.Key("name").String(),
		sec.Key("pwd").String(),
		sec.Key("host").String(),
		sec.Key("port").String(),
		sec.Key("dbname").String(),
		)
	DB,err = gorm.Open("mysql",str)
	if err != nil {
		panic(err)
	}
	table_prefix := sec.Key("table_prefix").String();
	gorm.DefaultTableNameHandler = func (db *gorm.DB,tableName string)string{
		return table_prefix+tableName
	}
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func closeDb(){
	defer DB.Close()
}

func getMyDB() *gorm.DB {
	return DB
}

