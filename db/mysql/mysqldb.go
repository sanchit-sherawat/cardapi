package mysql

import (
	"fmt"
	"time"

	"github.com/programmer-for-good/flashcardApi/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MysqlDB struct {
	db *gorm.DB
}

func (MysqlDB) Name() string {
	return "mysql"
}

func (f MysqlDB) Reader() *gorm.DB {
	return f.db
}

func (f MysqlDB) Writer() *gorm.DB {
	return f.db
}

var mysqlDB *MysqlDB

func UseMysqlDB() {
	dataSourceName := "root:sanchit@tcp(127.0.0.1:3306)/falshcard?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dataSourceName)
	if mysqlDB == nil {
		gdb, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{NowFunc: time.Now().Local,
			NamingStrategy: schema.NamingStrategy{
				NoLowerCase: true, // skip the snake_casing of names
				// 			//SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		})
		if err != nil {
			panic(err)
		}
		mysqlDB = &MysqlDB{gdb}
	}

	// Set as global in stance
	db.Instance(mysqlDB)
}
