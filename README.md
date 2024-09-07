# Org - File Organizer CLI Tool

`Org` is a command-line tool that helps organize files by sorting them into folders based on their extensions. It provides an easy and efficient way to categorize files within a directory, making them easier to find and manage.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
  - [`category`](#category)
  - [`add`](#add)
  - [`remove`](#remove)
  - [`organize`](#organize)
- [Flags](#flags)
- [Code Structure](#code-structure)
- [How It Works](#how-it-works)
- [Contributing](#contributing)
- [License](#license)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/username/FileOrganizer.git
   cd FileOrganizer
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the executable:
   ```bash
   go build -o org
   ```

4. Run the tool:
   ```bash
   ./org --help
   ```

## Usage

You can use the CLI tool to perform various operations like listing file categories, adding a new category, removing categories, and organizing files in a given directory.

### General Syntax

```bash
org [command] [flags] [arguments]
```

### Example Usage

- To list all categories:
  ```bash
  org category
  ```

- To add a category:
  ```bash
  org add documents --ext ".pdf,.docx" --folder "Documents"
  ```

- To organize a directory:
  ```bash
  org organize /path/to/directory
  ```

## Commands

### `category`

Lists all file categories currently available for organizing files. It can optionally display associated file extensions with the `--verbose` flag.

```bash
org category --verbose
```

---

### `add`

Adds a new category with a list of associated file extensions. Optionally, specify a folder name and extensions for the category.

```bash
org add [category] --ext [extensions] --folder [folder_name]
```

Example:
```bash
org add documents
```
This creates a category if it doesn't exist already. The folder name defaults to the category name and the extensions list is empty
<br>


```bash
org add documents -folder [folder_name]
```
This creates a category if it doesn't exist already and sets the folder name   
<br>

```bash
org add documents --ext ".pdf, .docx"
```
OR
```bash
org add documents --ext ".pdf" --ext ".docx"
```
This creates a category if it doesn't exist and adds a list of extensions to it
<br>

> **Note:**
> Also, you can use it to remove extensions you don't want in the category by simply appending a dash
> ```bash
> org add documents --ext "-.pdf"
> ```


---

### `remove`

Removes a file category from the list of available categories.

```bash
org remove [category]
```

Example:
```bash
org remove documents
```

---

### `organize`

Organizes the files in a specified directory by sorting them into folders based on the defined categories and extensions.

```bash
org organize [directory]
```

Example:
```bash
org organize /path/to/directory
```

## Flags

| Flag         | Description                                                        | Example                         |
|--------------|--------------------------------------------------------------------|---------------------------------|
| `-v`, `--verbose` | Displays detailed information when listing categories.         | `org category --verbose`        |
| `-f`, `--folder`  | Specifies the folder name for a new category.                   | `org add docs --folder Docs`    |
| `-e`, `--ext`     | A list of file extensions associated with a category.           | `org add docs --ext ".pdf,.doc"`|

---

## Code Structure

The project is organized into several packages and files for ease of maintenance and modularity:

```
.
├── cmd/
│   ├── root.go                # Entry point for Cobra commands
│   ├── handlers.go            # Contains commands for using the cli tool
├── constants/
|   ├── constants.go           # Stores constants used in the project
├── models/
│   ├── models.go              # Defines the FileCategory struct and methods related to file categories. Also defines the DataStore interface
│   ├── store_file.go          # Implements DataStore and methods for managing categories and extensions using a file for storage
├── main.go                    # Main entry point for the program
├── README.md                  # Project documentation
```

### `cmd/` Directory

- **root.go**: Contains the root command (`org`) and links all subcommands (like `category`, `add`, `remove`, etc.).
- **handlers.go**: Implements commands for listing, adding, and removing file categories.

### `models/` Directory

- **store.go**: Implements storage and retrieval of categories and associated extensions, allowing the CLI to manage file organization logic.
- **store_file.go**: Implements DataStore and methods for managing categories and extensions using a file for storage

## How It Works

### Categories and Extensions

- The tool uses categories to organize files. Each category is associated with a list of file extensions (e.g., `.pdf`, `.docx`, `.jpg`).
- Categories are stored internally in a data structure and can be added, removed, or listed using the corresponding commands.

### Organizing Files

1. **Directory Scan**: When you run `org organize [directory]`, the tool scans the given directory for files.
2. **File Matching**: Each file's extension is checked against the available categories.
3. **File Movement**: Files are moved into corresponding folders based on their category. If a folder for the category does not exist, it is created.

### Example Workflow

1. **Adding a Category**:
   - You can add a category for documents like PDFs and DOCX files by running:
     ```bash
     org add documents --ext ".pdf,.docx" --folder "Documents"
     ```

2. **Organizing Files**:
   - After adding categories, you can organize files in a directory by running:
     ```bash
     org organize /path/to/directory
     ```

3. **Category Listing**:
   - You can list all available categories and their associated extensions using:
     ```bash
     org category --verbose
     ```

---

## Contributing

We welcome contributions! Feel free to fork the repository and submit a pull request, or open an issue to report bugs or suggest improvements.

### Steps to Contribute

1. Fork the repository
2. Create a new feature branch (`git checkout -b feature/feature-name`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/feature-name`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

By following this structure, the tool becomes highly user-friendly and maintainable, with clear explanations of how to use it and how it works under the hood.

