package time

import (
	"time"

	"github.com/beevik/ntp"
)

func GetCurrentTime() (time.Time, error) {
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}
