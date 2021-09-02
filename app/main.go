package app

func RunApp(debug bool) {
	config, err := NewConfig()
	if err != nil {
		panic("failure to load config")
	}
	port := config.App.Port
	if port == "" {
		port = "8000"
	}
	NewEcho(config)
}
