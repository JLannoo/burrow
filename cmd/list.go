/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jlannoo/burrow/pkg/files"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all passwords",
	Long: `List all passwords stored in the password store.

This command will display all of the passwords that you have stored on your computer.`,
	Run: func(cmd *cobra.Command, args []string) {
		list, err := files.Manager.GetAllPasswords()
		if err != nil {
			fmt.Println("Error listing passwords:", err)
			return
		}

		for _, name := range list {
			fmt.Println(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
