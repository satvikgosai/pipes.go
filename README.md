# Pipes Screen Saver

A terminal-based screen saver that creates an animated pipe maze effect. Built with Go.

## Features

- Smooth pipe animation
- Adjustable speed
- Terminal window resize support
- Clean exit handling
- Cross-platform support

## Installation

### Prerequisites

- Go 1.24 or later
- A terminal emulator

### Building from Source

```bash
# Clone the repository
git clone https://github.com/satvikgosai/pipes.git
cd pipes

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
go install github.com/satvikgosai/pipes@latest
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

## Development

### Project Structure

```
pipes/
├── cmd.go         # Command-line interface
├── config.go      # Configuration management
├── matrix.go      # Pipe animation logic
├── terminal.go    # Terminal handling
├── main.go        # Application entry point
└── README.md      # Documentation
```

### Building for Development

```bash
# Build with debug symbols
go build -gcflags="-N -l" -o pipes
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
