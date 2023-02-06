/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("joke called")
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}
