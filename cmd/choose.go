/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// chooseCmd represents the choose command
var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "A brief description of your command",
	Long:  `Here you can choose any no of jokes you want to have!!!`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose called")
		chooseJokes(args)
	},
}

func chooseJokes(args []string) {
	var num int
	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Printf("Couldnot read the response. Error occured %v", err)
	}
	for i := 0; i < num; i++ {
		getRandomJoke()
	}
}
func init() {
	rootCmd.AddCommand(chooseCmd)
}
