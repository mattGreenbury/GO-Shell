package main

import (
	"bufio"         // Package bufio is used for buffered I/O operations.
	"fmt"           // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"os"            // Package os provides a platform-independent interface to operating system functionality.
	"os/exec"       // Package os/exec runs external commands.
	"path/filepath" // Package filepath implements utility routines for manipulating filename paths.
	"strings"       // Package strings implements simple functions to manipulate UTF-8 encoded strings.
)

// main is the entry point of the program.
// It reads user input from the command line and executes the corresponding commands.
// The valid commands are: exit, echo, type, pwd, and cd.
// If the command is not recognized, it is treated as an external command and executed accordingly.

func main() {
	// validCommands holds a map of commands that are recognized as valid internal commands.
	validCommands := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
		"pwd":  true,
		"cd":   true,
	}

	// reader is used to read input from the standard input.
	reader := bufio.NewReader(os.Stdin)
	for {
		// Prompt the user for input.
		fmt.Fprint(os.Stdout, "$ ")
		// Read the input until a newline character is encountered.
		input, _ := reader.ReadString('\n')
		// Trim leading and trailing whitespace from the input.
		input = strings.TrimSpace(input)
		// Split the input into parts based on spaces.
		parts := strings.Split(input, " ")
		// The first part of the input is considered the command.
		cmd := parts[0]

		// Check if the command is a valid internal command.
		if _, valid := validCommands[cmd]; valid {
			// Execute the corresponding action based on the command.
			switch cmd {
			case "exit":
				// Exit the program if the command is 'exit' or 'exit 0'.
				if len(parts) == 1 || (len(parts) == 2 && parts[1] == "0") {
					return
				} else {
					// Print an error message for invalid usage of the exit command.
					fmt.Println("Invalid usage of exit. Use 'exit' or 'exit 0'.")
				}
			case "echo":
				// Echo the message back to the user.
				message := strings.Join(parts[1:], " ")
				fmt.Println(message)
			case "type":
				// Check if the command is a built-in or an external command and display the result.
				if len(parts) < 2 {
					fmt.Println("type: usage: type command_name")
				} else {
					commandToCheck := parts[1]
					if _, exists := validCommands[commandToCheck]; exists {
						fmt.Printf("%s is a shell builtin\n", commandToCheck)
					} else {
						path, err := exec.LookPath(commandToCheck)
						if err != nil {
							fmt.Printf("%s: not found\n", commandToCheck)
						} else {
							fmt.Printf("%s is %s\n", commandToCheck, path)
						}
					}
				}
			case "pwd":
				// Print the current working directory.
				cwd, err := os.Getwd()
				if err != nil {
					fmt.Printf("pwd: %s\n", err)
				} else {
					fmt.Println(cwd)
				}
			case "cd":
				// Change the current working directory.
				if len(parts) < 2 {
					fmt.Println("cd: usage: cd <path>")
				} else {
					path := parts[1]
					// Handle the '~' character as the user's home directory.
					if path == "~" {
						var err error
						path, err = os.UserHomeDir()
						if err != nil {
							fmt.Printf("cd: failed to get home directory: %s\n", err)
							continue
						}
					} else if !filepath.IsAbs(path) {
						// Convert relative paths to absolute paths.
						cwd, _ := os.Getwd()
						path = filepath.Join(cwd, path)
					}
					// Attempt to change the directory.
					err := os.Chdir(path)
					if err != nil {
						errMsg := fmt.Sprintf("cd: %s: No such file or directory\n", path)
						fmt.Print(errMsg)
					}
				}
			}
		} else {
			// Execute the command as an external command if it is not recognized as an internal command.
			executeExternalCommand(parts)
		}
	}
}

// executeExternalCommand executes an external command by invoking the system's command interpreter.
func executeExternalCommand(parts []string) {
	// Create a new command based on the input parts.
	cmd := exec.Command(parts[0], parts[1:]...)
	// Set the standard input, output, and error streams for the command.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Run the command and check for errors.
	err := cmd.Run()
	if err != nil {
		// Print an error message if the command could not be found or executed.
		fmt.Printf("%s: command not found\n", parts[0])
	}
}
