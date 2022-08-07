/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package flip

import (
	"fmt"
    "os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/manifoldco/promptui"
)

// flipCmd represents the flip command
var FlipCmd = &cobra.Command{
	Use:   "flip",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		prompt := promptui.Select{
			Label: "Select Day",
			Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
				"Saturday", "Sunday"},
		}
	
		_, result, err := prompt.Run()
	
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	
		fmt.Printf("You choose %q\n", result)

		app := "echo"

		arg0 := "-e"
		arg1 := "Hello world"
		arg2 := "\n\tfrom"
		arg3 := "golang"
	
		cli := exec.Command(app, arg0, arg1, arg2, arg3)
		stdout, err := cli.Output()
	
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	
		// Print the output
		fmt.Println(string(stdout))

		viper.Set("choice.day", result)
		viper.WriteConfig() 
	},
}

func init() {
	// rootCmd.AddCommand(flipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
