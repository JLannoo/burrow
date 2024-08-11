/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "burrow",
	Short: "Burrow is a CLI tool for managing your passwords securely",
	Long: `Burrow is a CLI tool for managing your passwords securely.

It uses AES encryption to store your passwords in a file on your computer.
You can add, remove, and list passwords using the CLI.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
}
