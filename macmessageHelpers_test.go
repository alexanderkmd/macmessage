package macmessage

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMessageDateTimeConversion(t *testing.T) {
	date_input := 670509795321488000

	expectedStringValue := "2022-04-01 12:43:15 +0000 UTC"

	dt := convertTimestampToTime(date_input)

	assert.Equal(t, expectedStringValue, dt.UTC().String())
}

func TestTimeToMessageTimestampConversion(t *testing.T) {
	expected_output := 670509795345678912

	//time, err := time.Parse(time.RFC3339, "2022-04-01T12:43:15Z")
	ts := time.Date(2022, 4, 1, 12, 43, 15, 345678912, time.UTC)
	log.Debug(ts)
	assert.Equal(t, expected_output, convertTimeToTimestamp(ts))
}
