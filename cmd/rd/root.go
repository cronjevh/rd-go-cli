package rd

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:  "rd",
    Short: "rd - a simple CLI to transform and inspect strings",
    Long: `rd is a super fancy CLI (kidding)
   
One can use rd to modify or inspect strings straight from the terminal`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}