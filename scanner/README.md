# Olanta Scanner

Olanta Scanner is a static code analysis tool designed to find code smells and bugs in Java and Python projects. It uses embedded YAML rules to detect issues in the code and can either submit these issues to the Olanta Core server or print them to the console.

## Features

- Detects code smells and bugs in Java and Python code.
- Uses embedded YAML rules for flexibility and ease of customization.
- Can submit issues to the Olanta Core server or print them to the console.

## Installation

To build the Olanta Scanner, you need to have Go installed on your machine. Follow these steps to build and run the scanner:

1. Clone the repository:
    ```sh
    git clone https://github.com/olanta/olanta.git
    cd olanta/scanner
    ```

2. Install the required Go dependencies:
    ```sh
    go mod tidy
    ```

3. Build the project using Mage:
    ```sh
    mage -d build build
    ```

## Usage

To use the Olanta Scanner, run the built binary with the `scan` command, specifying the language and the path to the code to be analyzed.

### Example

1. Scan a Java project:
    ```sh
    ./bin/scanner scan java /path/to/your/java/project
    ```

2. Scan a Python project:
    ```sh
    ./bin/scanner scan python /path/to/your/python/project
    ```

### Command-line Options

- `--core-url`: Specify the URL of the Olanta Core server to submit the issues. If the server is not available, the issues will be printed to the console.

## Development

### Project Structure

- `cmd/scanner/main.go`: The main entry point for the scanner command-line tool.
- `internal/scanner`: Contains the core scanning logic and language-specific scanners.
- `internal/models`: Defines the data models used in the project.
- `internal/utils`: Utility functions for loading YAML rules and other common tasks.
- `internal/scanner/rules`: Contains the embedded YAML rules for Java and Python.
- `build`: Contains build scripts and configuration for Mage.

### Building and Running

1. Clean the build:
    ```sh
    mage -d build clean
    ```

2. Build the project:
    ```sh
    mage -d build build
    ```

3. Run the project:
    ```sh
    ./bin/scanner scan java /path/to/your/java/project
    ```

