package main

import (
	"fmt"

	"github.com/nyxstack/color"
)

func main() {
	fmt.Println("=== Basic Color Examples ===")

	// Basic text colors
	fmt.Println(color.Red + "This is red text" + color.Reset)
	fmt.Println(color.Green + "This is green text" + color.Reset)
	fmt.Println(color.Blue + "This is blue text" + color.Reset)
	fmt.Println(color.Yellow + "This is yellow text" + color.Reset)
	fmt.Println(color.Magenta + "This is magenta text" + color.Reset)
	fmt.Println(color.Cyan + "This is cyan text" + color.Reset)
	fmt.Println(color.White + "This is white text" + color.Reset)

	// Bright colors
	fmt.Println("\n=== Bright Colors ===")
	fmt.Println(color.BrightRed + "This is bright red text" + color.Reset)
	fmt.Println(color.BrightGreen + "This is bright green text" + color.Reset)
	fmt.Println(color.BrightBlue + "This is bright blue text" + color.Reset)
	fmt.Println(color.BrightYellow + "This is bright yellow text" + color.Reset)

	// Background colors
	fmt.Println("\n=== Background Colors ===")
	fmt.Println(color.BgRed + "Red background" + color.Reset)
	fmt.Println(color.BgGreen + "Green background" + color.Reset)
	fmt.Println(color.BgBlue + color.White + "Blue background with white text" + color.Reset)
	fmt.Println(color.BgYellow + color.Black + "Yellow background with black text" + color.Reset)

	// Text styles
	fmt.Println("\n=== Text Styles ===")
	fmt.Println(color.Bold + "Bold text" + color.Reset)
	fmt.Println(color.Italic + "Italic text" + color.Reset)
	fmt.Println(color.Underline + "Underlined text" + color.Reset)
	fmt.Println(color.Strikethrough + "Strikethrough text" + color.Reset)
	fmt.Println(color.Dim + "Dimmed text" + color.Reset)
	fmt.Println(color.Reverse + "Reversed text" + color.Reset)

	// Combining styles
	fmt.Println("\n=== Combined Styles ===")
	fmt.Println(color.Bold + color.Red + "Bold red text" + color.Reset)
	fmt.Println(color.Underline + color.Blue + "Underlined blue text" + color.Reset)
	fmt.Println(color.BgGreen + color.Bold + color.White + "Bold white text on green background" + color.Reset)
	fmt.Println(color.Italic + color.Magenta + "Italic magenta text" + color.Reset)

	// Creating colored output functions
	fmt.Println("\n=== Helper Functions ===")
	printSuccess("Operation completed successfully!")
	printError("An error occurred!")
	printWarning("This is a warning message.")
	printInfo("Here's some information.")
}

// Helper functions for common use cases
func printSuccess(message string) {
	fmt.Println(color.Bold + color.Green + "✓ " + message + color.Reset)
}

func printError(message string) {
	fmt.Println(color.Bold + color.Red + "✗ " + message + color.Reset)
}

func printWarning(message string) {
	fmt.Println(color.Bold + color.Yellow + "⚠ " + message + color.Reset)
}

func printInfo(message string) {
	fmt.Println(color.Bold + color.Cyan + "ℹ " + message + color.Reset)
}
