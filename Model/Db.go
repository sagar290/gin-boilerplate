package Model

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db         *gorm.DB
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Result     string
)

func init() {
	viper.SetEnvPrefix("gin-boilerplate")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/gin-boilerplate/")
	viper.AddConfigPath("$HOME/.gin-boilerplate")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// Connection := viper.GetString("DB_CONNECTION")
	Host := viper.GetString("DB_HOST")
	Port := viper.GetString("DB_PORT")
	Database := viper.GetString("DB_DATABASE")
	Username := viper.GetString("DB_USERNAME")
	Password := viper.GetString("DB_PASSWORD")

	fmt.Println(Host)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// defer Db.Close()
	// // avoid declrare but not used
	_ = Db
	// _ = Result
	Db.AutoMigrate(&User{})
	// fmt.Println("Hello world")
	var user User
	Db.First(&User{})
	fmt.Println(user.Name)

}
