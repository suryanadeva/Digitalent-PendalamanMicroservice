package main

import (
	"fmt"

	"github.com/suryanadeva/digitalent-microservice/auth-service/database"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/suryanadeva/digitalent-microservice/auth-service/config"
	"github.com/suryanadeva/digitalent-microservice/auth-service/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(cfg)
	}

	_, err = initDB(cfg.Database)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("DB Connection Success")
	}

	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :8003")
	log.Panic(http.ListenAndServe(":8003", router))
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(cfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&database.Auth{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
