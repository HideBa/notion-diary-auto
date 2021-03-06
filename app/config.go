package app

import (
	"log"
	"os"

	"github.com/HideBa/notion-diary-auto/infrastructure/calendar"
	"github.com/HideBa/notion-diary-auto/infrastructure/news"
	"github.com/HideBa/notion-diary-auto/infrastructure/notion"
	"github.com/HideBa/notion-diary-auto/infrastructure/weather"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB        DBConfig
	App       AppConfig
	Notion    notion.NotionRawConfig
	Weather   weather.WeatherConfig
	News      news.NewsConfig
	Auth      AuthConfig
	Calendar  calendar.GoogleCalendarConfig `envconfig:"GOOGLE"`
	DebugMode bool
}
type DBConfig struct {
	Name         string
	User         string
	Url          string `envconfig:"URL"`
	Pass         string
	MigrationVer uint
}

type AppConfig struct {
	Port   string `envconfig:"PORT"`
	Secret string
}

const DATABASE_MIGRATION_VERSION uint = 1

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
	c.DB.MigrationVer = DATABASE_MIGRATION_VERSION
	dationEnv := os.Getenv("DATION_ENV")
	if dationEnv != "production" {
		c.DebugMode = true
	}
	return &c, nil
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Port:   "8080",
		Secret: "secret",
	}
}
