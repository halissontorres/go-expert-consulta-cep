package config

import (
	"time"

	"github.com/halissontorres/go-expert-consulta-cep/configs/log"
)

const ConsultaCepTimeOut = 1 * time.Second

func GetLogger(p string) *log.Logger {
	return log.NewLogger(log.WithPrefix(p))
}
