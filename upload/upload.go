package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the file to a temporary location
	tempDir := "./temp"
	os.MkdirAll(tempDir, os.ModePerm)
	tempFilePath := filepath.Join(tempDir, fileHeader.Filename)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Copy the file contents to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Error copying file contents", http.StatusInternalServerError)
		return
	}

	// Open the Excel file
	f, err := excelize.OpenFile(tempFilePath)
	if err != nil {
		http.Error(w, "Error opening Excel file", http.StatusInternalServerError)
		return
	}

	// Get all the sheets in the Excel file
	sheetList := f.GetSheetMap()
	for _, sheet := range sheetList {
		// Read rows from the current sheet
		rows, err := f.GetRows(sheet)
		if err != nil {
			http.Error(w, "Error reading Excel sheet", http.StatusInternalServerError)
			return
		}

		// Process each row
		for _, row := range rows {
			// Print each cell value in the row
			for _, cell := range row {
				fmt.Printf("%s\t", cell)
			}
			fmt.Println()
		}
	}

	// Remove the temporary file after reading
	err = os.Remove(tempFilePath)
	if err != nil {
		log.Println("Error removing temporary file:", err)
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded, read, and removed successfully!"))
}
