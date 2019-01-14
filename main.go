package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserModel struct {
	Id      int    `gorm:"primary_key"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100	)"`
}

type Student struct {
	School    string
	ClassRoom string
}

func main() {
	//Create connection
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tcp?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
		fmt.Println(err)
	}
	log.Println("Connection Established")

	//Create table in database using GROM package
	db.DropTableIfExists(&UserModel{})
	db.AutoMigrate(&UserModel{})

	//Create Student table
	db.DropTableIfExists(&Student{})
	db.AutoMigrate(&Student{})

	fmt.Println("Table users created")

	//Insert data to Users table
	user1 := &UserModel{Id: 1, Name: "Yasitha", Address: "Alawwa"}
	user2 := &UserModel{Id: 2, Name: "Kelum", Address: "Colombo"}
	user3 := &UserModel{Id: 3, Name: "Hasitha", Address: "Minuwangoda"}
	user4 := &UserModel{Id: 4, Name: "Ishara", Address: "Wijerama"}
	user5 := &UserModel{Id: 5, Name: "Rehan", Address: "Borella"}

	//Insert into Student table
	student1 := &Student{School: "Sri Rahula National School", ClassRoom: "12-A"}
	db.Create(*student1)
	//Insert data to table
	db.Create(*user1)
	db.Create(*user2)
	db.Create(*user3)
	db.Create(*user4)
	db.Create(*user5)
	//Update data
	db.Find(&user1)
	user1.Address = "Warakapola"
	db.Save(&user1)

	//Delete Data from users table
	db.Table("user_models").Where("name= ?", "Yasitha").Delete(&UserModel{})
}
