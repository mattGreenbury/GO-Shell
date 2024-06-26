# Command Line Interface Tool

This Go program serves as a basic command line interface tool that allows users to execute a set of predefined internal commands as well as external commands directly from the command line. It is designed to provide a simple interface for performing common operations without the need for a full shell environment.

## Features

- **Internal Commands**: Supports a set of internal commands including:
  - `exit`: Exits the application.
  - `echo`: Prints text to the standard output.
  - `type`: Displays the content of a file.
  - `pwd`: Prints the current working directory.
  - `cd`: Changes the current directory to the one specified.
- **External Commands**: Any command not recognized as an internal command is treated as an external command and executed accordingly.

## Installation

To install and run this CLI tool, follow these steps:

1. Ensure you have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).
2. Clone this repository to your local machine.
3. Navigate to the directory containing the cloned repository.
4. Build the program by running `go build -o go-shell.exe`.
5. Execute the built program by running `./go-shell` on Unix-like systems or `go-shell.exe` on Windows.

## Usage

After starting the program, you will be presented with a prompt (`$ `). You can enter any of the supported internal commands or any external command available in your system's PATH.

## Example Usage

This section demonstrates how to use the CLI tool with various commands.

- **Echo Command**
$ echo Hello, World! 
Hello, World!

- **Print Working Directory (pwd)**
$ pwd /home/user

- **Change Directory (cd) to an Absolute Path**
$ cd /path/to/directory

- **Change Directory (cd) to a Relative Path**
$ cd relativeDirectory

- **Change to Home Directory**
$ cd ~

- **Display File Content (type)**
$ type filename.txt Contents of filename.txt

- **Exit the CLI Tool**
$ exit

## Contributing

Contributions to this project are welcome. Please feel free to fork the repository, make your changes, and submit a pull request.

## License

This project is open-sourced under the MIT License. See the LICENSE file for more details.