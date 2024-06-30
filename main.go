package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/Adam-Tay/pokecli/pkg/pokeapi"
)

var locationCount int = 1
var userPokemon map[string]pokeapi.Pokemon = make(map[string]pokeapi.Pokemon)

type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
}

func commandHelp(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("this command does not take an argument")
	}
	commands := getCommands()
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("this command does not take an argument")
	}
	fmt.Println("Exiting pokedex...")
	os.Exit(0)
	return nil
}

func commandMap(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("this command does not take an argument")
	}
	next := locationCount + 19
	locations, err := pokeapi.GetLocationsInPokemon(locationCount, next)
	if err != nil {
		fmt.Println(err)
		return err
	}
	locationCount = next
	fmt.Println("The next location(s)")
	for i := 0; i < len(locations); i++ {
		fmt.Println(locations[i].Name)
	}
	return nil
}

func commandMapB(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("this command does not take an argument")
	}
	prev := locationCount - 19
	if prev < 1 {
		return errors.New("error, cannot go backwards without first going forward")
	}
	locations, err := pokeapi.GetLocationsInPokemon(locationCount, prev)
	if err != nil {
		fmt.Println(err)
		return err
	}
	locationCount = prev
	fmt.Println("The previous location(s)")
	for i := 0; i < len(locations); i++ {
		fmt.Println(locations[i].Name)
	}
	return nil
}

func commandExplore(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("enter an area name next time")
	}
	areaId := args[0]
	fmt.Println("Exploring", areaId, "...")
	area, err := pokeapi.GetLocationInPokemonByName(areaId)

	if err != nil {
		return err
	}

	for _, pokemon := range area.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("enter a pokemon name next time")
	}

	poke, err := pokeapi.GetPokemonByName(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Throwing a Pokeball at", args[0]+"...")
	chance := rand.Intn(100)

	if poke.BaseExperience%100 > chance {
		fmt.Println(args[0], "escaped!")
		return nil
	}
	userPokemon[args[0]] = poke
	fmt.Println(args[0], "was caught!")

	return nil
}

func commandInspect(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("enter a pokemon name next time")
	}

	p, ok := userPokemon[args[0]]
	if !ok {
		fmt.Println("You have not caught", args[0], "yet")
		return nil
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Height:", p.Height)
	fmt.Println("Weight:", p.Weight)
	fmt.Println("Stats:")
	//fmt.Println("\t-hp:", p.Stats.Stat.)
	for _, stat := range p.Stats {
		fmt.Printf("  -"+stat.Stat.Name+": %d\n", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Println("  -" + t.Type.Name)
	}
	return nil
}

func commandPokedex(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("this command does not take an argument")
	}
	if len(userPokemon) == 0 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, p := range userPokemon {
		fmt.Println("  -", p.Name)
	}

	return nil
}
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of the  20 next locations in the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the 20 previous locations in the pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Usage: explore <area_name>, find pokemon in <area_name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Usage: catch <pokemon_name>, catch a pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Usage: inspect <pokemon_name>, display info of the pokemon you've caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display the names of all the pokemon you've caught",
			callback:    commandPokedex,
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
			parts := strings.Fields(input)

			if len(parts) == 0 {
				fmt.Println("Enter \"help\" for list of commands")
			} else {
				if cmd, ok := commands[parts[0]]; ok {
					err := cmd.callback(parts[1:])
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

}
