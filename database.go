package database

//Database : Interface database object
type Database interface {
	GetConnection() interface{}
}
