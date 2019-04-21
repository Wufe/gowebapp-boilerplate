package services

import (
	"github.com/wufe/boilerplateprj/infrastructure"
)

type HomeService interface {
	GetStatus() string
}

type homeServiceImpl struct {
	database infrastructure.DatabaseAccessor
}

func NewHomeService(database infrastructure.DatabaseAccessor) HomeService {
	return &homeServiceImpl{
		database: database,
	}
}

func (homeService *homeServiceImpl) GetStatus() string {
	return "OK"
}
