package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	month, year, units, money, monthlyIncome, totalInvested, lines := InitParams()

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		currentMonth, currentYear, currentPrice, err := ParseLine(lines[i])
		if err != nil {
			break
		}
		if month != currentMonth || year != currentYear {
			month = currentMonth
			year = currentYear
			money, units, totalInvested = UpdateParams(money, monthlyIncome, currentPrice, units, totalInvested)
			fmt.Println("current year:", currentYear, "current month:", currentMonth, "money:", money, "units:",
				units, "current price:", currentPrice, "cumulative worth:", Floor(currentPrice*float64(units)), "total invested:", Floor(totalInvested))
		}
	}
}

func Floor(value float64) int {
	return int(math.Floor(value))
}

// UpdateParams TODO: should support changing monthly income
func UpdateParams(money float64, monthlyIncome float64, currentPrice float64, units int, totalInvested float64) (float64, int, float64) {
	money += monthlyIncome
	unitsToBuy := int(math.Floor(money / currentPrice))
	units += unitsToBuy
	money -= float64(unitsToBuy) * currentPrice
	return money, units, totalInvested + monthlyIncome
}

func InitParams() (string, string, int, float64, float64, float64, []string) {
	month := ""
	year := ""
	units := 0
	money := 0.0
	monthlyIncome := 2000.0 // modify if needed
	totalInvested := 0.0
	absPath, _ := filepath.Abs("nikkei-225-index-historical-chart-data.csv") // update if desirable
	dat, _ := os.ReadFile(absPath)
	lines := strings.Split(string(dat), "\r\n")
	return month, year, units, money, monthlyIncome, totalInvested, lines
}

func ParseLine(line string) (string, string, float64, error) {
	lineParts := strings.Split(line, ",")
	dateParts := strings.Split(lineParts[0], "/")
	currentMonth := dateParts[0]
	currentYear := dateParts[2]
	currentPrice, err := strconv.ParseFloat(lineParts[1], 64)
	currentPrice /= 100
	if err != nil {
		fmt.Println(err)
		return "", "", -1, err
	}
	return currentMonth, currentYear, currentPrice, nil
}
