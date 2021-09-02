package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB        DBConfig
	App       AppConfig
	Notion    NotionConfig
	DebugMode DebugMode `default:"true"`
}
type DBConfig struct {
	DBName string
	DBUser string
	DBUrl  string `envconfig:"DB_URL"`
	DBPass string
}

type AppConfig struct {
	Port   string
	Secret string
}

type NotionConfig struct {
	BASE_URL string `envconfig:"NOTION_BASE_URL`
}

type DebugMode bool

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("failure to load env")
		return nil, err
	} else if err == nil {
		log.Print("config: .env loaded")
	}
	var c Config
	err = envconfig.Process("dation", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &c, nil
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
		Port:   "8080",
		Secret: "secret",
	}
}

func NewDebugMode() DebugMode {
	return true
}
