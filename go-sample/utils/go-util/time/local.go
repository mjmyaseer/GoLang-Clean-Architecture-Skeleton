package time

import (
	"go-sample/utils/go-util/config"
	"time"
)

func ToLocal(t time.Time) time.Time {
	return t.In(config.AppConf.Location)
}
