package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// อ่านค่า config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// ดึงค่า dsn
	dsn := viper.GetString("mysql.dsn")
	fmt.Println("DSN:", dsn)

	// เชื่อมต่อฐานข้อมูล
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// ตรวจสอบการเชื่อมต่อ
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	fmt.Println("Connection successful")
}
