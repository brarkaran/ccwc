/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

type cliOptions struct {
	bytes bool
	lines bool
	words bool
	chars bool
}

type fileStats struct {
	bytes uint64
	lines uint64
	words uint64
	chars uint64
}

var options cliOptions

func formatStats(options *cliOptions, stats fileStats) string {
	var result []string
	if options.bytes {
		result = append(result, strconv.FormatUint(stats.bytes, 10))
	}
	if options.lines {
		result = append(result, strconv.FormatUint(stats.lines, 10))
	}
	if options.words {
		result = append(result, strconv.FormatUint(stats.words, 10))
	}
	if options.chars {
		result = append(result, strconv.FormatUint(stats.chars, 10))
	}
	return strings.Join(result, " ")
}

func computeStats(buf *bufio.Reader) fileStats {
	var prevChar rune
	var bytesCount uint64
	var linesCount uint64
	var wordsCount uint64
	var charsCount uint64

	for {
		char, bytes, err := buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				if prevChar != rune(0) && !unicode.IsSpace(prevChar) {
					wordsCount++
				}
				break
			}
			log.Fatal(err)
		}
		bytesCount += uint64(bytes)
		charsCount++
		if char == '\n' {
			linesCount++
		}
		if !unicode.IsSpace(prevChar) && unicode.IsSpace(char) {
			wordsCount++
		}
		prevChar = char
	}
	return fileStats{bytes: bytesCount, lines: linesCount, words: wordsCount, chars: charsCount}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "ccwc [file]",
	Args: cobra.MinimumNArgs(0),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if !options.bytes && !options.chars && !options.words && !options.lines {
			options.bytes = true
			options.chars = true
			options.words = true
			options.lines = true
		}
		if len(args) == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("%s\n", formatStats(&options, computeStats(reader)))
		} else {
			for _, filename := range args {
				file, err := os.Open(filename)
				if err != nil {
					log.Fatal(err)
				}
				reader := bufio.NewReader(file)
				fmt.Printf("%s %s\n", formatStats(&options, computeStats(reader)), filename)
			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&options.bytes, "bytes", "c", false, "Count the number of bytes in the file")
	rootCmd.Flags().BoolVarP(&options.lines, "lines", "l", false, "Count the number of lines in the file")
	rootCmd.Flags().BoolVarP(&options.words, "words", "w", false, "Count the number of words in the file")
	rootCmd.Flags().BoolVarP(&options.lines, "characters", "m", false, "Count the number of characters in the file")
}
