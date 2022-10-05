package models

import (
	"crypto/md5"
	"encoding/hex"
)

type Url struct {
	LongUrl  string `json:"long_url" db:"longurl" binding:"required"`
	ShortUrl string `json:"short_url" db:"shorturl"`
}

func (url *Url) GetShorturl() {
	hash := md5.Sum([]byte(url.LongUrl))
	url.ShortUrl = hex.EncodeToString(hash[:])
}
