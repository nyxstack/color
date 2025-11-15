package main

import (
	"fmt"

	"github.com/nyxstack/color"
)

func main() {
	fmt.Println("=== 256 Color Examples ===")

	// Standard 16 colors (0-15)
	fmt.Println("\n=== Standard 16 Colors (0-15) ===")
	for i := 0; i < 16; i++ {
		fmt.Printf("%s%3d%s ", color.Color256(uint8(i)), i, color.Reset)
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}

	// 6x6x6 color cube (16-231)
	fmt.Println("\n\n=== 6x6x6 Color Cube (16-231) ===")
	for r := 0; r < 6; r++ {
		for g := 0; g < 6; g++ {
			for b := 0; b < 6; b++ {
				colorNum := 16 + (r * 36) + (g * 6) + b
				fmt.Printf("%s%3d%s ", color.Color256(uint8(colorNum)), colorNum, color.Reset)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}

	// Grayscale colors (232-255)
	fmt.Println("\n=== Grayscale Colors (232-255) ===")
	for i := 232; i <= 255; i++ {
		fmt.Printf("%s%3d%s ", color.Color256(uint8(i)), i, color.Reset)
		if (i-231)%8 == 0 {
			fmt.Println()
		}
	}

	// Demonstrating practical usage
	fmt.Println("\n\n=== Practical 256 Color Usage ===")

	// Different shades of blue
	blueShades := []uint8{17, 18, 19, 20, 21, 33, 39, 45, 51, 57, 63, 69, 75, 81, 87, 93, 99, 105, 111, 117, 123, 129, 135, 141, 147, 153, 159, 165, 171, 177, 183, 189, 195, 201, 207, 213, 219, 225, 231}
	fmt.Println("Blue gradient:")
	for _, shade := range blueShades {
		fmt.Printf("%s█%s", color.Color256(shade), color.Reset)
	}
	fmt.Println()

	// Rainbow effect
	fmt.Println("\nRainbow text:")
	rainbowColors := []uint8{196, 202, 208, 214, 220, 226, 190, 154, 118, 82, 46, 47, 48, 49, 50, 51, 45, 39, 33, 27, 21, 57, 93, 129, 165, 201, 200, 199, 198, 197}
	text := "Rainbow colored text example!"
	for i, char := range text {
		colorIndex := rainbowColors[i%len(rainbowColors)]
		fmt.Printf("%s%c%s", color.Color256(colorIndex), char, color.Reset)
	}
	fmt.Println()

	// Creating a simple chart with colors
	fmt.Println("\n=== Color Chart Example ===")
	data := []int{5, 8, 3, 12, 7, 15, 9, 6}
	labels := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug"}
	chartColors := []uint8{196, 208, 226, 46, 51, 33, 93, 165}

	for i, value := range data {
		fmt.Printf("%s%-3s%s: ", color.Bold, labels[i], color.Reset)
		for j := 0; j < value; j++ {
			fmt.Printf("%s█%s", color.Color256(chartColors[i]), color.Reset)
		}
		fmt.Printf(" (%d)\n", value)
	}

	// Temperature scale example
	fmt.Println("\n=== Temperature Scale ===")
	temperatures := []int{-10, 0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	tempColors := []uint8{21, 33, 39, 45, 51, 87, 123, 159, 195, 226, 220, 196}

	for i, temp := range temperatures {
		fmt.Printf("%s%3d°C%s ", color.Color256(tempColors[i]), temp, color.Reset)
		if (i+1)%6 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
