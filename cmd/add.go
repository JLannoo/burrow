/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name] [password]",
	Short: "Add a new password",
	Long: `Add a new password to the password store.

The password will be encrypted using AES encryption and stored in a file on your computer.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
