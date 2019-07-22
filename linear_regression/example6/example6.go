package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// scratchPredict makes a prediction for sales
// based on the from-scratch SGD-trained model.
func scratchPredict(tv float64) float64 {
	return 0.55*tv + 0.23
}

// sajariSinglePredict makes a prediction for sales
// based on the sajari single regression model.
func sajariSinglePredict(tv float64) float64 {
	return 0.57*tv + 0.22
}

// sajariMultiPredict makes a prediction for sales
// based on the sajari multiple regression model.
func sajariMultiPredict(tv, radio float64) float64 {
	return 0.55*tv + 0.35*radio + 0.05
}

func main() {

	// Open the test dataset file.
	f, err := os.Open("../data/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Loop over the test data predicting y and evaluating the prediction
	// with the mean squared error.
	var mSEScratch float64
	var mSESajariSingle float64
	var mSESajariMulti float64
	for i, record := range testData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the observed sales, or "y".
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the tv value.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the radio value.
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained models.
		yScratch := scratchPredict(tvVal)
		ySajariSingle := sajariSinglePredict(tvVal)
		ySajariMulti := sajariMultiPredict(tvVal, radioVal)

		// Add the to the mean squared error.
		mSEScratch += math.Pow(yObserved-yScratch, 2) / float64(len(testData))
		mSESajariSingle += math.Pow(yObserved-ySajariSingle, 2) / float64(len(testData))
		mSESajariMulti += math.Pow(yObserved-ySajariMulti, 2) / float64(len(testData))
	}

	// Output the RMSE to standard out.
	fmt.Printf("\nRMSE (Scratch) = %0.2f\n", math.Sqrt(mSEScratch))
	fmt.Printf("RMSE (Sajari Single) = %0.2f\n", math.Sqrt(mSESajariSingle))
	fmt.Printf("RMSE (Sajari Multi) = %0.2f\n\n", math.Sqrt(mSESajariMulti))
}
