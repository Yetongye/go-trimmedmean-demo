package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Yetongye/go-trimmedmean"
)

func generateRandomFloatSlice(n int) []float64 {
	rand.Seed(time.Now().UnixNano())
	s := make([]float64, n)
	for i := range s {
		s[i] = rand.NormFloat64()*10 + 50 // mean around 50, stddev around 10
	}
	return s
}

func generateRandomIntSlice(n int) []float64 {
	rand.Seed(time.Now().UnixNano())
	s := make([]float64, n)
	for i := range s {
		s[i] = float64(rand.Intn(100))
	}
	return s
}

func generateNormalData(n int) []float64 {
	data := make([]float64, n)
	for i := range data {
		data[i] = rand.NormFloat64()*10 + 50
	}
	return data
}

func bootstrapStats(data []float64, B int, trim float64) (meanMean, meanTrim float64, seMean, seTrim float64) {
	n := len(data)
	means := make([]float64, B)
	tmeans := make([]float64, B)

	for b := 0; b < B; b++ {
		sample := make([]float64, n)
		for i := 0; i < n; i++ {
			sample[i] = data[rand.Intn(n)]
		}
		// mean
		sum := 0.0
		for _, x := range sample {
			sum += x
		}
		means[b] = sum / float64(n)

		// trimmed mean
		tm, _ := trimmedmean.TrimmedMean(sample, trim)
		tmeans[b] = tm
	}

	// Compute means
	sumMean, sumTrim := 0.0, 0.0
	for i := 0; i < B; i++ {
		sumMean += means[i]
		sumTrim += tmeans[i]
	}
	meanMean = sumMean / float64(B)
	meanTrim = sumTrim / float64(B)

	// Compute standard errors
	for i := 0; i < B; i++ {
		seMean += math.Pow(means[i]-meanMean, 2)
		seTrim += math.Pow(tmeans[i]-meanTrim, 2)
	}
	seMean = math.Sqrt(seMean / float64(B-1))
	seTrim = math.Sqrt(seTrim / float64(B-1))

	return
}

func main() {
	floatData := generateRandomFloatSlice(100)
	intData := generateRandomIntSlice(100)

	// Calculate trimmed means with a 5% trim
	floatMean, _ := trimmedmean.TrimmedMean(floatData, 0.05)
	intMean, _ := trimmedmean.TrimmedMean(intData, 0.05)

	fmt.Printf("Trimmed Mean of 100 floats (0.05): %.2f\n", floatMean)
	fmt.Printf("Trimmed Mean of 100 ints   (0.05): %.2f\n", intMean)

	rand.Seed(time.Now().UnixNano())
	data := generateNormalData(100)
	meanMean, meanTrim, seMean, seTrim := bootstrapStats(data, 1000, 0.1)

	fmt.Printf("Bootstrap Mean:         %.2f (SE = %.2f)\n", meanMean, seMean)
	fmt.Printf("Bootstrap Trimmed Mean: %.2f (SE = %.2f)\n", meanTrim, seTrim)
}
