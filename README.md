# README for CCWC Project

## Overview

The CCWC ([Coding Challenges Word Count](https://codingchallenges.fyi/challenges/challenge-wc)) is a custom command-line interface (CLI) tool designed to implement various text file analysis functions, similar to the Unix `wc` command. This tool is developed as part of a series of coding challenges, aiming to demonstrate proficiency in Go programming, particularly in creating CLI applications and handling file and input/output operations.

## Features

CCWC supports several options to analyze the contents of text files or standard input (stdin). The tool can be used to count the following:

- Number of bytes (`-c` or `--bytes`)
- Number of lines (`-l` or `--lines`)
- Number of words (`-w` or `--words`)
- Number of characters (`-m` or `--characters`)

The program can read input from a specified file or from stdin, making it versatile for use in various scenarios such as piping output from other commands.

## Installation

To install CCWC, you will need to have Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/).

Once Go is installed, you can build the CCWC tool using the following steps:

1. Clone the repository to your local machine.
2. Navigate to the root directory of the project.
3. Run `go build`. This will compile the source code and generate an executable.

## Usage

After building the project, you can use CCWC by running the executable from the command line. The basic syntax is as follows:

```bash
./ccwc [options] [file]
```

If no file is specified, CCWC will read from stdin.

### Options

- `-c` or `--bytes`: Count the number of bytes.
- `-l` or `--lines`: Count the number of lines.
- `-w` or `--words`: Count the number of words.
- `-m` or `--characters`: Count the number of characters.

### Examples

Counting the number of words in a file:

```bash
./ccwc -w myfile.txt
```

Counting the number of lines from stdin:

```bash
cat myfile.txt | ./ccwc -l
```

## Contributing

Contributions to the CCWC project are welcome. If you have suggestions for improvements or encounter any issues, please feel free to open an issue or submit a pull request on the project repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Note**: This README is for a project created as part of coding challenges. The actual functionality and behavior of the tool might vary based on the specific requirements and implementation details of those challenges.
