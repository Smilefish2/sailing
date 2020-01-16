package starter

import (
	"errors"
	"fmt"
	"github.com/Smilefish0/sailing/configer"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/micro/go-micro/util/log"
	"os"
	"sync"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB Global DB connection
var db *gorm.DB
var dbOnce sync.Once

func DB() *gorm.DB {

	if db == nil {
		dbOnce.Do(func() {
			var err error

			dbConfig := configer.DatabaseConfig()
			if dbConfig.GetConnection() == "mysql" {
				db, err = gorm.Open("mysql",
					fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%s&parseTime=True&loc=Local",
						dbConfig.GetUsername(),
						dbConfig.GetPassword(),
						dbConfig.GetHost(),
						dbConfig.GetPort(),
						dbConfig.GetDatabase(),
						dbConfig.GetCharset(),
					),
				)
				// DB = DB.Set("gorm:table_options", "CHARSET=utf8")
			} else if dbConfig.GetConnection() == "postgres" {
				db, err = gorm.Open("postgres",
					fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
						dbConfig.GetUsername(),
						dbConfig.GetPassword(),
						dbConfig.GetHost(),
						dbConfig.GetDatabase(),
					),
				)
			} else if dbConfig.GetConnection() == "sqlite" {
				db, err = gorm.Open("sqlite3",
					fmt.Sprintf("%v/%v", os.TempDir(),
						dbConfig.GetDatabase(),
					),
				)
			} else {
				panic(errors.New("not supported database adapter"))
			}

			if err == nil {
				if os.Getenv("APP_DEBUG") != "" {
					db.LogMode(true)
				}
			} else {
				panic(err)
			}

			log.Logf("[Init] 服务启动：%s\n", dbConfig.GetConnection())
		})
	}
	return db
}
