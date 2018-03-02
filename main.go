package main

import (
	"GoStarter/config"
	"GoStarter/controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	// Disable Console Color
	//gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	config := config.FromFile("./config/config.json")

	c := controller.New(config)
	logrus.Fatal(c.Start())
}

func initDB(config config.Main) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", config.DbConnURL())
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	if config.Debug.DB {
		db = db.Debug()
	}
	/*
		migrations := &migrate.FileMigrationSource{
			Dir: "db/migrations",
		}
		if _, err := migrate.Exec(db.DB(), "mysql", migrations, migrate.Up); err != nil {
			return nil, err
		}
	*/
	db.DB().SetMaxOpenConns(100)
	return db, nil
}
