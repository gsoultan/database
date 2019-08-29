package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //dialect for postgres
)

type postgresql struct {
	db *gorm.DB
}

//NewPostgreSQL : New Instance of MYSQL
func NewPostgreSQL(connectionString string) (Database, error) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error: Connection failed to database:", err.Error())
		return nil, err
	}
	return &postgresql{db: db}, nil
}

func (m *postgresql) GetConnection() interface{} {
	return m.db
}
