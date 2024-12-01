package ornament

import "github.com/spf13/cobra"

type Ornament interface {
	Run(cmd *cobra.Command, args []string) error
	Part1(file string) (string, error)
	Part2(file string) (string, error)
}
