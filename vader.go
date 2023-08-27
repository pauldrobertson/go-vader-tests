package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "github.com/jonreiter/govader"
)


func main(){

// Get input string
fmt.Print("Input: ")
reader := bufio.NewReader(os.Stdin)
        // ReadString will block until the delimiter is entered
        input, err := reader.ReadString('\n')
        if err != nil {
                fmt.Println("An error occured while reading input. Please try again", err)
                return
        }

        // remove the delimeter from the string
        input = strings.TrimSuffix(input, "\n")

// Get sentiment scores
analyzer := govader.NewSentimentIntensityAnalyzer()
sentiment := analyzer.PolarityScores(input)

// Print scores for sentence
fmt.Println("Compound score:", sentiment.Compound)
fmt.Println("Positive score:", sentiment.Positive)
fmt.Println("Neutral score:", sentiment.Neutral)
fmt.Println("Negative score:", sentiment.Negative)
}
