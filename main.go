package main

import ("fmt"
		"bufio"
		"os"
		"github.com/Adam-Tay/pokecli/pkg/pokeapi"
	)

var locationCount int = 1

type cliCommand struct {
	name string
	description string
	callback func() error
}


func commandHelp() error {
	commands := getCommands()
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Exiting pokedex...")
	os.Exit(0)
	return nil
}

func commandMap() error {
	next := locationCount + 19
	locations, err := pokeapi.GetLocationsInPokemon(locationCount, next)
	locationCount = next
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("The next location(s)")
	for i := 0; i < len(locations); i++ {
		fmt.Println(locations[i].Name)
	}
	return nil
}

func commandMapB() error {
	fmt.Println("The previous location(s)")
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Display the names of the  20 next locations in the pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Display the names of the 20 previous locations in the pokemon world",
			callback: commandMapB,
		}, 
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	fmt.Println("Welcome to your pokedex, type \"help\" for a list of commands")

	for {
		fmt.Print("Enter your command: ")
		for scanner.Scan() {
			input := scanner.Text()
		
			if cmd, ok := commands[input]; ok {
				err := cmd.callback()
				if err != nil {
					fmt.Println("Error executing command")
				}

				break
				
			} else {
				fmt.Println("Command not found:", input)
				break
			}
		}
		
	}
	
	
}