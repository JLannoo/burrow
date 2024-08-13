/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jlannoo/burrow/pkg/auth"
	"github.com/jlannoo/burrow/pkg/crypto"
	"github.com/jlannoo/burrow/pkg/files"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get a password",
	Long: `Get a password from the password store.

The password will be decrypted using AES encryption and displayed in the terminal.`,
	Args: cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		auth.Manager.Authenticate()
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		var password []byte
		var user []byte

		for len(password) == 0 {
			unlockKey, err := crypto.GenerateUnlockKey(string(auth.Manager.HashedMasterPassword))
			if err != nil {
				fmt.Println("Error generating encryption key", err)
				return
			}

			fileBytes, err := files.Manager.ReadFromFile(name)
			if err != nil {
				fmt.Println("Could not find password file for", name)
				return
			}

			split := files.Manager.SplitBytes(fileBytes)
			encryptedPassword := split[0]

			if len(split) > 1 {
				encryptedUser := split[1]
				user, err = crypto.Decrypt(encryptedUser, unlockKey)
				if err != nil {
					fmt.Println("Error decrypting username:", err)
					return
				}
			}

			password, err = crypto.Decrypt(encryptedPassword, unlockKey)
			if len(password) == 0 || err != nil {
				fmt.Println("Incorrect master password, please try again")
				auth.Manager.GetAuth()
			}
		}

		if display, _ := cmd.Flags().GetBool("display"); !display {
			err := clipboard.Init()
			clipboard.Write(clipboard.FmtText, password)

			if err != nil {
				fmt.Println("Error copying password to clipboard")
				return
			}

			if user != nil {
				fmt.Printf("Password with username '%s' for %s copied to clipboard\n", user, name)
			} else {
				fmt.Printf("Password for %s copied to clipboard\n", name)
			}
		} else {
			if user != nil {
				fmt.Printf("Username for %s: %s\n", name, user)
			}
			fmt.Printf("Password for %s: %s\n", name, password)
		}
	},
}

func init() {
	getCmd.Flags().BoolP("display", "d", false, "Display password in terminal instead of copying to clipboard")
	rootCmd.AddCommand(getCmd)
}
