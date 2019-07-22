package main

import (
	"fmt"
	"strings"

	"github.com/machinebox/sdk-go/textbox"
)

func main() {

	// Connect to MachineBox.
	machBoxIP := "http://localhost:8080"
	client := textbox.New(machBoxIP)

	// Example text input.
	positiveStatement := "I am so excited to be teaching to super awesome, fun workshop!"
	//negativeStatement := "It is sad, depressing, and unfortunate that this workshop will terminate at the end of the day."

	// Analyze the text with TextBox.
	analysis, err := client.Check(strings.NewReader(positiveStatement))
	if err != nil {
		fmt.Println(err)
	}

	// Calculate the sentiment.
	sentimentTotal := 0.0
	for _, sentence := range analysis.Sentences {
		sentimentTotal += sentence.Sentiment
	}

	// Higher sentitment is more positive, and lower is more negative.
	fmt.Printf("\nSentiment: %0.2f\n\n", sentimentTotal/float64(len(analysis.Sentences)))
}
