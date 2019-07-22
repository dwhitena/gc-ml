package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
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
	reader.FieldsPerRecord = 2

	// Read in all of the CSV records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// features and labels will hold all the float values that
	// will eventually be used in our training.
	var features [][]float64
	var labels []float64

	// Sequentially move the rows into the slices of floats.
	for idx, record := range rawCSVData {

		// Skip the header row.
		if idx == 0 {
			continue
		}

		// Add the FICO score feature.
		featureVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		features = append(features, []float64{featureVal})

		// Add the class label.
		labelVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		labels = append(labels, labelVal)
	}

	// New Logistic regression model from goml (optimization method:
	// Batch Gradient Ascent, Learning rate: 1e-4, Regulatization term: 6,
	// Max Iterations: 800, Dataset to learn fron: features,
	// Expected results dataset: labels)
	model := linear.NewLogistic(base.BatchGA, 1e-4, 1, 800, features, labels)

	// Train the logistic regression model.
	if err = model.Learn(); err != nil {
		log.Fatal(err)
	}
}
