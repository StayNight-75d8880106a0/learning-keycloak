package mysql

import (
	"fmt"
	"learning-keycloak/internal/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlClient struct {
	DB *gorm.DB
}

var mySqlClient *MySqlClient

func NewConnectToMySQL() *MySqlClient {

	config := config.NewMySqlConfig()

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local`,
		config.MYSQL_USER,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_DATABASE,
		config.MYSQL_CHARSET,
	)

	db, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if errDB != nil {
		panic("Failed To Connect MySQL!: " + errDB.Error())
	}

	connectToSQLDB, errConnectSQLDB := db.DB()

	if errConnectSQLDB != nil {
		panic("Failed To Connect MySQL!: " + errConnectSQLDB.Error())
	}

	connectToSQLDB.SetMaxOpenConns(101)
	connectToSQLDB.SetMaxIdleConns(26)
	connectToSQLDB.SetConnMaxLifetime(31 * time.Minute)
	connectToSQLDB.SetConnMaxIdleTime(11 * time.Minute)

	if mySqlClient == nil {
		mySqlClient = &MySqlClient{
			DB: db,
		}
		log.Println("Success Connect To MySQL ✅🎌")
	}

	return mySqlClient

}

func NewGetInstaceMySQL() *MySqlClient {
	if mySqlClient == nil {
		panic("MySQL Client is not initialized. Please call NewConnectToMySQL first.")
	}
	return mySqlClient
}
