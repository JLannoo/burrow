package crypto

import (
	"crypto/rand"
	"math/big"
	"strings"
)

var LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
var UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var NUMBERS = "01234567890123456789" // Repeat to make odds more balanced
var SPECIAL_CHARACTERS = "!@#$%^&*()_+:!@#$%^&*()_+:" // Repeat to make odds more balanced

type PasswordGenerator struct {
	// Length of password (default 16)
	Length int

	// Allow special characters in password
	SpecialCharacters bool
	// Maximum number of special characters in password
	SpecialMax int

	// Allow Numbers in password
	Numbers bool
	// Maximum number of numbers in password
	NumbersMax int

	// Allow Uppercase letters in password
	Uppercase bool
	// Maximum number of uppercase letters in password
	UppercaseMax int

	// Allow Lowercase letters in password
	Lowercase bool
	// Maximum number of lowercase letters in password
	LowercaseMax int
}

func (pg *PasswordGenerator) Generate() string {
	// Default length to 16
	if pg.Length == 0 {
		pg.Length = 16
	}

	// If bool is set but max is not, set max to 1000
	if pg.SpecialCharacters && pg.SpecialMax == 0 {
		pg.SpecialMax = 1000
	}
	if pg.Numbers && pg.NumbersMax == 0 {
		pg.NumbersMax = 1000
	}
	if pg.Uppercase && pg.UppercaseMax == 0 {
		pg.UppercaseMax = 1000
	}
	if pg.Lowercase && pg.LowercaseMax == 0 {
		pg.LowercaseMax = 1000
	}

	// Generate password
	var password strings.Builder
	for i := 0; i < pg.Length; i++ {
		// Generate random character
		char := pg.generateRandomCharacter()

		// Add character to password
		password.WriteString(char)
	}

	return password.String()
}

func (pg *PasswordGenerator) generateRandomCharacter() string {
	// Generate random character
	var characters strings.Builder
	if pg.SpecialCharacters {
		if pg.SpecialMax > 0 {
			characters.WriteString(SPECIAL_CHARACTERS)
			pg.SpecialMax--
		}
	}
	if pg.Numbers {
		if pg.NumbersMax > 0 {
			characters.WriteString(NUMBERS)
			pg.NumbersMax--
		}
	}
	if pg.Uppercase {
		if pg.UppercaseMax > 0 {
			characters.WriteString(UPPERCASE)
			pg.UppercaseMax--
		}
	}
	if pg.Lowercase {
		if pg.LowercaseMax > 0 {
			characters.WriteString(LOWERCASE)
			pg.LowercaseMax--
		}
	}

	return selectRandomCharacter(characters.String())
}

func selectRandomCharacter(characters string) string {
	index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
	return string(characters[index.Int64()])
}
