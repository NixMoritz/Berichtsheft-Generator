package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func getWeekDates(year int, month time.Month, day int) (int, string, string, map[string]string) {
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	_, isoWeek := date.ISOWeek()

	// Calculate the start and end dates of the week
	monday := date.AddDate(0, 0, -int(date.Weekday())+1)
	sunday := monday.AddDate(0, 0, 6)

	// Details for each day of the week
	days := map[string]string{
		"Montag":     monday.Format("02.01.2006"),
		"Dienstag":   monday.AddDate(0, 0, 1).Format("02.01.2006"),
		"Mittwoch":   monday.AddDate(0, 0, 2).Format("02.01.2006"),
		"Donnerstag": monday.AddDate(0, 0, 3).Format("02.01.2006"),
		"Freitag":    monday.AddDate(0, 0, 4).Format("02.01.2006"),
	}

	return isoWeek, monday.Format("02.01"), sunday.Format("02.01.2006"), days
}

func generateAndWriteFileNames(year int, outputFolder string) error {
	for month := time.January; month <= time.December; month++ {
		// Create a file for each week in the month
		for day := 1; day <= 31; day += 7 {
			// Skip if the date is invalid (e.g., February 30th)
			if time.Date(year, month, day, 0, 0, 0, 0, time.UTC).IsZero() {
				continue
			}

			isoWeek, startDate, endDate, days := getWeekDates(year, month, day)
			fileName := fmt.Sprintf("KW%d_%s-%s.txt", isoWeek, startDate, endDate)
			// Derive the file path relative to the output folder
			filePath := filepath.Join(outputFolder, fileName)

			file, err := os.Create(filePath)
			if err != nil {
				return err
			}

			// Insert the template with actual dates for each day into the file
			template := fmt.Sprintf(`Montag, %s

-   

Dienstag, %s
    
-   

Mittwoch, %s

-   

Donnerstag, %s

-   

Freitag, %s

-   

_________________________________________________________________________________________________________________________

Berufsschule:

Montag, %s

-   

Mittwoch, %s

-   

_________________________________________________________________________________________________________________________

Thema:`, days["Montag"], days["Dienstag"], days["Mittwoch"], days["Donnerstag"], days["Freitag"], days["Montag"], days["Mittwoch"])

			_, err = file.WriteString(template)
			if err != nil {
				return err
			}

			file.Close()
		}
	}
	return nil
}

func main() {
	var year int
	var createNewFolder bool
	var outputFolder string

	// Check if command-line arguments are provided
	if len(os.Args) > 1 {
		// Parse the provided year argument
		argYear, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error parsing the year argument:", err)
			return
		}
		year = argYear

		// Check if a second argument is provided and is "y" or "1"
		if len(os.Args) > 2 && (os.Args[2] == "y" || os.Args[2] == "1") {
			createNewFolder = true
		}

		// Check if a third argument is provided as the output directory
		if len(os.Args) > 3 {
			outputFolder = os.Args[3]
		}
	} else {
		// If no argument is provided, use the current year
		year = time.Now().Year()
	}

	// Get the current working directory if outputFolder is not specified
	if outputFolder == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
			return
		}
		outputFolder = currentDir
	}

	// If createNewFolder is true, create a new folder based on the provided year
	if createNewFolder {
		outputFolder = filepath.Join(outputFolder, fmt.Sprintf("Berichtshefte_%d", year))
		err := os.Mkdir(outputFolder, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating a new folder:", err)
			return
		}
	}

	// Example usage
	err := generateAndWriteFileNames(year, outputFolder)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Your Berichtshefte have been generated.\nYear used for operation: %d\nThe Berichtshefte are located at: %s\n", year, outputFolder)
}
