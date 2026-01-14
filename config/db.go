//CONFIG - SETUP DATABASE CONNECTION

package config

import (
	"log"
	"os"

	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DBUSER") + ":" +
		os.Getenv("DBPASS") + "@tcp(" +
		os.Getenv("DBADDRESS") + ")/" +
		os.Getenv("DBNAME") +
		"?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
}
