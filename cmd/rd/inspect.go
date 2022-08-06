package rd

import (
    "fmt"

    "github.com/cronjevh/rd-go-cli/pkg/rd"
    "github.com/spf13/cobra"
)

var onlyDigits bool
var inspectCmd = &cobra.Command{
    Use:   "inspect",
    Aliases: []string{"insp"},
    Short:  "Inspects a string",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {

        i := args[0]
        res, kind := rd.Inspect(i, onlyDigits)

        pluralS := "s"
        if res == 1 {
            pluralS = ""
        }
        fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
    },
}

func init() {
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
    rootCmd.AddCommand(inspectCmd)
}