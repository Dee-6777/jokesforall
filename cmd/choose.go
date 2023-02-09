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
	Short: "Choose the number of jokes you want to have",
	Long:  `Here you can choose any no of jokes you want to have!!!`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose called")
		ChooseJokes(args)
	},
}

func ChooseJokes(args []string) bool {
	num, err := strconv.Atoi(args[0])
	if err != nil {
		log.Printf("Not a valid argument. Error %v", err)
	}
	for i := 0; i < num; i++ {
		fmt.Println(GetRandomJoke())
	}
	if len(args) == 0 && err == nil {
		return true
	}
	return false
}

func init() {
	rootCmd.AddCommand(chooseCmd)
}
