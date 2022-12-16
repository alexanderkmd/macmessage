package macmessage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// Returns date_int value from DB in time.Time format
func convertTimestampToTime(dateInput int) time.Time {
	timeFloat := float64(dateInput) / 1000000000
	// Используется дата, таймстемп которой начинается 2001-01-01, вот и добавляем разницу
	timeFloat += 978307200
	log.Debugf("%+v - %+v", dateInput, timeFloat)

	dt := time.Unix(int64(timeFloat), 0)

	log.Debugf("%v ", dt)
	return dt
}

func convertTimeToTimestamp(time time.Time) int {
	// Используется дата, таймстемп которой начинается 2001-01-01, вот и вычитаем разницу
	out := time.UnixNano() - 978307200*1000000000
	return int(out)
}
