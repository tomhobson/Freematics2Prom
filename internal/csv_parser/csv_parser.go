package csv_parser

import (
	"fmt"
	"github.com/tomhobson/freematics2prom/internal/domain"
	"github.com/tomhobson/freematics2prom/pkg/models"
	"strconv"
	"strings"
)

type csvParser struct {
	pidToName map[string]string
}

func NewCsvParser() domain.FreematicsCSVParser {
	uc := csvParser{
		pidToName: map[string]string{
			// Standard PIDs
			"104": "Engine Load",
			"105": "Coolant Temperature",
			"10A": "Fuel Pressure",
			"10B": "Intake Manifold Pressure",
			"10C": "Engine RPM",
			"10D": "Vehicle Speed",
			"10E": "Timing Advance",
			"10F": "Intake Air Temperature",
			"110": "MAF Air Flow Rate",
			"111": "Throttle Position",
			"11f": "Run Time Since Engine Start",
			"121": "Distance Traveled with malfunction indicator lamp on",
			"12f": "Fuel Level Input",
			"131": "Distance traveled since codes cleared",
			"133": "Barometric Pressure",
			"142": "Control module voltage",
			"143": "Absolute Load Value",
			"15B": "Hybrid battery pack remaining life",
			"15C": "Engine oil temperature",
			"15E": "Engine fuel rate",

			// Freematics Specific PIDS
			"11": "UTC Date",
			"10": "UTC Time",
			"A":  "Latitude",
			"B":  "Longitude",
			"C":  "Altitude",
			"D":  "Speed",
			"E":  "Course",
			"F":  "Number of satellites in use",
			"20": "Accelerometer data",
			"21": "Gyroscope data",
			"22": "Magnitude field data",
			"23": "MEMS temperature",
			"24": "Battery voltage",
			"25": "Orientation data",
			"81": "Cellular network signal level",
			"82": "CPU temperature",
			"83": "CPU hall sensor data",
		},
	}
	return uc
}

type AccelerometerData struct {
	X, Y, Z float64
}

type GyroscopeData struct {
	X, Y, Z float64
}

func (c csvParser) ParseCSV(csvData string) ([]models.FreematicsData, error) {
	lines := strings.Split(csvData, "\n")
	dataPoints := make([]models.FreematicsData, 0)

	for _, line := range lines {
		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			continue
		}

		pid := parts[0]
		values := strings.Split(parts[1], ";")

		name, ok := c.pidToName[pid]
		if !ok {
			fmt.Printf("Unknown PID: %s\n", pid)
			name = "Unknown"
		}

		var value interface{}

		// Apply unit conversion and type conversion for specific PIDs
		switch pid {
		case "10D":
			speedKmph := convertToKmph(values[0])
			value = speedKmph
		case "10A":
			pressureKpa := convertToKpa(values[0])
			value = pressureKpa
		case "10E":
			advanceDegrees := convertToDegrees(values[0])
			value = advanceDegrees
		case "24":
			voltageV := convertToVoltage(values[0])
			value = voltageV
		case "20":
			accelerometer := parseAccelerometerData(values)
			value = accelerometer
		case "21":
			gyroscope := parseGyroscopeData(values)
			value = gyroscope
		default:
			// Use raw string value if no specific conversion is needed
			value = parts[1]
		}

		dataPoint := models.FreematicsData{
			Name:  name,
			Value: value,
		}

		dataPoints = append(dataPoints, dataPoint)
	}

	return dataPoints, nil
}

func convertToKmph(speed string) string {
	return fmt.Sprintf("%s"+" km/h", speed)
}

func convertToKpa(pressureKpa string) float64 {
	pressure, _ := strconv.ParseFloat(pressureKpa, 64)
	return pressure
}

func convertToDegrees(advanceDegrees string) float64 {
	degrees, _ := strconv.ParseFloat(advanceDegrees, 64)
	return degrees
}

func convertToVoltage(voltage string) float64 {
	volt, _ := strconv.ParseFloat(voltage, 64)
	return volt * 0.01 // 1 unit = 0.01V
}

func parseAccelerometerData(values []string) AccelerometerData {
	x, _ := strconv.ParseFloat(values[0], 64)
	y, _ := strconv.ParseFloat(values[1], 64)
	z, _ := strconv.ParseFloat(values[2], 64)
	return AccelerometerData{X: x, Y: y, Z: z}
}

func parseGyroscopeData(values []string) GyroscopeData {
	x, _ := strconv.ParseFloat(values[0], 64)
	y, _ := strconv.ParseFloat(values[1], 64)
	z, _ := strconv.ParseFloat(values[2], 64)
	return GyroscopeData{X: x, Y: y, Z: z}
}
