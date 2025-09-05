package utils

import (
	"api/constants"
	"time"
)

func DateNow() string {
	now := time.Now()
	return now.Format(constants.DATE_FORMAT)
}