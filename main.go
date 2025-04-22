package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rtanx/caesarcipher/cacipher"
	"github.com/spf13/pflag"
)

func main() {
	var decFlag bool
	var shift int
	var inputFile string
	var outputFile string
	var showHelp bool

	pflag.BoolVarP(&decFlag, "decrypt", "d", false, "Perform decryption instead of encryption to the input text")
	pflag.IntVarP(&shift, "shift", "s", 13, "Shift value")
	pflag.StringVarP(&inputFile, "in", "i", "", "Input file (default: stdin or command-line argument)")
	pflag.StringVarP(&outputFile, "out", "o", "", "Output file (default: stdout)")
	pflag.BoolVarP(&showHelp, "help", "h", false, "Show help")

	pflag.Parse()

	if showHelp || (len(os.Args) == 0) {
		fmt.Println("Caesar Cipher Tool - Encrypt or decrypt text using ROT-N cipher")
		fmt.Println("\nUsage:")
		pflag.PrintDefaults()
		os.Exit(0)
	}

	// Setup output
	var outputWriter *os.File
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		outputWriter = file
		defer outputWriter.Close()
	} else {
		outputWriter = os.Stdout
	}

	// Create a cipher instance
	c := cacipher.NewCipher(shift, nil, outputWriter, decFlag)

	// Handle input based on what's provided
	// Case 1: Input file specified
	if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		c.SetReader(file)
		if err := c.Transform(); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing file: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Case 2: Text provided as command-line arguments
	args := pflag.Args()
	if len(args) > 0 {
		inputText := strings.Join(args, " ")
		if err := c.TransformText(inputText); err != nil {
			fmt.Fprintf(os.Stderr, "Error processing text: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Case 3: No input file or text arguments, use stdin
	fmt.Print("Enter text (Ctrl+D to finish):\n\n")

	c.SetReader(os.Stdin)
	if err := c.Transform(); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing input: %v\n", err)
		os.Exit(1)
	}

}
