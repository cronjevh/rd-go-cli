package incubator

import (
	"errors"
	"fmt"
	"strconv"
	"os"

	"github.com/cronjevh/rd-go-cli/pkg/rd"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverses a string",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		validate := func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("Invalid number")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Number",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if len(os.Args) <= 3 {
			fmt.Println("if less than trigger Arg length",len(os.Args))
		}

		fmt.Printf("You choose %q\n", result)

		res := rd.Reverse(result)
		fmt.Println(res)
	},
}

func init() {
	IncubatorCmd.AddCommand(reverseCmd)
}
