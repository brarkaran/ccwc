/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var countBytes bool
var countLines bool
var countWords bool
var countChars bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "ccwc [file]",
	Args: cobra.MinimumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		if countBytes {
			fileInfo, err := os.Stat(fileName)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d %s\n", fileInfo.Size(), fileName)
		}
		if countLines {
			file, err := os.Open(fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			fileScanner := bufio.NewScanner(file)
			lineCount := 0
			for fileScanner.Scan() {
				lineCount++
			}
			fmt.Printf("%d %s\n", lineCount, fileName)
		}
		if countWords {
			file, err := os.Open(fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			fileScanner := bufio.NewScanner(file)
			wordCount := 0
			for fileScanner.Scan() {
				line := fileScanner.Text()
				words := strings.Fields(line)
				wordCount += len(words)
			}
			fmt.Printf("%d %s\n", wordCount, fileName)
		}
		if countChars {
			file, err := os.Open(fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			fileScanner := bufio.NewScanner(file)
			fileScanner.Split(bufio.ScanRunes)
			charCount := 0
			for fileScanner.Scan() {
				charCount++
			}
			fmt.Printf("%d %s\n", charCount, fileName)
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
	rootCmd.Flags().BoolVarP(&countBytes, "bytes", "c", false, "Count the number of bytes in the file")
	rootCmd.Flags().BoolVarP(&countLines, "lines", "l", false, "Count the number of lines in the file")
	rootCmd.Flags().BoolVarP(&countWords, "words", "w", false, "Count the number of words in the file")
	rootCmd.Flags().BoolVarP(&countChars, "characters", "m", false, "Count the number of characters in the file")
}
