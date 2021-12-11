package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	MySqlDb *gorm.DB
	Es      *elastic.Client
)
