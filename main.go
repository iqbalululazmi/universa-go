package main

import (
	"fmt"
	"log"
	"os"
	"universa-api/Models"
	"universa-api/Routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error
var db *gorm.DB

func getEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	// Config DB
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", getEnv("DBhost"), getEnv("DBport"), getEnv("DBuser"), getEnv("DBpassword"), getEnv("DBname"))
	fmt.Println("psqlInfo:", psqlInfo)
	db, errdb := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if errdb != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", errdb)
	}

	if errdb != nil {
		fmt.Println("Status:", errdb)
	}

	db.AutoMigrate(&Models.User{})
	pgl, err := db.DB()
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer pgl.Close()

	r := Routes.SetupRouter()
	//running
	r.Run()
}
