package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//Calculate thesum of the deviations of the valuesin the  x-axis...
func CalculateSumDeviationX(data []int64) (sumDeviationX float64) {
	var (
		mean float64
	)
	for i := 1; i <= len(data); i++ {
		mean += float64(i)
	}
	mean /= float64(len(data))
	sumDeviationX = 0.0
	for i := 1; i <= len(data); i++ {
		sumDeviationX += math.Pow(float64(i)-mean, 2)
	}
	return
}

//Calculate the sum of the squared dviations of the y-axis...
func CalculateSumDeviationY(data []int64) (sumDeviationY float64) {
	var (
		mean float64
	)
	for _, v := range data {
		mean += float64(v)
	}
	mean /= float64(len(data))
	sumDeviationY = 0.0
	for _, v := range data {
		sumDeviationY += math.Pow(float64(v)-mean, 2)
	}
	return
}

//Calculate the  value of `r` the Pearsons correlation coefficient
func PearsonsCorrelation(sumDeviationX, sumDeviationY float64, data []int64) (r float64) {
	var (
		n    = float64(len(data))
		sumX float64
		sumY float64
	)
	for i := 1; i <= len(data); i++ {
		sumX += float64(i)
		sumY += float64(data[i-1])
	}
	meanX := sumX / n
	meanY := sumY / n
	numerator := 0.0
	denominator := 0.0
	for i := 1; i <= len(data); i++ {
		numerator += (float64(i) - meanX) * (float64(data[i-1]) - meanY)
	}
	denominator += sumDeviationX + sumDeviationY
	r = (numerator / denominator)
	return
}

//Calculate the standard deviation of both x and y values...
func StadardDeviation(data []int64) (stdDevX, stdDevY float64) {
	var (
		n    = float64(len(data))
		sumX float64
		sumY float64
	)
	for i := 1; i <= len(data); i++ {
		sumX += float64(i)
		sumY += float64(data[i-1])
	}
	meanX := sumX / n
	meanY := sumY / n
	for i := 0; i < len(data); i++ {
		stdDevX += math.Pow((float64(i) - meanX), 2)
		stdDevY += math.Pow((float64(data[(i)]) - meanY), 2)
	}
	stdDevX = math.Sqrt(stdDevX / float64(len(data)))
	stdDevY = math.Sqrt(stdDevY / float64(len(data)-1))
	return
}

//Calculate the slope of the Linear Regression Line function...
func CalculateSlope(r, stdDevX, stdDevY float64) (b float64) {
	b = (r * stdDevY) / stdDevX
	return
}

//Calculate the y-intercept of the Linear Regression Line function...
func CalculateIntercept(data []int64, b float64) (a float64) {
	var (
		sumX float64
		sumY float64
		n    = float64(len(data))
	)
	for i := 0; i < len(data); i++ {
		sumX += float64(i)
		sumY += float64(data[i])
	}
	meanX := sumX / n
	meanY := sumY / n
	a = meanY - (b * meanX)
	return
}

//Calculate the y-value for a given x-value on the Linear Regression Line function...
func LinearRegression(data []int64) (a, b, r float64) {
	var (
		sumDeviationX float64
		sumDeviationY float64
		stdDevX       float64
		stdDevY       float64
	)
	stdDevX, stdDevY = StadardDeviation(data)
	sumDeviationX = CalculateSumDeviationX(data)
	sumDeviationY = CalculateSumDeviationY(data)
	r = PearsonsCorrelation(sumDeviationX, sumDeviationY, data)
	b = CalculateSlope(r, stdDevX, stdDevY)
	a = CalculateIntercept(data, b)

	return
}

func main() {
	var data []int64
	//read from a file
	file, err := os.Open(os.Args[1])
	if err != nil {
		println("Error opening file: ", err)
		os.Exit(0)
	}
	defer file.Close()

	//read data line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			println("Error parsing number: ", err)
			continue
		}
		data = append(data, num)
	}
	if err := scanner.Err(); err != nil {
		println("Error reading file: ", err)
	}
	a, b, r := LinearRegression(data)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\nPearson Correlation Coefficient: %.10f\n", b, a, r)
}
