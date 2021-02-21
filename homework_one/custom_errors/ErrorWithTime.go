package custom_errors

import (
	"fmt"
	"time"
)

type ErrorWithTime struct {
	text string
	time time.Time
}

func New(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now(),
	}
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s, time: %s", e.text, e.time)
}
