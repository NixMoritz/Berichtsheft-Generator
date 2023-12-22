# Berichtsheft Generator

## Overview

Berichtsheft Generator is a versatile program written in GoLang. It is designed to generate weekly reports (Berichtshefte) for a specified year, creating individual text files for each week with specific templates.
## Usage

### Prerequisites

Make sure you have [Go](https://golang.org/) installed on your machine.

### Using the `main.go` Variant

1. **Adjusting Parameters:**
   Modify the parameters in the `main` function of the `main.go` file:
   - `year`: Set the default year for report generation.
   - `createNewFolder`: Set to true to create a new folder for the reports.
   - `outputFolder`: Set the default output folder path.

2. **Run the `main.go` Variant:**
   Execute the `main.go` file using the Go compiler:
   ```bash
   go run main.go
   ```
You can also provide arguments when calling the program:
   ```bash
   go run main.go [year] [createNewFolder] [outputFolder]
   ```
Alternatively, build the executable:
   ```bash
   go build main.go
   ```
Run the generated executable.

3. **Example:**
   ```bash
   go run main.go 2024 y C:\Reports
   ```

### Running the Executable

1. **Run the Executable:**
   Execute the `Berichtsheft-Generator.exe` binary in your command prompt or terminal.

2. **Command-Line Arguments:**
   The program accepts optional command-line arguments:
    - `Year`: Specify the year for report generation.
    - `Create New Folder`: Optionally, provide "y" or "1" to create a new folder for the reports.
    - `Output Folder`: Optionally, provide the path to the output folder as the third argument.

3. **Example:**
   ```bash
   Berichtsheft-Generator.exe 2024 y C:\Reports
   ```

## Notes

- If no command-line arguments are provided, the program defaults to the current year.
- If no output folder is specified, the program uses the current working directory.
- If the second argument is neither "y" nor "1", files will be saved in the current working directory or the output directory given in argument 3.

## Output

The generated reports are saved as text files in the specified output folder. Each file follows the naming convention: `KW{ISOWeek}_{StartDate}-{EndDate}.txt`.

## Template

The content of each report file follows a template with entries for each weekday, details about vocational school days, and a thematic section.

```plaintext
Montag, xx.xx.xx

-   

Dienstag, xx.xx.xx
    
-   

Mittwoch, xx.xx.xx

-   

Donnerstag, xx.xx.xx

-   

Freitag, xx.xx.xx

-   

_________________________________________________________________________________________________________________________

Berufsschule:

Montag, xx.xx.xx

-   

Mittwoch, xx.xx.xx

-   

_________________________________________________________________________________________________________________________

Thema:
```

Feel free to customize the template based on your reporting needs.

## License

This program is open-source and distributed under the [MIT License](LICENSE). Feel free to use, modify, and share it. If you find any issues or have suggestions, please create an [issue](https://github.com/NixMoritz/Berichtsheft-Generator/issues).
