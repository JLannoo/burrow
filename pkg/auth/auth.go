package auth

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/jlannoo/burrow/pkg/crypto"
	"github.com/jlannoo/burrow/pkg/files"
	"golang.org/x/term"
)

type Auth struct {
	HashedMasterPassword []byte
	ExpireTime           int64
}

func (a *Auth) AuthFileExists() bool {
	_, err := files.Manager.ReadFromMasterPasswordFile()
	return err == nil
}

func (a *Auth) AuthIsExpired() bool {
	currentTime := time.Now().UnixMilli()

	updateTime, err := files.Manager.GetFileUpdateTime(files.Manager.MasterPasswordFileName)
	if err != nil {
		return true
	}

	return currentTime-updateTime > a.ExpireTime
}

func (a *Auth) IsAuthed() bool {
	exists := a.AuthFileExists()
	expired := a.AuthIsExpired()

	return exists && !expired
}

func (a *Auth) GetAuth() ([]byte, error) {
	fileExists := a.AuthFileExists()

	if !fileExists {
		fmt.Println("No master password set. Please set one now.")
		fmt.Println("This password will be used to encrypt and decrypt your passwords from now on.\nIf you lose it, you will lose access to your passwords.")
	} else {
		fmt.Println("You are currently not authenticated. Please enter your master password to authenticate.")
	}

	fmt.Println("Enter master password: ")
	masterPassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}

	hashed := crypto.HashSHA256(string(masterPassword))

	if fileExists {
		passwordInFile, err := files.Manager.ReadFromMasterPasswordFile()
		if err != nil {
			return nil, err
		}

		if !crypto.CompareSHA256(string(masterPassword), passwordInFile) {
			return nil, errors.New("password entered is different from previously set one. Please retry")
		}
	}

	err = files.Manager.WriteToMasterPasswordFile(hashed)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func (a *Auth) Authenticate() ([]byte, error) {
	if !a.IsAuthed() {
		_, err := a.GetAuth()
		if err != nil {
			fmt.Println("Error authenticating:", err)
			os.Exit(1)
		}
	}

	masterPassword, err := files.Manager.ReadFromMasterPasswordFile()
	if err != nil {
		fmt.Println("Error reading master password:", err)
		os.Exit(1)
	}

	a.HashedMasterPassword = masterPassword

	return masterPassword, nil
}

var Manager = &Auth{
	ExpireTime: 10 * 60 * 1000, // 10 minutes
}
