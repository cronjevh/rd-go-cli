/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package incubator

import (
	// "fmt"

	"github.com/spf13/cobra"
)

// incubatorCmd represents the incubator command
var IncubatorCmd = &cobra.Command{
	Use:   "incubator",
	Short: "A brief description of your command",
	Long: `Incubator submenu`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// rootCmd.AddCommand(incubatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incubatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incubatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
