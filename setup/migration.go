package setup

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"idev-cms-service/config"
	"idev-cms-service/repository/mongodb"

	migrate "github.com/xakep666/mongo-migrate"
)

const (
	MigrateCollName = "_migrations"
)

func (s *setup) MigrateMongoDB(appConfig *config.Config) {
	ctx := context.Background()
	mongoConfig := &mongodb.Config{
		URI:      appConfig.MongoDBEndpoint,
		DBName:   appConfig.MongoDBName,
		CollName: MigrateCollName,
		Timeout:  appConfig.MongoDBTimeout,
	}
	repo, err := mongodb.New(ctx, mongoConfig)
	s.panicIfErr(err)

	migrate.SetDatabase(repo.DB)
	migrate.SetMigrationsCollection(MigrateCollName)

	if len(os.Args) == 1 {
		if err = migrate.Up(migrate.AllAvailable); err != nil {
			s.log.Panic("Error migration up", err.Error())
			return
		}
		return
	}

	// Command for use migration
	// go run main.go new <filename>
	// go run main.go migrate
	// go run main.go rollback

	option := os.Args[1]
	switch option {
	case "new":
		s.doNewMigrateFile()
	case "migrate":
		s.doMigrate()
	case "rollback":
		s.doRollback()
	}
	os.Exit(0)
}

func (s *setup) doMigrate() {
	s.log.Println("Beginning migration ...")
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		s.panicIfErr(err)
	}
	s.log.Println("Migration done.")
}

func (s *setup) doRollback() {
	s.log.Println("Beginning rollback ...")
	if err := migrate.Down(migrate.AllAvailable); err != nil {
		s.panicIfErr(err)
	}
	s.log.Println("Rollback done.")
}

func (s *setup) doNewMigrateFile() {
	if len(os.Args) != 3 {
		s.log.Fatal("Invalid command to new migration file")
		return
	}

	folder := "migrations"
	path, err := os.Getwd()
	if err != nil {
		s.log.Fatal(err.Error())
	}

	if _, err := os.Stat(path + "/" + folder); os.IsNotExist(err) {
		_ = os.Mkdir(path+"/"+folder, 0755)
	}

	file := fmt.Sprintf("%s/%s_%s.go", folder, time.Now().Format("20060102150405"), strings.ReplaceAll(os.Args[2], "-", "_"))

	from, err := os.Open(fmt.Sprintf("%s/template.txt", folder))
	if err != nil {
		s.log.Fatal("File migration template not found !!")
	}
	defer func() {
		_ = from.Close()
	}()

	to, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		s.panicIfErr(err)
	}
	defer func() {
		_ = to.Close()
	}()

	_, err = io.Copy(to, from)
	if err != nil {
		s.panicIfErr(err)
	}
	s.log.Printf(fmt.Sprintf("%s %s\n", "New migration file created =>", file))
}
