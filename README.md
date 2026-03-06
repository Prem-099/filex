# filex

filex is a terminal-based file explorer written in GO.
It allows users to navigate directories,manage files,
and preview content directly from terminal using keyboard controls.

## Features

- Navigate directories from terminal
- Open files using system default apps
- File preview panel
- Search files
- Rename files and folders
- Delete files and folders
- Copy and Move files
- Recursive folder copy
- Sorting by name, size and time
- Toggle show hidden files like .vscode etc
- Help Screen
- Keyboard driven interface

## DEMO

## Installation

```bash
go install github.com/Prem-099/filex/cmd/filex@latest
```
- Add to path after installation

## Usage Examples

- Start in current directory:
```bash
filex
```

- Start in specific directory:
```bash
filex [path]
```

- Show version
```bash
filex --version
```

- Show help
```bash
filex --help
```

## Keybindings

| Key | Action |
|----|------|
| ↑ ↓ | Move selection |
| Enter | Open file / folder |
| Backspace | Go to parent directory |
| / | Search files |
| r | Rename file |
| d | Delete file |
| c | Copy |
| x | Cut |
| v | Paste |
| y | Clear clipboard |
| n | Sort by name |
| s | Sort by size |
| t | Sort by time |
| . | Toggle hidden files |
| ? | Show help |
| q | Quit |

## Project Structure

```bash
├───cmd                
│   └───filex         - CLI entry point
└───internal
    ├───app           - application logic
    ├───explorer      - navigation logic
    ├───fs            - filesystem absraction
    └───ui            - terminal rendering
```

## Requirements

GO 1.20 or later

## Note
- Best when used in cmd in full screen

## Current Issues
- can not open files which have spaces in-between words of file name
- no proper preview for binary files
- continuously renders terminal 

## Author

Prem Chandu Palivela

## License 

This project is licensed under MIT License
