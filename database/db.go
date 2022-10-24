package database

import (
		"fmt" 
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
	)

var Db *gorm.DB
func DatabaseInit() *gorm.DB  {

	Db = connectDB()
	return Db

	
	
}

func connectDB() (*gorm.DB) {
	dsn  := "root:@tcp(127.0.0.1:3306)/go_shopcart?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn ), &gorm.Config{})

  if err !=nil {

	fmt.Println("Error...")

	return nil

}

return db
}

func InitDb(){
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/go_shopcart?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot Connect to database")
	}

	fmt.Println("Connect Database...")



}