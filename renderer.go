package main

import (
	"fmt"
	"strings"
)

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorBold   = "\033[1m"
)

// GetANSIColor returns the ANSI color code for a color
func GetANSIColor(c Color) string {
	switch c {
	case ColorCyan:
		return "\033[46m"  // Cyan background
	case ColorYellow:
		return "\033[43m"  // Yellow background
	case ColorPurple:
		return "\033[45m"  // Magenta background
	case ColorGreen:
		return "\033[42m"  // Green background
	case ColorRed:
		return "\033[41m"  // Red background
	case ColorBlue:
		return "\033[44m"  // Blue background
	case ColorOrange:
		return "\033[48;5;208m"  // Orange background (256 color)
	default:
		return ""
	}
}

// InitScreen initializes the alternate screen buffer
func InitScreen() {
	fmt.Print("\033[?1049h") // Enable alternative buffer
	fmt.Print("\033[?25l")   // Hide cursor
}

// ClearScreen clears the terminal screen
func ClearScreen() {
	fmt.Print("\033[2J\033[H")
}

// RestoreScreen restores the main screen buffer
func RestoreScreen() {
	fmt.Print("\033[?25h")   // Show cursor
	fmt.Print("\033[?1049l") // Disable alternative buffer
}

// Render renders the game board
func Render(b *Board) {
	// Move cursor to home position
	fmt.Print("\033[H")

	// Create a copy of the grid to render
	display := make([][]Color, BoardHeight)
	for i := range display {
		display[i] = make([]Color, BoardWidth)
		copy(display[i], b.Grid[i])
	}

	// Add current tetromino to display
	if b.Current != nil {
		for row := 0; row < len(b.Current.Shape); row++ {
			for col := 0; col < len(b.Current.Shape[row]); col++ {
				if b.Current.Shape[row][col] == 0 {
					continue
				}

				x := b.Current.X + col
				y := b.Current.Y + row

				if y >= 0 && y < BoardHeight && x >= 0 && x < BoardWidth {
					display[y][x] = b.Current.Color
				}
			}
		}
	}

	// Build output
	var output strings.Builder

	// Print header
	output.WriteString(ColorBold + "╔════════════════════╗" + ColorReset + "\r\n")
	output.WriteString(ColorBold + "║   GO TETRIS !!     ║" + ColorReset + "\r\n")
	output.WriteString(ColorBold + "╠════════════════════╣" + ColorReset + "\r\n")

	// Print board
	for row := 0; row < BoardHeight; row++ {
		output.WriteString(ColorBold + "║" + ColorReset)
		for col := 0; col < BoardWidth; col++ {
			color := display[row][col]
			if color == ColorNone {
				output.WriteString("  ")
			} else {
				output.WriteString(GetANSIColor(color) + "  " + ColorReset)
			}
		}
		output.WriteString(ColorBold + "║" + ColorReset)

		// Print stats on the right
		switch row {
		case 1:
			output.WriteString(fmt.Sprintf("  Score: %d", b.Score))
		case 2:
			output.WriteString(fmt.Sprintf("  Lines: %d", b.Lines))
		case 3:
			output.WriteString(fmt.Sprintf("  Level: %d", b.Level))
		case 5:
			output.WriteString("  Next:")
		case 6, 7, 8, 9:
			if b.Next != nil && row-6 < len(b.Next.Shape) {
				output.WriteString("  ")
				for col := 0; col < len(b.Next.Shape[row-6]); col++ {
					if b.Next.Shape[row-6][col] == 1 {
						output.WriteString(GetANSIColor(b.Next.Color) + "  " + ColorReset)
					} else {
						output.WriteString("  ")
					}
				}
			}
		case 12:
			output.WriteString("  Controls:")
		case 13:
			output.WriteString("  WASD or Arrows: Move/Rotate")
		case 14:
			output.WriteString("  Space: Hard drop")
		case 15:
			output.WriteString("  Q: Quit")
		}
		output.WriteString("\r\n")
	}

	// Print footer
	output.WriteString(ColorBold + "╚════════════════════╝" + ColorReset + "\r\n")

	if b.GameOver {
		output.WriteString("\r\n")
		output.WriteString(strings.Repeat(" ", 5) + ColorBold + "GAME OVER!" + ColorReset + "\r\n")
		output.WriteString(strings.Repeat(" ", 4) + "Press Q to quit\r\n")
	}

	// Output everything at once
	fmt.Print(output.String())
}
