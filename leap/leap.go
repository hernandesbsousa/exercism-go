package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	isLeap := false
	if year%4 == 0 {
		isLeap = true
		if year%100 == 0 {
			isLeap = false
			if year%400 == 0 {
				isLeap = true
			}
		}
	}
	return isLeap
}
