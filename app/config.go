package app

type Config struct {
	db        DBConfig
	app       AppConfig
	debugMode DebugMode
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
