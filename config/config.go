package config

import (
	"fmt"
	"frmgol/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connDB *gorm.DB

//Connect Function
func Connect() {
	var userDatabase, passDatabase, portDatabase, hostDatabase, nameDatabase string

	userDatabase = utils.GetEnv("DB_USER")
	passDatabase = utils.GetEnv("DB_PASSWORD")
	portDatabase = utils.GetEnv("DB_PORT")
	hostDatabase = utils.GetEnv("DB_HOST")
	nameDatabase = utils.GetEnv("DB_NAME")

	dsn :=
		" host=" + hostDatabase +
			" user=" + userDatabase +
			" password=" + passDatabase +
			" dbname=" + nameDatabase +
			" port=" + portDatabase +
			" sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	} else {
		fmt.Println("Koneksi Sukses")
	}

	connDB = db
}

//GetConnection func
func GetConnection() *gorm.DB {
	return connDB
}
