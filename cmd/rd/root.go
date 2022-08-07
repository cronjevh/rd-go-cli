package rd

import (
	"fmt"
	"os"

	"github.com/cronjevh/rd-go-cli/cmd/rd/flip"
	"github.com/cronjevh/rd-go-cli/cmd/rd/incubator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "0.0.1"
var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "rd",
	Version: version,
	Short:   "rd - a simple CLI to transform and inspect strings",
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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".rd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	initConfig()
	cobra.OnInitialize()

	// setDefaults()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// Add my subcommand palette
	// rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(incubator.IncubatorCmd)
	rootCmd.AddCommand(flip.FlipCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toolbox.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
