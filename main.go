package main

import (
	"time"
)

func main() {
	// Setup terminal
	if err := SetupTerminal(); err != nil {
		panic(err)
	}
	defer RestoreTerminal()

	// Initialize screen
	InitScreen()
	defer RestoreScreen()
	ClearScreen()

	// Create game board
	board := NewBoard()

	// Input channel
	inputChan := make(chan []byte, 10)
	go func() {
		for {
			key, err := GetKey()
			if err == nil && len(key) > 0 {
				inputChan <- key
			}
		}
	}()

	// Game loop
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	lastDrop := time.Now()

	for !board.GameOver {
		// Calculate drop interval based on level
		dropInterval := time.Duration(500-board.Level*30) * time.Millisecond
		if dropInterval < 100*time.Millisecond {
			dropInterval = 100 * time.Millisecond
		}

		select {
		case keys := <-inputChan:
			// Handle key input
			if len(keys) == 1 {
				// Single byte keys
				switch keys[0] {
				case 'q', 'Q', 3: // q, Q, or Ctrl+C
					return
				case 'a', 'A':
					board.MoveTetromino(-1, 0)
				case 'd', 'D':
					board.MoveTetromino(1, 0)
				case 'w', 'W':
					board.RotateTetromino()
				case 's', 'S':
					board.DropTetromino()
					lastDrop = time.Now()
				case ' ':
					board.HardDrop()
					lastDrop = time.Now()
				}
			} else if len(keys) == 3 && keys[0] == 27 && keys[1] == 91 {
				// Arrow keys (ESC [ X)
				switch keys[2] {
				case 68: // Left arrow
					board.MoveTetromino(-1, 0)
				case 67: // Right arrow
					board.MoveTetromino(1, 0)
				case 65: // Up arrow
					board.RotateTetromino()
				case 66: // Down arrow
					board.DropTetromino()
					lastDrop = time.Now()
				}
			}
			Render(board)

		case <-ticker.C:
			// Auto drop
			if time.Since(lastDrop) >= dropInterval {
				board.DropTetromino()
				lastDrop = time.Now()
				Render(board)
			}
		}
	}

	// Game over screen
	Render(board)

	// Wait for quit
	for {
		key, err := GetKey()
		if err == nil && len(key) > 0 && (key[0] == 'q' || key[0] == 'Q' || key[0] == 3) {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
