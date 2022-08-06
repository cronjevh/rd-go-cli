package rd

import (
 "fmt"
 "os"

 "github.com/spf13/cobra"
 "github.com/spf13/viper"
)

var version = "0.0.1"
var cfgFile string

var rootCmd = &cobra.Command{
    Use:  "rd",
    Version: version,
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
