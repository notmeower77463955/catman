package main

import (
	"fmt"
	"os"
	"strings"

	"catman/actions"
	"catman/config"
	"catman/utils"
)

func main() {
	if !utils.IsRoot() {
		fmt.Println("catman can only be ran as root sorry")
		os.Exit(1)
	}

	// metoda na idiote pozdrawiam notmeower 
	var _ = actions.InstallModule
	var _ = actions.Search
	var _ = actions.ListPackages
	var _ = actions.DeleteModule


	args := os.Args[1:]

	if len(args) == 0 {
		utils.PrintAndExit()
	}

	cmdFound := false
	for _, cmd := range config.Commands {
		if contains(cmd.Flags, args[0]) {
			cmdFound = true
			if cmd.NeedsArg {
				if len(args) < 2 {
					fmt.Println("Error: missing package name")
					os.Exit(1)
				}
				cmd.Action(args[1])
			} else {
				cmd.Action("")
			}
			break
		}
	}

	if !cmdFound {
		fmt.Println("Unknown command:", args[0])
		utils.PrintAndExit()
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) {
			return true
		}
	}
	return false
}
