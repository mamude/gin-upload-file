package helpers

import (
	"strconv"
	"strings"
	"time"
)

func ParseCPF(value string) string {
	return value
}

func ParseCNPJ(value string) string {
	if value == "NULL" {
		return ""
	}
	return value
}

func ParseDate(value string) time.Time {
	if value == "NULL" {
		return time.Now()
	}
	layout := "2006-01-02"
	date, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return date
}

func ParseDecimal(value string) float64 {
	if value == "NULL" {
		return 0
	}
	value = strings.Replace(value, ",", ".", 1)
	decimal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(err)
	}
	return decimal
}
