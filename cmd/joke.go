/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Generate a random joke",
	Long:  `A joke generator CLI for some giggle and laughs...Have fun!!!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("joke called")
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}

// joke structure
type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// function to return random joke
func getRandomJoke() bool {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Could not unmarshall response %v", err)
	}
	fmt.Println(string(joke.Joke))
	return len(joke.Joke) > 0
}

// gets the joke from the API
func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not get a joke. Error %v", err)
		os.Exit(1)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Joke generator CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. Error %v", err)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read the response. Error%v", err)
	}
	return responseBytes
}
