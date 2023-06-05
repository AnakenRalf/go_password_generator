# Password Generator CLI Tool

The Password Generator CLI Tool is a command-line utility written in Golang that generates random passwords based on user-specified options. It allows users to define the desired password length and choose whether to include numbers, uppercase letters, and special characters in the generated password.

##Usage

    python password_generator.py [-h] [-l LENGTH] [-n] [-u] [-s]

## Options

    -h, --help: Show the help message and usage information.
    -l LENGTH, --length LENGTH: Specify the desired length of the password (default: 8).
    -n, --numbers: Include numbers in the generated password.
    -u, --uppercase: Include uppercase letters in the generated password.
    -s, --special: Include special characters in the generated password.

## Example Usage

Generate a password with a length of 12 characters, including numbers, uppercase letters, and special characters:

    password_generator.exe --length 12 --numbers --uppercase --special

or with short version of arguments:

    password_generator.exe -l 12 -n -u -s

This command will output a randomly generated password that meets the specified criteria.
