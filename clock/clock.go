package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	return validate(hour, minute)
}

func validate(hour, minute int) Clock {
	for hour > 24 || minute > 59 {
		if hour > 24 {
			hour -= 24
		}
		if minute > 59 {
			minute -= 60
			hour++
		}
	}

	for hour < 0 || minute < 0 {
		if hour < 0 {
			hour += 24
		}
		if minute < 0 {
			minute += 60
			hour--
		}
	}

	if hour == 24 {
		hour = 0
	}
	return Clock{hour, minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutesToAdd int) Clock {
	clock.minute += minutesToAdd
	return validate(clock.hour, clock.minute)
}
