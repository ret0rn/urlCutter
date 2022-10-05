package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ret0rn/urlCutter/models"
	"github.com/ret0rn/urlCutter/utils"
)

type Db struct {
	db *sqlx.DB
}

func NewDb() (*Db, error) {
	var dbConf models.DbConfig
	err := utils.ParseConfig("./configs/database.json", &dbConf)
	if err != nil {
		return nil, err
	}

	var dbInfo = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbConf.Db_user, dbConf.Db_password, dbConf.Db_host, dbConf.Db_port, dbConf.Db_name, dbConf.Sslmode)
	db, err := sqlx.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	db.MustExec(UrlsScheme)

	return &Db{db: db}, nil
}

var UrlsScheme = `
CREATE TABLE IF NOT EXISTS urls (
	id BIGSERIAL PRIMARY KEY,
	shorturl VARCHAR(60) NOT NULL UNIQUE,
	longurl VARCHAR NOT NULL,
	created_at timestamptz NOT NULL DEFAULT (now())
);
`
