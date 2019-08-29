package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //dialect of mysql
)

type mysql struct {
	db *gorm.DB
}

//NewMySQL : New Instance of MYSQL
func NewMySQL(connectionString string) (Database, error) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Error: Connection failed to database:", err.Error())
		return nil, err
	}
	return &mysql{db: db}, nil
}

func (m *mysql) GetConnection() interface{} {
	return m.db
}
