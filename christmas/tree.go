package christmas

import (
	"os"
	"strconv"

	"github.com/EdgeLordKirito/ChristmasTree/ornament"
	"github.com/spf13/cobra"
)

type Tree struct {
	rootCmd *cobra.Command
}

func New(name string) *Tree {
	var rootCommand = &cobra.Command{
		Use: name,
	}
	result := Tree{rootCmd: rootCommand}
	return &result
}

func (self *Tree) AddOrnament(day int, deco ornament.Ornament) {
	self.rootCmd.AddCommand(createSubCommand(day, deco))
}

func createSubCommand(day int, deco ornament.Ornament) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "Day" + strconv.Itoa(day),
		Short: "Sub Command for Day " + strconv.Itoa(day),
		Args:  cobra.ExactArgs(1),
		RunE:  deco.Run,
	}
	return cmd
}

func (self *Tree) Run() {
	if err := self.rootCmd.Execute(); err != nil {
		os.Exit(1) // Let Cobra handle printing the error
	}
}
