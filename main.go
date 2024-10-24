package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		err := fmt.Errorf("you must inform a time")
		log.Fatal(err)
	}

	timeArg := os.Args[1]
	timeAsFloat, err := strconv.ParseFloat(timeArg, 64)

	if err != nil {
		log.Fatal("you must inform a time as decimal")
	}

	result, err := decimalToTime(timeAsFloat)

	if err != nil {
		log.Fatal("unable to parse the informed time", err)
	}

	fmt.Printf("The decimal time %.2f in time is %s\n", timeAsFloat, result)
}

func decimalToTime(timeAsFloat float64) (string, error) {
	hours := int(timeAsFloat / 1)
	minutes := math.Round(math.Mod(timeAsFloat, 1) * 60)

	duration, err := time.ParseDuration(fmt.Sprintf("%dh%dm", hours, int(minutes)))

	if err != nil {
		return "", err
	}

	return duration.String(), nil
}
