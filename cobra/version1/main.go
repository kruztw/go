package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func initConfig() {
}

func init() {
	rootCmd = &cobra.Command{
		Short: "This is Short",
		Long: `
			This is supperrrrrrrrrrrrrrrrrrrrrrrrr 
			Long commentttttttttttt
			`,

		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			fmt.Printf("hello %v\n", name)
		},
	}

	cobra.OnInitialize(initConfig)

	var v string
	rootCmd.Flags().StringVarP(&v, "name", "n", "default", "your name")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("run rootCmd failed: %v\n", err)
	}
}
