/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [name] [length]",
	Short: "Generate a new password",
	Long: `Generate a new password with the specified length.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
