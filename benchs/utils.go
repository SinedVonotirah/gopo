package benchs

import (
	"os"

	"github.com/SinedVonotirah/gopo/shared/logging"
)

var (
	ORM_MULTI    int
	ORM_MAX_IDLE int
	ORM_MAX_CONN int
	ORM_SOURCE   string
)

func checkErr(err error) {
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("gorm.Open error")
		os.Exit(2)
	}
}

func wrapExecute(b *B, cbk func()) {
	b.StopTimer()
	defer b.StartTimer()
	cbk()
}
