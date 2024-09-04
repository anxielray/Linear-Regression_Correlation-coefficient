package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the file path as a command-line argument")
	}

	filePath := os.Args[1]
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data := string(file)
	r := correlation(data)
	a, b := InterceptAndSlope(data)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\nPearson Correlation Coefficient: %.10f\n", a, b, r)
}

func InterceptAndSlope(data string) (slope, intercept float64) {
	var (
		Xs []float64
		Ys []float64
	)
	for x, value := range strings.Fields(data) {
		y, _ := strconv.ParseFloat(value, 64)
		Ys = append(Ys, y)
		Xs = append(Xs, float64(x))
	}

	n := float64(len(Ys))
	sumX := sum(Xs)
	sumY := sum(Ys)
	sumX2 := sumSquares(Xs)
	sumXY := sumProduct(Xs, Ys)

	slope = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept = (sumY - slope*sumX) / n
	return
}

func sum(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum
}

func sumSquares(Xs []float64) float64 {
	sum := 0.0
	for _, value := range Xs {
		sum += value * value
	}
	return sum
}

func sumProduct(Xs, Ys []float64) float64 {
	sum := 0.0
	for i := range Xs {
		sum += Xs[i] * Ys[i]
	}
	return sum
}

func correlation(data string) float64 {
	var (
		Xs []float64
		Ys []float64
	)

	for x, value := range strings.Fields(data) {
		y, _ := strconv.ParseFloat(value, 64)
		Ys = append(Ys, y)
		Xs = append(Xs, float64(x))
	}

	// n := float64(len(Ys))
	meanX := mean(Xs)
	meanY := mean(Ys)

	numerator := 0.0
	sumX2 := 0.0
	sumY2 := 0.0
	for i := range Xs {
		xDiff := Xs[i] - meanX
		yDiff := Ys[i] - meanY
		numerator += xDiff * yDiff
		sumX2 += xDiff * xDiff
		sumY2 += yDiff * yDiff
	}

	denominator := math.Sqrt(sumX2 * sumY2)
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func mean(data []float64) float64 {
	return sum(data) / float64(len(data))
}
