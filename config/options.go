package config

import (
	"catman/actions"
	"catman/utils"
)

type Command struct {
	Flags    []string
	NeedsArg bool
	Action   func(string)
}

var Commands = []Command{
	{
		Flags:    []string{"-h", "--help", "--herp"},
		NeedsArg: false,
		Action:   func(_ string) { utils.PrintAndExit() },
	},
	{
		Flags:    []string{"-i", "--install"},
		NeedsArg: true,
		Action:   actions.InstallModule,
	},
	{
		Flags:    []string{"-s", "--search"},
		NeedsArg: true,
		Action:   actions.Search,
	},
	{
		Flags:    []string{"-l", "--list"},
		NeedsArg: false,
		Action:   func(_ string) { actions.ListPackages() },
	},
	{
		Flags:    []string{"-d", "--delete"},
		NeedsArg: true,
		Action:   actions.DeleteModule,
	},
}
