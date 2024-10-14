package pkg

import (
	"strconv"
	"time"
)

func StringToTime(s string) (time.Duration, error) {

	switch s[len(s)-1] {

	case 's':
		return getDuration(s[0:len(s)-1], time.Second)

	case 'm':
		return getDuration(s[0:len(s)-1], time.Minute)

	case 'h':
		return getDuration(s[0:len(s)-1], time.Hour)

	}

	return time.Second, nil

}

func getDuration(digit string, tm time.Duration) (time.Duration, error) {

	n, err := strconv.Atoi(digit)

	if err != nil {
		return time.Second, err
	}

	return tm * time.Duration(n), nil

}
