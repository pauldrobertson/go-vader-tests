package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    // "github.com/jonreiter/govader"
    "github.com/knuppe/vader"
)


func main(){

v, err := vader.NewVader("lexicons/en/en.zip")
        if err != nil {
                panic(err)
        }

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

sentiment := v.PolarityScores(input)

fmt.Println("Compound score:", sentiment.Compound)
fmt.Println("Positive score:", sentiment.Positive)
fmt.Println("Neutral score:", sentiment.Neutral)
fmt.Println("Negative score:", sentiment.Negative)
fmt.Println("Sentiment:",sentiment.Sentiment())
}