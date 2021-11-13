package app

import (
	"github.com/HideBa/notion-diary-auto/adapter/controller"
	"github.com/HideBa/notion-diary-auto/infrastructure/calendar"
	"github.com/HideBa/notion-diary-auto/infrastructure/database"
	"github.com/HideBa/notion-diary-auto/infrastructure/news"
	"github.com/HideBa/notion-diary-auto/infrastructure/notion"
	"github.com/HideBa/notion-diary-auto/infrastructure/router"
	"github.com/HideBa/notion-diary-auto/infrastructure/weather"
	"github.com/HideBa/notion-diary-auto/interactor"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func NewEcho(config *Config) {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	db := database.NewDB(config.DB.Url)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("fail to create db instance")
	}
	//migration
	if err := database.Migrate(driver, config.DB.MigrationVer); err != nil {
		panic("fail to connect with DB")
	}

	ngw := notion.NewNotionDiary(&config.Notion)
	wgw := weather.NewWeather(&config.Weather)
	nsgw := news.NewNews(&config.News)
	cgw := calendar.NewCalendar(&config.Calendar)
	repo := database.NewRepository(database.NewDB(config.DB.Url))

	uc := interactor.NewInteractor(&ngw, &wgw, &nsgw, &cgw, &repo.User)
	internalCon := controller.NewInternalController(uc)
	outerCon := controller.NewOuterController((*controller.NotionRawConfig)(&config.Notion))
	apiV1 := e.Group("/api/v1")
	router.Api(apiV1, internalCon, outerCon)

	if config.DebugMode == true {
		e.Logger.SetLevel(log.ERROR)
	}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))
	e.Use(NewFirebaseAuthMiddleware(&config.Auth))

	e.Logger.Fatal(e.Start(":" + (config.App.Port)))
}
