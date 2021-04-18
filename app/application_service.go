package app

import (
	"github.com/0x000def42/microshards-go-config/utils/error_factory"
	"github.com/0x000def42/microshards-go-config/utils/logger"
)

type ApplicationService struct {
	Log logger.Logger
	Err error_factory.ErrorFactory
}
