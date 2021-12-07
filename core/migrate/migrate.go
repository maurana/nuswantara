package migrate

import (
	"fmt"

	"github.com/maurana/nuswantara/core/config"
	mg "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getURL() (sourceURL string, databaseURL string) {
	sourceURL = "file://./migrations"
	databaseURL = fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s",
		config.Cfg().MysqlUser,
		config.Cfg().MysqlPassword,
		config.Cfg().MysqlHost,
		config.Cfg().MysqlPort,
		config.Cfg().MysqlDatabase,
	)
	return
}

func Up() error {
	m, err := mg.New(getURL())
	if err != nil {
		return err
	}
	err = m.Up()
	return ignoreErrNoChange(err)
}

func Down() error {
	m, err := mg.New(getURL())
	if err != nil {
		return err
	}
	err = m.Down()
	return ignoreErrNoChange(err)
}

func Steps(n int) error {
	m, err := mg.New(getURL())
	if err != nil {
		return err
	}
	err = m.Steps(n)
	return ignoreErrNoChange(err)
}

func Drop() error {
	m, err := mg.New(getURL())
	if err != nil {
		return err
	}
	err = m.Drop()
	return ignoreErrNoChange(err)
}

func ignoreErrNoChange(err error) error {
	if err != nil && err != mg.ErrNoChange {
		return err
	}
	return nil
}