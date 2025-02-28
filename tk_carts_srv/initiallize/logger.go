package initiallize

import "github.com/Numsina/tk_carts/tk_carts_srv/logger"

var l *logger.Logger

func InitLogger() *logger.Logger {
	if l == nil {
		return logger.NewLogger()
	}
	return l
}
