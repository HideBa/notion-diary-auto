package app

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	db        DBConfig
	app       AppConfig
	debugMode DebugMode
}

type SampleConfig struct {
	port string
	db   string
}
type DBConfig struct {
	DBName string
	DBUser string
	DBUrl  string
	DBPass string
}

type AppConfig struct {
	Port   int
	Secret string
}

type DebugMode bool

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Failure to load .env file")
	}
	var c SampleConfig
	err = envconfig.Process("daytion", &c)
	fmt.Printf("env--------- %v", c)
	return &Config{
		db:        NewDBConfig(),
		app:       NewAppConfig(),
		debugMode: NewDebugMode(),
	}
}

func NewDBConfig() DBConfig {
	return DBConfig{
		DBName: "hoge",
		DBUser: "hoge",
		DBUrl:  "hoge",
		DBPass: "hoge",
	}
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Port:   8080,
		Secret: "secret",
	}
}

func NewDebugMode() DebugMode {
	return true
}
