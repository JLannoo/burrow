/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/jlannoo/burrow/pkg/auth"
	"github.com/jlannoo/burrow/pkg/crypto"
	"github.com/jlannoo/burrow/pkg/files"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [name] [length]",
	Short: "Generate a new password",
	Long:  `Generate a new password with the specified length.`,
	Args:  cobra.RangeArgs(1, 2),
	PreRun: func(cmd *cobra.Command, args []string) {
		auth.Manager.Authenticate()
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		user := cmd.Flag("username").Value.String()

		length := 16
		var err error
		if len(args) == 2 {
			length, err = strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Error parsing length:", err)
				return
			}
		}

		pg := crypto.PasswordGenerator{
			Length:            length,
			SpecialCharacters: true,
			Numbers:           true,
			Lowercase:         true,
			Uppercase:         true,
		}

		password := pg.Generate()
		if err != nil {
			fmt.Println("Error generating password:", err)
			return
		}

		unlockKey, err := crypto.GenerateUnlockKey(string(auth.Manager.HashedMasterPassword))
		if err != nil {
			fmt.Println("Error generating unlock key:", err)
			return
		}

		encryptedUser := []byte{}
		if user != "" {
			encryptedUser, err = crypto.Encrypt([]byte(user), unlockKey)
			if err != nil {
				fmt.Println("Error encrypting username:", err)
				return
			}
		}

		encryptedPassword, err := crypto.Encrypt([]byte(password), unlockKey)
		if err != nil {
			fmt.Println("Error encrypting password:", err)
			return
		}

		fileBytes := files.Manager.JoinBytes(encryptedPassword, encryptedUser)

		err = files.Manager.WriteToFile(fileBytes, name)
		if err != nil {
			fmt.Println("Could not create password file:", err)
			return
		}

		fmt.Println("Password added successfully!")
	},
}

func init() {
	generateCmd.Flags().StringP("username", "u", "", "Username associated to the password")
	rootCmd.AddCommand(generateCmd)
}
