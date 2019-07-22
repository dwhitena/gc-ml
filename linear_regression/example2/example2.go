package main

import (
	"bufio"
	"log"
	"os"
	"path"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	// Open the advertising dataset file.
	f, err := os.Open("../data/advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	adDF := dataframe.ReadCSV(f)

	// Calculate the number of elements in each set.
	trainingNum := (4 * adDF.Nrow()) / 5
	testNum := adDF.Nrow() / 5
	if trainingNum+testNum < adDF.Nrow() {
		trainingNum++
	}

	// Create the subset indices.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// Enumerate the training indices.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate the test indices.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Create the subset dataframes.
	trainingDF := adDF.Subset(trainingIdx)
	testDF := adDF.Subset(testIdx)

	// Create a map that will be used in writing the data
	// to files.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// Create the respective files.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// Save the filtered dataset file.
		f, err := os.Create(path.Join("../data/", setName))
		if err != nil {
			log.Fatal(err)
		}

		// Create a buffered writer.
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
