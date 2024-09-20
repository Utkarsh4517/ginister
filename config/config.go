package config

import (
	"fmt"
	"os"
)

func GenerateConfigFile(projectName string) {
	configFile := fmt.Sprintf("%s/config/config.go", projectName)
	file, _ := os.Create(configFile)
	defer file.Close()

	content := `package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
		return
	}
	DB = database
	fmt.Println("Database connection successful!")
}`

	file.WriteString(content)
	fmt.Println("Config file created successfully!")
}
