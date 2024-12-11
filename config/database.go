package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DB1 *gorm.DB
	DB2 *gorm.DB
	DB3 *gorm.DB
)

func ConnectDB() {
	dbHost := "172.16.6.31"
	dbPort := "1433"
	dbUser := "sa"
	dbPass := "pass,123"

	connectToDB := func(dbName string) *gorm.DB {
		connection := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", dbHost, dbUser, dbPass, dbPort, dbName)
		db, err := gorm.Open(sqlserver.Open(connection), &gorm.Config{SkipDefaultTransaction: true})

		if err != nil {
			panic(fmt.Sprintf("Failed to connect to %s database: %s", dbName, err))
		}

		log.Printf("Connected to %s database successfully\n", dbName)
		return db
	}

	DB1 = connectToDB("Portal_Ext_Fieldcoll")
	DB2 = connectToDB("Portal_fieldcoll")
	DB3 = connectToDB("Portal_SLA_Dokumen")
}

// func ConnectDB1() {
// 	dbHost := "172.16.6.31"
// 	dbPort := "1433"
// 	dbUser := "sa"
// 	dbPass := "pass,123"
// 	dbName := "Portal_Ext_Fieldcoll"
// 	connection := "server=" + dbHost + ";user id=" + dbUser + ";password=" + dbPass + ";port=" + dbPort + ";database=" + dbName + ";encrypt=disable"

// 	db, err := gorm.Open(sqlserver.Open(connection), &gorm.Config{SkipDefaultTransaction: true})

// 	if err != nil {
// 		panic("db conection failed")
// 	}

// 	DB1 = db
// 	log.Println("db1 conected successfully")
// }

// func ConnectDB2() {
// 	dbHost := "172.16.6.31"
// 	dbPort := "1433"
// 	dbUser := "sa"
// 	dbPass := "pass,123"
// 	dbName := "Portal_fieldcoll"
// 	connection := "server=" + dbHost + ";user id=" + dbUser + ";password=" + dbPass + ";port=" + dbPort + ";database=" + dbName + ";encrypt=disable"

// 	db, err := gorm.Open(sqlserver.Open(connection), &gorm.Config{SkipDefaultTransaction: true})

// 	if err != nil {
// 		panic("db conection failed")
// 	}

// 	DB2 = db
// 	log.Println("db2 conected successfully")
// }

// func ConnectDB3() {
// 	dbHost := "172.16.6.31"
// 	dbPort := "1433"
// 	dbUser := "sa"
// 	dbPass := "pass,123"
// 	dbName := "Portal_SLA_Dokumen"
// 	connection := "server=" + dbHost + ";user id=" + dbUser + ";password=" + dbPass + ";port=" + dbPort + ";database=" + dbName + ";encrypt=disable"

// 	db, err := gorm.Open(sqlserver.Open(connection), &gorm.Config{SkipDefaultTransaction: true})

// 	if err != nil {
// 		panic("db conection failed")
// 	}

// 	DB2 = db
// 	log.Println("db2 conected successfully")
// }
