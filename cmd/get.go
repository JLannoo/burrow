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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get a password",
	Long: `Get a password from the password store.

The password will be decrypted using AES encryption and displayed in the terminal.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		fmt.Println("Enter master password:")
		masterPassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Error reading master password:", err)
			return
		}

		unlockKey, err := crypto.GenerateUnlockKey(string(masterPassword))
		if err != nil {
			fmt.Println("Error generating encryption key", err)
			return
		}

		encryptedPassword, err := files.Manager.ReadFromFile(name)
		if err != nil {
			fmt.Println("Could not find password file for", name)
			return
		}

		password, err := crypto.Decrypt(encryptedPassword, unlockKey)
		if err != nil {
			fmt.Println("Error decrypting password, check your master password")
			return
		}

		fmt.Printf("Password for %s: %s\n", name, password)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
