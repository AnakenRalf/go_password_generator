package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	lenght               int
	useNumbers           bool
	useUpperCase         bool
	useSpecialCharacters bool
	showHelp             bool
)

const (
	lowercase         = "abcdefghijklmnopqrstuvwxyz"
	uppercase         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers           = "0123456789"
	specialCharacters = "!@#$%^&*()_+-=[]{}|;:,./<>?"
)

func init() {
	// Define flags
	flag.IntVar(&lenght, "l", 8, "length of password")
	flag.IntVar(&lenght, "lenght", 8, "length of password")
	flag.BoolVar(&useNumbers, "n", false, "use numbers")
	flag.BoolVar(&useNumbers, "numbers", false, "use numbers")
	flag.BoolVar(&useUpperCase, "u", false, "use uppercase")
	flag.BoolVar(&useUpperCase, "uppercase", false, "use uppercase")
	flag.BoolVar(&useSpecialCharacters, "s", false, "use special characters")
	flag.BoolVar(&useSpecialCharacters, "special", false, "use special characters")
	flag.BoolVar(&showHelp, "h", false, "Show help message")
	flag.BoolVar(&showHelp, "help", false, "Show help message")

	flag.Parse()
}

func main() {
	if showHelp || len(os.Args) == 1 {
		printHelp()
		return
	}

	rand.Seed(time.Now().UnixNano())

	var passwordBuilder strings.Builder

	passwordBuilder.WriteString(lowercase)

	if useNumbers {
		passwordBuilder.WriteString(numbers)
	}

	if useUpperCase {
		passwordBuilder.WriteString(uppercase)
	}

	if useSpecialCharacters {
		passwordBuilder.WriteString(specialCharacters)
	}

	password := generateRandomPassword(passwordBuilder.String(), lenght)
	fmt.Println(password)

}

func generateRandomPassword(lowercase string, length int) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = lowercase[rand.Intn(len(lowercase))]
	}
	return string(password)
}

func printHelp() {
	fmt.Println("Password Generator CLI Tool")
	fmt.Println("Usage: password-generator [options]")
	fmt.Println("Options:")
	fmt.Println("\t-l  --lenght       Specify the desired length of the password (default: 8).")
	fmt.Println("\t-n  --numbers      Include numbers in the generated password.")
	fmt.Println("\t-u  --uppercase    Include uppercase letters in the generated password")
	fmt.Println("\t-s  --special      Include special characters in the generated password.")
	fmt.Println("\t-h  --help         Show the help message and usage information")
}
