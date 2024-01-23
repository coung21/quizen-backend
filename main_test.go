package main

import (
	"log"
	"os"
	"quizen/config"
	"quizen/db"
	"testing"
)

func TestMain(m *testing.M) {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)

		_, err := db.Connect(
			config.MysqlUser,
			config.MysqlPassword,
			config.MysqlDb,
			config.MysqlHost,
			config.MysqlPort,
		)

		if err != nil {
			log.Fatal("cannot connect to db:", err)
		}

		os.Exit(m.Run())
	}
}
