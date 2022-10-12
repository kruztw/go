package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "Short",
	Long:  `Long`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("call Run")
		argT, _ := cmd.Flags().GetString("name")
		fmt.Printf("argT = %v\n", argT)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	fmt.Println("call initConfig")
}

func init() {
	fmt.Println("call init")
	cobra.OnInitialize(initConfig)

	var v string
	rootCmd.Flags().StringVarP(&v, "name", "n", "default", "your name")
}
