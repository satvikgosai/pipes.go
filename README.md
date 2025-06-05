# Pipes Screen Saver

A terminal-based screen saver that creates an animated pipe maze effect, built with Go.

## Features

- Smooth pipe animation
- Adjustable speed
- Terminal window resize support
- Clean exit handling
- Cross-platform support
- Multiple pipe styles
- Color themes

## Installation

### Prerequisites

- Go 1.24 or later
- A terminal emulator

### Building from Source

```bash
# Clone the repository
git clone https://github.com/satvikgosai/pipes.go.git
cd pipes.go

# Build the binary
go build -o pipes
```

### Manual Installation (Optional)

```bash
# Install to /usr/local/bin (requires sudo on some systems)
sudo install pipes /usr/local/bin/
```

### Installation via go install

```bash
# Install directly using go install
go install github.com/satvikgosai/pipes.go@latest
```

## Usage

### Basic Usage

```bash
./pipes
```

### Adjusting Speed

The speed can be adjusted from 0 (slowest) to 100 (fastest):

```bash
# Using the long form
./pipes --speed 75

# Using the short form
./pipes -s 75
```

### Exiting

Press `Ctrl+C` to exit the screen saver.

## Configuration

The following parameters can be configured:

- `--speed` or `-s`: Animation speed (0-100)
  - Default: 50
  - Higher values = faster animation
- `--theme` or `-t`: Color theme
  - Default: default
  - Options: default, red, green, blue, cyan, magenta, yellow, rainbow
- `--style` or `-l`: Pipe style
  - Default: default
  - Options: default, single, thick, rounded, dotted

### Examples

```bash
# Using rainbow theme with thick pipes
./pipes --theme rainbow --style thick

# Using blue theme with rounded pipes at high speed
./pipes -t blue -l rounded -s 80
```

## Development

### Project Structure

```
pipes.go/
├── cmd.go         # Command-line interface
├── config.go      # Configuration management
├── matrix.go      # Pipe animation logic
├── terminal.go    # Terminal handling
├── main.go        # Application entry point
└── README.md      # Documentation
```

### Purpose

This project is for learning purposes on the basis of build your own X philosophy and is inspired by https://github.com/pipeseroni/pipes.sh 

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
