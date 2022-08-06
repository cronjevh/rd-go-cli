/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package rd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
)

// configCmd represents the config command
// var configCmd = &cobra.Command{
// 	Use:   "config",
// 	Short: "A brief description of your command",
// 	Long: `A longer description`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// 4. Get client config from the flags.

// 		config, _ := configFlags.ToRESTConfig()

// 		// 5. Create a client-go instance for config.
// 		client, _ := kubernetes.NewForConfig(config)

// 		vinfo, _ := client.Discovery().ServerVersion()
// 		fmt.Println(vinfo)
// 	},
// }

func init() {
	// 1. Create a flags instance.
	configFlags := genericclioptions.NewConfigFlags(true)

	// 2. Create a cobra command.
	cmd := &cobra.Command{
		Use:  "config",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			builder := resource.NewBuilder(configFlags)

			namespace := ""
			if configFlags.Namespace != nil {
				namespace = *configFlags.Namespace
			}
			// enforceNamespace := namespace != ""

			obj, _ := builder.
				// Scheme teaches the Builder how to instantiate resources by names.
				WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
				// Where to look up.
				NamespaceParam(namespace).
				// What to look for.
				ResourceTypeOrNameArgs(true, args...).
				// Do look up, please.
				Do().
				// Convert the result to a runtime.Object
				Object()

			fmt.Println("# YAML ConfigMap representation")
			printr := printers.NewTypeSetter(scheme.Scheme).ToPrinter(&printers.YAMLPrinter{})
			if err := printr.PrintObj(obj, os.Stdout); err != nil {
				panic(err.Error())
			}
		},
	}

	// 3. Register flags with cobra.
	configFlags.AddFlags(cmd.PersistentFlags())

	rootCmd.AddCommand(cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
