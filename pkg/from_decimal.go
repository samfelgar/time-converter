package pkg

import (
	"fmt"
	"math"
	"time"
)

func FromDecimal(timeAsFloat float64) (string, error) {
	hours := int(timeAsFloat / 1)
	minutes := math.Round(math.Mod(timeAsFloat, 1) * 60)

	duration, err := time.ParseDuration(fmt.Sprintf("%dh%dm", hours, int(minutes)))

	if err != nil {
		return "", err
	}

	return duration.String(), nil
}
