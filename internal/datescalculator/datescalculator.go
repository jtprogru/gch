package datescalculator

import (
	"errors"
	"time"
)

var ErrIncorrectDateFormat = errors.New("Incorrect date format! Use 2006-01-02")

const defaultDateFormat = "2006-01-02"

type DatesResponse struct {
	Today      string
	PastDate   string
	FutureDate string
}

func Calc(dateInput string, daysCount int) (*DatesResponse, error) {
	inputDate, err := time.Parse(defaultDateFormat, dateInput)
	if err != nil {
		return nil, ErrIncorrectDateFormat
	}

	return &DatesResponse{
		Today:      time.Now().Format("2006-01-02"),
		PastDate:   inputDate.AddDate(0, 0, -daysCount).Format(defaultDateFormat),
		FutureDate: inputDate.AddDate(0, 0, daysCount).Format(defaultDateFormat),
	}, nil
}
