package main

import "fmt"

func callbackHelp() {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Here are your available commands:")

	availableCommand := getCommands()
	for _, command := range availableCommand {
		fmt.Printf(" - %s: %s\n", command.name, command.description)
	}
	fmt.Println("")
}