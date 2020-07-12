package db

import (
	"os"

	"github.com/go-easylog/el"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

//Connect DBへの接続
func Connect() *gorm.DB {
	DBMS := os.Getenv("DB_TYPE")
	USER := os.Getenv("POSTGRES_DBNAME")
	PASS := os.Getenv("POSTGRES_PASSWORD")
	DBNAME := os.Getenv("POSTGRES_DBNAME")

	CONNECT := "host=" + os.Getenv("POSTGRES_DBHOST") +
		" port=543" +
		" user=" + USER +
		" dbname=" + DBNAME +
		" password=" + PASS +
		" sslmode=disable"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		el.Warn("Output warning")
		//panic(err.Error())
	}
	//glog.Info("DB connect")

	return db
}
