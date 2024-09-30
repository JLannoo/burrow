/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jlannoo/burrow/pkg/auth"
	"github.com/jlannoo/burrow/pkg/files"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a password",
	Long: `Remove a password from the password store.

The password will be permanently deleted from your computer.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		auth.Manager.Authenticate()
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		err := files.Manager.RemovePassword(name)
		if err != nil {
			fmt.Println("Error removing password:", err)
			return
		}

		fmt.Println("Password removed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
