package app

func Run(debug bool) {
	config := NewConfig()
	port := config.app.Port
	if port == 0 {
		port = 8000
	}
	NewEcho(config)
}
