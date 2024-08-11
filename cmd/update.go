/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [name] [password]",
	Short: "Update a password",
	Long: `Update a password in the password store.

The password will be re-encrypted and stored in a file on your computer.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
