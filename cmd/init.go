/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jlannoo/burrow/pkg/crypto"
	"github.com/jlannoo/burrow/pkg/files"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the password store",
	Long: `Initialize the password store.

This command will create a new directory on your computer to store your passwords, as well as setup your master password.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Make .key
		keyBytes, err := crypto.GenerateRandomKey()
		if err != nil {
			fmt.Println("Error generating random key:", err)
			return
		}

		err = files.Manager.WriteToSecretKeyFile(keyBytes)
		if err != nil {
			fmt.Println("Error writing key to file:", err)
			return
		}

		fmt.Printf("Wrote key to file %s\n", files.Manager.SecretKeyFileName)
		fmt.Printf("KEEP THIS KEY SAFE! If you lose it, you will lose access to your passwords.\n")

		fmt.Println("Password store initialized!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
