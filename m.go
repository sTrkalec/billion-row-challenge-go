package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func main() {
	start := time.Now()
	measurement, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}

	defer measurement.Close()

	dados := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurement)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolor := strings.Index(rawData, ";")
		location := rawData[:semicolor]
		rawTemp := rawData[semicolor+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurement, ok := dados[location]

		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {

			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		dados[location] = measurement
	}

	locations := make([]string, 0, len(dados))

	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	for _, name := range locations {
		measurement := dados[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f\n", name, measurement.Min, measurement.Max, measurement.Sum/float64(measurement.Count))
	}

	fmt.Println("Tempo de execução: ", time.Since(start))
}
