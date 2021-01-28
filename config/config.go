package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Connect Function
func Connect() *gorm.DB {
	var userDatabase, passDatabase, portDatabase, hostDatabase, nameDatabase string

	userDatabase = "postgres"
	passDatabase = "password"
	portDatabase = "5432"
	hostDatabase = "localhost"
	nameDatabase = "penduduk"

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

	return db

}
