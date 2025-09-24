package validate

import (
	"fmt"
	"time"
)

func ValidateDate(date time.Time) error {
	currentDate := time.Now()
	today := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location())
	eventDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	if eventDate.Before(today) {
		return fmt.Errorf("you can't create event in the past")
	}
	return nil
}
