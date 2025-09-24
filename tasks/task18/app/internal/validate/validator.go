package validate

import (
	"fmt"
	"time"
)

func ValidateDate(date time.Time) error {
	currentDate := time.Now()
	if date.Year() < currentDate.Year() ||
		(date.Year() == currentDate.Year() && date.YearDay() < currentDate.YearDay()) {
		return fmt.Errorf("you can't create event in the past")
	}
	return nil
}
