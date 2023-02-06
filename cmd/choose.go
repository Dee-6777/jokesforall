/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chooseCmd represents the choose command
var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "A brief description of your command",
	Long:  `Here you can choose any no of jokes you want to have!!!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose called")
	},
}

func init() {
	rootCmd.AddCommand(chooseCmd)
}
