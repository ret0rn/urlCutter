package database

import (
	"fmt"

	"github.com/ret0rn/urlCutter/models"
)

func (d *Db) AddUrl(url *models.Url) error {
	const insertQuery = `INSERT INTO urls (longurl, shorturl) VALUES ($1, $2)`
	_, err := d.db.Exec(insertQuery, url.LongUrl, url.ShortUrl)
	if err != nil {
		return fmt.Errorf("[AddUrl] - could not exec query. error: %v", err)
	}
	return nil
}

func (d *Db) CheckUrl(url *models.Url) {
	const selectQuery = `SELECT shorturl FROM urls WHERE longurl = $1`
	_ = d.db.QueryRowx(selectQuery, url.LongUrl).Scan(&url.ShortUrl)
}

func (d *Db) FindUrl(url *models.Url) {
	const selectQuery = `SELECT longurl FROM urls WHERE shorturl = $1`
	_ = d.db.QueryRowx(selectQuery, url.ShortUrl).Scan(&url.LongUrl)
}
