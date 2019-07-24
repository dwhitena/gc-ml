package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// Open the Advertising data set.
	f, err := os.Open("../data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// We should have 4 fields per line. By setting
	// FieldsPerRecord to 4, we can validate that each of
	// the rows in our CSV has the correct number of fields.
	reader.FieldsPerRecord = 4

	// features and response will hold our successfully parsed
	// TV and Sales values respectively.
	line := 1
	var features []float64
	var response []float64
	for {

		// Read in a row. Check if we are at the end of
		// the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip the header.
		if line == 1 {
			line++
			continue
		}

		// If we had a parsing error, log the error
		// and move on.
		if err != nil {
			log.Println(err)
			line++
			continue
		}

		// Try to parse the values we want (the TV and Sales values)
		// as floats.
		var tv float64
		if tv, err = strconv.ParseFloat(record[0], 64); err != nil {
			log.Printf("Unexpected type in TV column at line %d\n", line)
			fmt.Println(record)
			line++
			continue
		}

		var sales float64
		if sales, err = strconv.ParseFloat(record[3], 64); err != nil {
			log.Printf("Unexpected type in Sales column at line %d\n", line)
			line++
			continue
		}

		// Append the records to our feature and response slices.
		features = append(features, tv)
		response = append(response, sales)
		line++
	}

	// Train our linear regression model.
	w, b, _ := sgdTrain(features, response, 0.1, 200)

	// Print our results.
	fmt.Printf("\nRegression Formula:\ny = %0.2f * x + %0.2f\n\n", w, b)
}

// sgdTrain calculates the ideal parameters using SGD for
// a linear regression model.
func sgdTrain(features, response []float64, lr float64, epochs int) (float64, float64, float64) {

	// Initialize the weight and bias.
	w := 0.0
	b := 0.0

	// Set the number of observations in the data.
	n := float64(len(response))

	// Loop over the number of epochs.
	loss := 0.0
	for i := 0; i < epochs; i++ {

		// Calculate current predictions.
		var predictions []float64
		for _, x := range features {
			predictions = append(predictions, w*x+b)
		}

		// Calculate the loss for this epoch.
		loss = 0.0
		for idx, p := range predictions {
			loss += squaredError(p, response[idx]) / n
		}

		// Output some info to standard out so we know
		// how training is progressing.
		if i%10 == 0 {
			fmt.Printf("Epoch %d, Loss %0.4f\n", i, loss)
		}

		// Calculate the gradients for w and b.
		wGradient := 0.0
		bGradient := 0.0
		for idx, p := range predictions {
			wGradient += -(2 / n) * (features[idx] * (response[idx] - p))
			bGradient += -(2 / n) * (response[idx] - p)
		}

		// Update the weight and bias.
		w = w - (lr * wGradient)
		b = b - (lr * bGradient)
	}

	return w, b, loss
}

// squaredError returns the error associted with a particular
// pair of prediction and observation.
func squaredError(observation, prediction float64) float64 {
	return math.Pow(observation-prediction, 2)
}
