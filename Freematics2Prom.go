package main

import (
	"fmt"
	"github.com/tomhobson/freematics2prom/internal/csv_parser"
	"github.com/tomhobson/freematics2prom/internal/file_store"
	"os"
)

func main() {
	// Check if the command-line argument (file path) is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}

	// Extract the file path from the command-line argument
	filePath := os.Args[1]

	// Create an instance of the FileStore
	fileStoreInstance := file_store.NewFileStore()
	// Create an instance of the csvParser as well
	csvParserInstance := csv_parser.NewCsvParser()

	// Call the ReadFile method to read the content of the file
	fileContent, err := fileStoreInstance.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	data, err := csvParserInstance.ParseCSV(fileContent)

	for _, row := range data {
		fmt.Printf("%s, %v\n", row.Name, row.Value)
	}
}
