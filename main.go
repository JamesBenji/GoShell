package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// "bufio" to access a Scanner that can read the standard input provided by the "os" module
// "strings" allows the user input to be split into individual words that can be interpreted later
// A Read-Evaluate-Print-Loop will be implemented

const shellName = "Cloud1"

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s> ", shellName)

		if !input.Scan() {
			break
		}

		words := strings.Fields(input.Text())
		if len(words) == 0 {
			continue
		}

		pipeIndex := -1
		for index, word := range words {
			if word == "|" {
				pipeIndex = index
				break
			}
		}

		if pipeIndex >= 0 {
			if pipeIndex == 0 || pipeIndex == len(words)-1 {
				fmt.Fprintln(os.Stderr, "usage: command | command")
				continue
			}

			left := words[:pipeIndex]
			right := words[pipeIndex+1:]

			reader, writer, err := os.Pipe()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			first := exec.Command(left[0], left[1:]...)
			first.Stdin = os.Stdin
			first.Stdout = writer
			first.Stderr = os.Stderr

			second := exec.Command(right[0], right[1:]...)
			second.Stdin = reader
			second.Stdout = os.Stdout
			second.Stderr = os.Stderr

			if err := first.Start(); err != nil {
				writer.Close()
				reader.Close()
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			writer.Close()
			if err := second.Run(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			reader.Close()

			if err := first.Wait(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			continue
		}

		if words[0] == "exit" {
			break
		}

		if words[0] == "cd" {
			if len(words) != 2 {
				fmt.Fprintln(os.Stderr, "usage: cd <directory>")
				continue
			}

			if err := os.Chdir(words[1]); err != nil {
				fmt.Fprintln(os.Stderr, "cd:", err)
			}
			continue
		}

		redirectIndex := -1
		for index, word := range words {
			if word == ">" || word == "<" {
				redirectIndex = index
				break
			}
		}

		output := os.Stdout
		inputFile := os.Stdin

		if redirectIndex >= 0 {
			if redirectIndex == 0 || redirectIndex != len(words)-2 {
				fmt.Fprintln(os.Stderr, "usage: command < filename or usage: command > filename")
				continue
			}

			filename := words[redirectIndex+1]

			var err error
			if words[redirectIndex] == "<" {
				inputFile, err = os.Open(filename)
			} else {
				output, err = os.Create(filename)
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			words = words[:redirectIndex]
		}

		command := exec.Command(words[0], words[1:]...)
		command.Stdin = inputFile
		command.Stdout = output
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if inputFile != os.Stdin {
			inputFile.Close()
		}

		if output != os.Stdout {
			output.Close()
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read input:", err)
	}
}
