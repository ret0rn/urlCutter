package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ret0rn/urlCutter/database"
	"github.com/ret0rn/urlCutter/models"
	"github.com/ret0rn/urlCutter/utils"
	"github.com/sirupsen/logrus"
)

type UrlCutterService struct {
	config *models.ServiceConfig
	logger *logrus.Logger
	db     *database.Db
}

func NewService() (*UrlCutterService, error) {
	var config models.ServiceConfig
	err := utils.ParseConfig("./configs/service.json", &config)
	if err != nil {
		return nil, fmt.Errorf("[NewService] - cant parse config; error: %s", err.Error())
	}
	logger, err := utils.NewLogger(logrus.DebugLevel)
	if err != nil {
		return nil, fmt.Errorf("[NewService] - cant create logger; error: %s", err.Error())
	}
	db, err := database.NewDb()
	if err != nil {
		return nil, fmt.Errorf("[NewService] - cant connect to database; error: %s", err.Error())
	}
	return &UrlCutterService{
		config: &config,
		logger: logger,
		db:     db,
	}, nil
}

func (s *UrlCutterService) Run() error {
	eng := gin.Default()

	eng.POST("/generate", s.CreateShortUrlHandler)
	eng.GET("/:short_url", s.GetLongUrlHandler)

	return eng.Run(s.config.Port)
}
