# Correa CLI

This is a CLI tool for simple tasks that I do everyday.

## Table of Contents

- [Correa CLI](#correa-cli)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Windows](#windows)
    - [Linux / MacOS](#linux--macos)
  - [Building](#building)
    - [Cross-compilation](#cross-compilation)
  - [Contributing](#contributing)

## Installation

### Windows

1. Download the Windows artifact
2. Extract the contents on any folder
3. Create a .bat file named `correa.bat` with the following content:

```bat
@echo off
"<path to artifact>" %*
```

4. Add the folder to the PATH environment variable
5. Run `correa` on the command line

### Linux / MacOS

1. Download the Linux artifact
2. Extract the contents on any folder
3. Add the folder to the PATH environment variable
4. Run `correa` on the command line
5. If you get a permission denied error, run `chmod +x correa`

## Building

To build the project, run the following command:

```sh
go build -o correa
```

This will create an executable file named `correa` in the current directory.

### Cross-compilation

To build the project for a different operating system or architecture, set the `GOOS` and `GOARCH` environment variables before running the build command. For example, to build for Windows on a 64-bit architecture:

```sh
GOOS=windows GOARCH=amd64 go build -o correa.exe
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.