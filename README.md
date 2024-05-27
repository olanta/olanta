# Olanta

Olanta is a static code analysis tool designed to detect code smells and bugs. It uses YAML files for defining rules and Bleve for indexing.

## Supported Languages

- Java
- Python

### Prerequisites

Make sure you have the following installed:
- Go 1.22 or later
- Makefile

### Building the Project

To build Olanta, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/olanta/olanta.git
    cd olanta
    ```

2. Run the build script:
    ```sh
    go run build/main.go -d build
    ```

    This command compiles the Olanta project and prepares it for use.

### Running the Tool

To run Olanta, use the following command:

```sh
./olanta scan [language] [path]
```

Replace `[language]` with either `java` or `python`, and `[path]` with the path to the code you want to analyze. For example:

```sh
./olanta scan java /path/to/java/project
```

### Adding New Rules

You can add new rules by editing the YAML files in the `rules` directory. For example, to add a new Java rule, update the `java_rules.yaml` file.
