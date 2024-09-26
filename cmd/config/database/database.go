package database 

import ("fmt"
		"os"
	
		"gorm.io/gorm"
		"gorm.io/driver/postgres"
		)

var Db *gorm.DB


func DatabaseConnection() *gorm.DB{
	var err error

	Db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("PORTDB"),os.Getenv("DB_path"))), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	return Db
}
