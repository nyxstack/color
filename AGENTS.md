# AGENTS.md

This file provides context for AI agents and language models working with the NYX Color package.

## Package Overview

**NYX Color** is a lightweight, zero-dependency Go package for terminal colors and text styling using ANSI escape sequences.

## Key Components

### Core File: `color.go`
- **Constants**: Pre-defined ANSI escape sequences for colors and styles
- **Functions**: `RGB()`, `Hex()`, `Color256()` for dynamic color generation
- **Features**: 16 standard colors, bright variants, backgrounds, text styles, 256-color palette, 24-bit RGB

### Test Suite: `color_test.go`
- Comprehensive unit tests for all functions
- Benchmarks for performance testing
- Edge case testing (invalid inputs, boundary conditions)

### Examples Directory
- `basic.go` - Standard colors and text styling
- `color_256.go` - 256-color palette demonstrations
- `advanced.go` - RGB, HEX, gradients, and complex examples
- `demo.go` - Showcase file for screenshots

## API Reference

### Colors (Constants)
```go
color.Red, color.Green, color.Blue, color.Yellow, color.Magenta, color.Cyan, color.White, color.Black
color.BrightRed, color.BrightGreen, color.BrightBlue // ... etc
color.BgRed, color.BgGreen, color.BgBlue // ... etc (backgrounds)
color.Reset // Always use after colored text
```

### Text Styles (Constants)
```go
color.Bold, color.Italic, color.Underline, color.Strikethrough
color.Dim, color.Reverse, color.Hidden, color.Blink
```

### Dynamic Functions
```go
color.RGB(r, g, b uint8) string        // 24-bit RGB colors
color.Hex(hex string) string           // HEX colors (#FF0000 or FF0000)
color.Color256(i uint8) string         // 256-color palette (0-255)
```

## Usage Patterns

### Basic Pattern
```go
fmt.Println(color.Red + "Error message" + color.Reset)
```

### Combination Pattern
```go
fmt.Println(color.Bold + color.Red + "Bold red text" + color.Reset)
```

### Helper Function Pattern
```go
func printError(msg string) {
    fmt.Println(color.Bold + color.Red + "✗ " + msg + color.Reset)
}
```

## Common Use Cases

1. **CLI Applications** - Status messages, error handling, user feedback
2. **Logging Systems** - Color-coded log levels (INFO, WARN, ERROR)
3. **Debug Output** - Highlighting important information
4. **Progress Indicators** - Colored progress bars and status updates
5. **Data Visualization** - Terminal-based charts and graphs

## Design Principles

- **Zero Dependencies** - Pure Go standard library only
- **High Performance** - Pre-computed escape sequences, no allocations
- **Simple API** - String concatenation based, no complex objects
- **Cross-Platform** - Works on all platforms that support ANSI escape sequences

## Testing Notes

- All functions have unit tests with expected ANSI output
- Benchmarks measure performance of color generation
- Edge cases include invalid hex values, out-of-range inputs
- Examples serve as integration tests

## Package Stability

- **Status**: Stable, ready for v1.0.0
- **API**: Stable, backward compatible
- **Dependencies**: None (zero external dependencies)
- **Go Version**: Requires Go 1.18+ (uses generics in tests)