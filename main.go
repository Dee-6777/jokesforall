/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"time"

	"github.com/Dee-6777/jokesforall/cmd"
	"github.com/Dee-6777/jokesforall/ui"
)

func main() {
	cmd.Execute()
	var choice string
	DurationOfTime := time.Duration(6) * time.Second
	f := func() {
		fmt.Println("Want to work with UI (y/n)?")
	}
	Timer1 := time.AfterFunc(DurationOfTime, f)
	defer Timer1.Stop()
	fmt.Scanln(&choice)
	if choice == "y" {
		ui.Greet()
	}
}
