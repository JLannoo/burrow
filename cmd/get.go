/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get a password",
	Long: `Get a password from the password store.

The password will be decrypted using AES encryption and displayed in the terminal.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
