package service

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/ret0rn/urlCutter/models"
)

const regexUrl = `^(http[s]?:\/\/|www\.|ftp[s]?:\/\/)[\w-]+\.\w+\S*`

func (s *UrlCutterService) CreateShortUrlHandler(ctx *gin.Context) {
	var url models.Url
	if err := ctx.ShouldBind(&url); err != nil {
		s.logger.Warnf("binding: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matched, _ := regexp.MatchString(regexUrl, url.LongUrl)
	if !matched {
		s.logger.Warn("regex: bad Url format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad Url format"})
		return
	}

	s.db.CheckUrl(&url)
	if url.ShortUrl == "" {
		url.GetShorturl()
		err := s.db.AddUrl(&url)
		if err != nil {
			s.logger.Warnf("db: %s", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, url)
}

func (s *UrlCutterService) GetLongUrlHandler(ctx *gin.Context) {
	var url models.Url
	url.ShortUrl = ctx.Param("short_url")
	if url.ShortUrl == "" {
		s.logger.Warn("short_url is empty")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "short_url is empty"})
		return
	}

	s.db.FindUrl(&url)
	if url.LongUrl == "" {
		ctx.JSON(http.StatusNotFound, "url not found")
		return
	}
	ctx.Redirect(http.StatusSeeOther, url.LongUrl)
}
