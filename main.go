package main

import (
	"idev-cms-service/app"
	"idev-cms-service/config"
	_ "idev-cms-service/docs"
	_ "idev-cms-service/migrations"
	"idev-cms-service/service/util/logs"
	"idev-cms-service/setup"
)

// @termsOfService  http://swagger.io/terms/
// @contact.name    the developer
// @contact.email   devmaster@idev.com
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  BearerAuth
// @name                        Authorization
// @in                          header like: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXvCj9...

func main() {

	// load config
	appConfig := config.New()

	// init log
	log := logs.New(appConfig)

	// init setup
	s := setup.New(appConfig, log)

	// migration mongo db
	s.MigrateMongoDB(appConfig)

	// setup jaeger
	closer := s.Jaeger()
	defer func() {
		if err := closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	// setup app route and start service
	conf := &app.Config{appConfig, log, s.Logger()}
	app.New(conf).RegisterRoute().Start()
}
