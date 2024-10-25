package pkg

import "time"

func ToDecimal(rawDuration string) (float64, error) {
	duration, err := time.ParseDuration(rawDuration)

	if err != nil {
		return 0, err
	}

	hours := duration.Hours()
	minutes := (duration.Minutes() - (duration.Hours() * 60)) / 60

	return hours + minutes, nil
}
