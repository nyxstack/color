package main

import (
	"fmt"
	"time"

	"github.com/nyxstack/color"
)

func main() {
	// Header with title
	fmt.Println()
	fmt.Printf("  %s%s╔══════════════════════════════════════════════════════════════╗%s\n", color.Bold, color.Cyan, color.Reset)
	fmt.Printf("  %s%s║                    🎨 NYX COLOR PACKAGE 🎨                   ║%s\n", color.Bold, color.Cyan, color.Reset)
	fmt.Printf("  %s%s║               Terminal Colors Made Beautiful                 ║%s\n", color.Bold, color.Cyan, color.Reset)
	fmt.Printf("  %s%s╚══════════════════════════════════════════════════════════════╝%s\n", color.Bold, color.Cyan, color.Reset)
	fmt.Println()

	// Basic Colors Demo
	fmt.Printf("%s%s▶ Basic Colors:%s\n", color.Bold, color.White, color.Reset)
	colors := []struct {
		name  string
		code  string
		emoji string
	}{
		{"Red", color.Red, "🔴"},
		{"Green", color.Green, "🟢"},
		{"Blue", color.Blue, "🔵"},
		{"Yellow", color.Yellow, "🟡"},
		{"Magenta", color.Magenta, "🟣"},
		{"Cyan", color.Cyan, "🔷"},
	}

	for _, c := range colors {
		fmt.Printf("  %s %s%-8s%s %s●●●●●%s\n", c.emoji, c.code, c.name, color.Reset, c.code, color.Reset)
	}
	fmt.Println()

	// Text Styles Demo
	fmt.Printf("%s%s▶ Text Styles:%s\n", color.Bold, color.White, color.Reset)
	fmt.Printf("  %s💪 Bold Text%s\n", color.Bold+color.Yellow, color.Reset)
	fmt.Printf("  %s📝 Italic Text%s\n", color.Italic+color.Green, color.Reset)
	fmt.Printf("  %s📑 Underlined Text%s\n", color.Underline+color.Blue, color.Reset)
	fmt.Printf("  %s❌ Strikethrough Text%s\n", color.Strikethrough+color.Red, color.Reset)
	fmt.Println()

	// RGB Colors Demo
	fmt.Printf("%s%s▶ RGB & HEX Colors:%s\n", color.Bold, color.White, color.Reset)
	rgbColors := []struct {
		name string
		hex  string
	}{
		{"🌅 Sunset Orange", "#FF6B35"},
		{"🌊 Ocean Blue", "#006994"},
		{"🌸 Cherry Blossom", "#FFB7C5"},
		{"🍃 Forest Green", "#228B22"},
		{"🌙 Midnight Purple", "#2E0854"},
	}

	for _, c := range rgbColors {
		fmt.Printf("  %s%-20s%s %s██████%s\n", color.Hex(c.hex), c.name, color.Reset, color.Hex(c.hex), color.Reset)
	}
	fmt.Println()

	// Progress Bar Demo
	fmt.Printf("%s%s▶ Animated Progress:%s\n", color.Bold, color.White, color.Reset)
	for i := 0; i <= 100; i += 20 {
		fmt.Print("  📊 Loading: [")
		for j := 0; j < 25; j++ {
			if j < i/4 {
				if j < 6 {
					fmt.Printf("%s█%s", color.Red, color.Reset)
				} else if j < 12 {
					fmt.Printf("%s█%s", color.Yellow, color.Reset)
				} else {
					fmt.Printf("%s█%s", color.Green, color.Reset)
				}
			} else {
				fmt.Printf("%s░%s", color.RGB(60, 60, 60), color.Reset)
			}
		}
		fmt.Printf("] %s%d%%%s\r", color.Bold, i, color.Reset)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println()

	// Log Levels Demo
	fmt.Printf("%s%s▶ Log Level Styling:%s\n", color.Bold, color.White, color.Reset)
	logs := []struct {
		level string
		hex   string
		icon  string
		msg   string
	}{
		{"INFO", "#17A2B8", "ℹ️", "Server started successfully"},
		{"WARN", "#FFC107", "⚠️", "High memory usage detected"},
		{"ERROR", "#DC3545", "❌", "Database connection failed"},
		{"SUCCESS", "#28A745", "✅", "Backup completed"},
	}

	for _, log := range logs {
		timestamp := "15:04:05"
		fmt.Printf("  %s[%s]%s %s[%-7s]%s %s %s\n",
			color.RGB(128, 128, 128), timestamp, color.Reset,
			color.Hex(log.hex), log.level, color.Reset,
			log.icon, log.msg)
	}
	fmt.Println()

	// 256 Color Gradient Demo
	fmt.Printf("%s%s▶ 256 Color Gradient:%s\n", color.Bold, color.White, color.Reset)
	fmt.Print("  🌈 ")
	for i := 0; i < 50; i++ {
		colorCode := uint8(196 + (i % 6))
		if i < 8 {
			colorCode = 196 // Red
		} else if i < 16 {
			colorCode = 208 // Orange
		} else if i < 24 {
			colorCode = 226 // Yellow
		} else if i < 32 {
			colorCode = 46 // Green
		} else if i < 40 {
			colorCode = 51 // Cyan
		} else {
			colorCode = 21 // Blue
		}
		fmt.Printf("%s▓%s", color.Color256(colorCode), color.Reset)
	}
	fmt.Println()
	fmt.Println()

	// Background Colors Demo
	fmt.Printf("%s%s▶ Background Colors:%s\n", color.Bold, color.White, color.Reset)
	fmt.Printf("  %s%s  RED BG  %s ", color.BgRed, color.White, color.Reset)
	fmt.Printf("%s%s  GREEN BG  %s ", color.BgGreen, color.Black, color.Reset)
	fmt.Printf("%s%s  BLUE BG  %s\n", color.BgBlue, color.White, color.Reset)
	fmt.Println()

	// Footer
	fmt.Printf("  %s%s┌─────────────────────────────────────────────────────────────┐%s\n", color.Bold, color.BrightBlack, color.Reset)
	fmt.Printf("  %s%s│  🚀 github.com/nyxstack/color - Fast & Lightweight Colors  │%s\n", color.Bold, color.BrightBlack, color.Reset)
	fmt.Printf("  %s%s│  📦 Support: 16/256/RGB/HEX colors + Text Styling          │%s\n", color.Bold, color.BrightBlack, color.Reset)
	fmt.Printf("  %s%s└─────────────────────────────────────────────────────────────┘%s\n", color.Bold, color.BrightBlack, color.Reset)
	fmt.Println()
}
