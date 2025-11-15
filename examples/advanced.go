package main

import (
	"fmt"
	"math"
	"time"

	"github.com/nyxstack/color"
)

func main() {
	fmt.Println("=== Advanced Color Examples ===")

	// RGB colors
	fmt.Println("\n=== RGB Colors ===")
	fmt.Printf("%sPure Red (255,0,0)%s\n", color.RGB(255, 0, 0), color.Reset)
	fmt.Printf("%sPure Green (0,255,0)%s\n", color.RGB(0, 255, 0), color.Reset)
	fmt.Printf("%sPure Blue (0,0,255)%s\n", color.RGB(0, 0, 255), color.Reset)
	fmt.Printf("%sOrange (255,165,0)%s\n", color.RGB(255, 165, 0), color.Reset)
	fmt.Printf("%sPurple (128,0,128)%s\n", color.RGB(128, 0, 128), color.Reset)
	fmt.Printf("%sPink (255,192,203)%s\n", color.RGB(255, 192, 203), color.Reset)

	// HEX colors
	fmt.Println("\n=== HEX Colors ===")
	fmt.Printf("%sRed (#FF0000)%s\n", color.Hex("#FF0000"), color.Reset)
	fmt.Printf("%sGreen (#00FF00)%s\n", color.Hex("00FF00"), color.Reset)
	fmt.Printf("%sBlue (#0000FF)%s\n", color.Hex("0000FF"), color.Reset)
	fmt.Printf("%sGold (#FFD700)%s\n", color.Hex("FFD700"), color.Reset)
	fmt.Printf("%sCoral (#FF7F50)%s\n", color.Hex("#FF7F50"), color.Reset)
	fmt.Printf("%sTeal (#008080)%s\n", color.Hex("008080"), color.Reset)

	// Gradient examples
	fmt.Println("\n=== RGB Gradients ===")
	createRedToBlueGradient()
	createRainbowGradient()

	// Color palettes
	fmt.Println("\n=== Color Palettes ===")
	displayPalette("Sunset", []string{"#FF6B35", "#F7931E", "#FFD23F", "#6BCF7F", "#4D96FF", "#9B59B6"})
	displayPalette("Ocean", []string{"#006994", "#13A5B7", "#26BCCF", "#4DD5EA", "#80E6FF", "#B3F2FF"})
	displayPalette("Forest", []string{"#2D5016", "#3E6B1C", "#4F7F23", "#61942A", "#73A931", "#85BD38"})

	// Progress bar example
	fmt.Println("\n=== Progress Bar Example ===")
	animatedProgressBar()

	// Log level styling
	fmt.Println("\n=== Log Level Styling ===")
	logExamples()

	// ASCII art with colors
	fmt.Println("\n=== Colored ASCII Art ===")
	coloredASCIIArt()

	// Color mixing example
	fmt.Println("\n=== Color Mixing ===")
	colorMixingDemo()
}

// Create a gradient from red to blue
func createRedToBlueGradient() {
	fmt.Println("\nRed to Blue Gradient:")
	for i := 0; i < 50; i++ {
		// Interpolate between red (255,0,0) and blue (0,0,255)
		ratio := float64(i) / 49.0
		r := uint8(255 * (1 - ratio))
		b := uint8(255 * ratio)
		fmt.Printf("%s█%s", color.RGB(r, 0, b), color.Reset)
	}
	fmt.Println()
}

// Create a rainbow gradient
func createRainbowGradient() {
	fmt.Println("\nRainbow Gradient:")
	for i := 0; i < 60; i++ {
		// Create HSV to RGB conversion for rainbow effect
		hue := float64(i) * 6.0 / 60.0 // 0 to 6
		r, g, b := hsvToRGB(hue, 1.0, 1.0)
		fmt.Printf("%s█%s", color.RGB(uint8(r*255), uint8(g*255), uint8(b*255)), color.Reset)
	}
	fmt.Println()
}

// Display a color palette
func displayPalette(name string, colors []string) {
	fmt.Printf("%s%s Palette:%s\n", color.Bold, name, color.Reset)
	for i, hexColor := range colors {
		fmt.Printf("%s  %-8s  %s", color.Hex(hexColor), hexColor, color.Reset)
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
	if len(colors)%3 != 0 {
		fmt.Println()
	}
	fmt.Println()
}

// Animated progress bar
func animatedProgressBar() {
	fmt.Println("Loading...")
	for progress := 0; progress <= 100; progress += 5 {
		fmt.Printf("\r%sProgress: %s", color.Bold, color.Reset)

		// Progress bar with gradient colors
		for i := 0; i < 50; i++ {
			if i < progress/2 {
				// Filled portion with gradient from red to green
				ratio := float64(i) / 25.0 // 0 to 2
				if ratio > 1.0 {
					ratio = 1.0
				}
				r := uint8(255 * (1 - ratio))
				g := uint8(255 * ratio)
				fmt.Printf("%s█%s", color.RGB(r, g, 0), color.Reset)
			} else {
				// Empty portion
				fmt.Printf("%s░%s", color.RGB(100, 100, 100), color.Reset)
			}
		}

		fmt.Printf(" %s%d%%%s", color.Bold, progress, color.Reset)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\n" + color.Green + "Complete!" + color.Reset)
}

// Log level examples
func logExamples() {
	logMessage("DEBUG", "#6C757D", "Database connection established")
	logMessage("INFO", "#17A2B8", "Server started on port 8080")
	logMessage("WARN", "#FFC107", "API rate limit approaching")
	logMessage("ERROR", "#DC3545", "Failed to connect to database")
	logMessage("FATAL", "#721C24", "System critical error")
}

func logMessage(level, hexColor, message string) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("%s[%s]%s %s[%-5s]%s %s\n",
		color.RGB(128, 128, 128), timestamp, color.Reset,
		color.Hex(hexColor), level, color.Reset,
		message)
}

// Colored ASCII art
func coloredASCIIArt() {
	art := []string{
		"    🌟    ",
		"   ⭐⭐⭐   ",
		"  🌟⭐🌟⭐🌟  ",
		" ⭐🌟⭐🌟⭐🌟⭐ ",
		"🌟⭐🌟⭐🌟⭐🌟⭐🌟",
		" ⭐🌟⭐🌟⭐🌟⭐ ",
		"  🌟⭐🌟⭐🌟  ",
		"   ⭐⭐⭐   ",
		"    🌟    ",
	}

	colors := []string{"#FFD700", "#FFA500", "#FF6347", "#FF1493", "#8A2BE2", "#4169E1", "#00CED1", "#00FA9A", "#ADFF2F"}

	for i, line := range art {
		fmt.Printf("%s%s%s\n", color.Hex(colors[i]), line, color.Reset)
	}
}

// Color mixing demonstration
func colorMixingDemo() {
	colors := map[string][3]uint8{
		"Red":     {255, 0, 0},
		"Green":   {0, 255, 0},
		"Blue":    {0, 0, 255},
		"Yellow":  {255, 255, 0},
		"Cyan":    {0, 255, 255},
		"Magenta": {255, 0, 255},
	}

	for name1, rgb1 := range colors {
		for name2, rgb2 := range colors {
			if name1 >= name2 {
				continue
			}

			// Mix the colors (average RGB values)
			mixedR := (rgb1[0] + rgb2[0]) / 2
			mixedG := (rgb1[1] + rgb2[1]) / 2
			mixedB := (rgb1[2] + rgb2[2]) / 2

			fmt.Printf("%s%s%s + %s%s%s = %s%s+%s%s\n",
				color.RGB(rgb1[0], rgb1[1], rgb1[2]), name1, color.Reset,
				color.RGB(rgb2[0], rgb2[1], rgb2[2]), name2, color.Reset,
				color.RGB(mixedR, mixedG, mixedB), name1, name2, color.Reset)
		}
	}
}

// HSV to RGB conversion
func hsvToRGB(h, s, v float64) (float64, float64, float64) {
	for h < 0 {
		h += 6
	}
	for h >= 6 {
		h -= 6
	}

	i := math.Floor(h)
	f := h - i
	p := v * (1 - s)
	q := v * (1 - s*f)
	t := v * (1 - s*(1-f))

	switch int(i) {
	case 0:
		return v, t, p
	case 1:
		return q, v, p
	case 2:
		return p, v, t
	case 3:
		return p, q, v
	case 4:
		return t, p, v
	case 5:
		return v, p, q
	default:
		return 0, 0, 0
	}
}
