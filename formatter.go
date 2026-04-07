package main

import "github.com/fatih/color"

func printColored(label, message string) {
	color.Green("[%s] ", label)
	color.Yellow(message)
}
