package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {

	// Open the training dataset file.
	f, err := os.Open("../data/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Read in all of the CSV records
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// In this case we are going to try and model our sales
	// y by the tv and radio features plus an intercept.  As such, let's create
	// the struct needed to train a model using github.com/sajari/regression.
	var r regression.Regression
	r.SetObserved("sales")
	r.SetVar(0, "tv")
	r.SetVar(1, "radio")

	// Loop over the CSV records adding the training data.
	for i, record := range trainingData {

		// Skip the header.
		if i == 0 {
			continue
		}

		// Parse the diabetes progression measure, or "y".
		yVal, err := strconv.ParseFloat(record[3], 64)
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

		// Add these points to the regression value.
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))
	}

	// Train/fit the regression model.
	r.Run()

	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}
