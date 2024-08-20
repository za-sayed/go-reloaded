package main

import (
	"fmt"
	"go-reloaded/functions"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: go run . <input-file> <output-file>")
		return
	}

	inputFileName := args[0]
	outputFileName := args[1]
	if filepath.Ext(inputFileName) != ".txt" || filepath.Ext(outputFileName) != ".txt" {
		fmt.Println("Error: Both files must have a .txt extension.")
		return
	}

	data, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer data.Close()

	fileContent, err := io.ReadAll(data)
	if err != nil {
		fmt.Println("Error reading file content:", err)
		return
	}

	for _, char := range fileContent {
		if char > 127 {
			fmt.Println("Error: The file contains non-ASCII characters.")
			return
		}
	}

	fileContentStr := string(fileContent)
	fileContentStr = strings.Join(strings.Fields(fileContentStr), " ")
	fileContentStr = functions.FormatPunctuation(fileContentStr)
	fileContentStr = functions.FormatQuotation(fileContentStr)
	fileContentStr = functions.ConvertAToAn(fileContentStr)
	fileContentStr = functions.ApplyTransformations(fileContentStr)

	err = os.WriteFile(outputFileName, []byte(fileContentStr), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
}
