package rd

import (
    "fmt"

    "github.com/cronjevh/rd-go-cli/pkg/rd"
    "github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
    Use:   "reverse",
    Aliases: []string{"rev"},
    Short:  "Reverses a string",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        res := rd.Reverse(args[0])
        fmt.Println(res)
    },
}

func init() {
    rootCmd.AddCommand(reverseCmd)
}