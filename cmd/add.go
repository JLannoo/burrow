/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"syscall"

	"github.com/jlannoo/burrow/pkg/crypto"
	"github.com/jlannoo/burrow/pkg/files"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add a new password",
	Long: `Add a new password to the password store.

The password will be encrypted using AES encryption and stored in a file on your computer.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fmt.Println("Enter master password: ")
		masterPassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Error reading master password:", err)
			return
		}

		fmt.Println("Enter password to store: ")
		password, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Error reading password:", err)
			return
		}

		unlockKey, err := crypto.GenerateUnlockKey(string(masterPassword))
		if err != nil {
			fmt.Println("Error generating unlock key:", err)
			return
		}

		encryptedPassword, err := crypto.Encrypt([]byte(password), unlockKey)
		if err != nil {
			fmt.Println("Error encrypting password:", err)
			return
		}

		err = files.Manager.WriteToFile(encryptedPassword, name)
		if err != nil {
			fmt.Println("Could not create password file:", err)
			return
		}

		fmt.Println("Password added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
