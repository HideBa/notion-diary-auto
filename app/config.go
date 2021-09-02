package app

import (
	"log"
	"os"

	"github.com/HideBa/notion-diary-auto/infrastructure/notion"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB        DBConfig
	App       AppConfig
	Notion    notion.NotionConfig
	DebugMode bool
}
type DBConfig struct {
	Name string
	User string
	Url  string `envconfig:"DB_URL"`
	Pass string
}

type AppConfig struct {
	Port   string `envconfig:"PORT"`
	Secret string
}

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
	dationEnv := os.Getenv("DATION_ENV")
	if dationEnv != "production" {
		c.DebugMode = true
	}
	return &c, nil
}

// func NewDBConfig() DBConfig {
// 	return DBConfig{
// 		DBName: "hoge",
// 		DBUser: "hoge",
// 		DBUrl:  "hoge",
// 		DBPass: "hoge",
// 	}
// }

func NewAppConfig() AppConfig {
	return AppConfig{
		Port:   "8080",
		Secret: "secret",
	}
}

// func NewDebugMode() DebugMode {
// 	return true
// }
